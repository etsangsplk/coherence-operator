package coherencerole

import (
	"context"
	"encoding/json"
	"fmt"
	coh "github.com/oracle/coherence-operator/pkg/apis/coherence/v1"
	"io/ioutil"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"net/http"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

// scale will scale a role up or down
func (r *ReconcileCoherenceRole) scale(role *coh.CoherenceRole, cohInternal *unstructured.Unstructured, existing *coh.CoherenceInternal, desired int32, current int32, sts *appsv1.StatefulSet) (reconcile.Result, error) {
	var policy coh.ScalingPolicy

	if role.Spec.ScalingPolicy == nil {
		// the scaling policy is not set the look at the storage enabled flag
		if role.Spec.StorageEnabled == nil || *role.Spec.StorageEnabled {
			// storage enabled is either not set or is true so do safe scaling
			policy = coh.ParallelUpSafeDownScaling
		} else {
			// storage enabled is false so do parallel scaling
			policy = coh.ParallelScaling
		}
	} else {
		// scaling policy is set so use it
		policy = *role.Spec.ScalingPolicy
	}

	switch policy {
	case coh.SafeScaling:
		return r.safeScale(role, cohInternal, existing, desired, current, sts)
	case coh.ParallelScaling:
		return r.parallelScale(role, cohInternal, existing, desired, current)
	case coh.ParallelUpSafeDownScaling:
		if desired > current {
			return r.parallelScale(role, cohInternal, existing, desired, current)
		} else {
			return r.safeScale(role, cohInternal, existing, desired, current, sts)
		}
	default:
		// shouldn't get here, but better safe than sorry
		return r.safeScale(role, cohInternal, existing, desired, current, sts)
	}
}

// safeScale will scale a role up or down by one and requeue the request.
func (r *ReconcileCoherenceRole) safeScale(role *coh.CoherenceRole, cohInternal *unstructured.Unstructured, existing *coh.CoherenceInternal, desired int32, current int32, sts *appsv1.StatefulSet) (reconcile.Result, error) {
	logger := log.WithValues("Namespace", role.Name, "Name", role.Name)

	if sts.Status.ReadyReplicas != current {
		logger.Info(fmt.Sprintf("Role %s is not StatusHA - re-queing scaling request. Stateful set ready replicas is %d", role.Name, sts.Status.ReadyReplicas))
	}

	ha := current == 1 || r.IsStatusHA(role, sts)

	if ha {
		var replicas int32

		if desired > current {
			replicas = current + 1
		} else {
			replicas = current - 1
		}

		logger.Info(fmt.Sprintf("Role %s is StatusHA, safely scaling from %d to %d (final desired replicas %d)", role.Name, current, replicas, desired))

		// use the parallel method to just scale by one
		_, err := r.parallelScale(role, cohInternal, existing, replicas, current)
		if err == nil {
			if replicas == desired {
				// we're at the desired size so finished scaling
				return reconcile.Result{Requeue: false}, nil
			} else {
				// scaled by one but not yet at the desired size - requeue request after one minute
				return reconcile.Result{Requeue: true, RequeueAfter: time.Minute}, nil
			}
		} else {
			// failed
			return r.handleErrAndRequeue(err, role, fmt.Sprintf(failedToScaleRole, role.Name, current, replicas, err.Error()), logger)
		}
	}

	// Not StatusHA - wait one minute
	logger.Info(fmt.Sprintf("Role %s is not StatusHA - re-queing scaling request", role.Name))
	return reconcile.Result{Requeue: true, RequeueAfter: r.statusHARetry}, nil
}

// parallelScale will scale the role by the required amount in one request.
func (r *ReconcileCoherenceRole) parallelScale(role *coh.CoherenceRole, cohInternal *unstructured.Unstructured, existing *coh.CoherenceInternal, desired int32, current int32) (reconcile.Result, error) {
	// update the CoherenceInternal, this should trigger an update of the Helm install to scale the StatefulSet

	// Update this CoherenceRole's status
	role.Status.Status = coh.RoleStatusScaling
	role.Status.Replicas = desired
	err := r.client.Status().Update(context.TODO(), role)
	if err != nil {
		// failed to update the CoherenceRole's status
		// ToDo - handle this properly by re-queuing the request and then in the reconcile method properly handle setting status even if the role is in the desired state
		log.Error(err, "failed to update role status")
	}

	existing.Spec.ClusterSize = desired
	cohInternal.Object["spec"] = existing.Spec
	err = r.client.Update(context.TODO(), cohInternal)
	if err != nil {
		// send a failed scale event
		msg := fmt.Sprintf("failed to scale Helm install %s in CoherenceRole %s from %d to %d", role.Name, role.Name, current, desired)
		r.events.Event(role, corev1.EventTypeNormal, "SuccessfulScale", msg)

		return reconcile.Result{}, err
	}

	// send a successful scale event
	msg := fmt.Sprintf("scaled Helm install %s in CoherenceRole %s from %d to %d", role.Name, role.Name, current, desired)
	r.events.Event(role, corev1.EventTypeNormal, eventReasonScale, msg)

	return reconcile.Result{}, nil
}

// IsStatusHA will return true if the cluster represented by the role is StatusHA.
func (r *ReconcileCoherenceRole) IsStatusHA(role *coh.CoherenceRole, sts *appsv1.StatefulSet) bool {
	list := corev1.PodList{}
	opts := client.ListOptions{}
	opts.InNamespace(role.Namespace)
	opts.MatchingLabels(sts.Spec.Selector.MatchLabels)

	if log.Enabled() {
		log.Info("Checking StatefulSet "+sts.Name+" for StatusHA", "Namespace", role.Name, "Name", role.Name)
	}

	err := r.client.List(context.TODO(), &opts, &list)
	if err != nil {
		log.Error(err, "Error getting list of Pods for StatefulSet "+sts.Name)
		return false
	}

	if len(list.Items) == 0 {
		if log.Enabled() {
			log.Info("Cannot find any Pods for StatefulSet " + sts.Name + " - assuming StatusHA is true")
		}
		return true
	}

	for _, pod := range list.Items {
		if pod.Status.Phase == "Running" {
			ip := pod.Status.PodIP
			ha, err := IsPodStatusHA(ip)
			if log.Enabled() {
				log.Info("Checking pod " + pod.Name + " for StatusHA")
			}
			if err == nil {
				return ha
			}
		} else {
			log.Info("Skipping StatusHA checking for pod " + pod.Name + " as Pod status not in running phase")
		}
	}

	return false
}

// Determine whether a Pod's Coherence Services are StatusHA.
func IsPodStatusHA(podIP string) (bool, error) {
	cl := &http.Client{}

	url := fmt.Sprintf(servicesFormat, podIP)
	response, err := cl.Get(url)
	if err != nil {
		if log.Enabled() {
			log.Info("Error accessing podIP " + podIP + " URL " + url + "\n" + err.Error())
		}
		return false, err
	}

	data, _ := ioutil.ReadAll(response.Body)
	services := RestData{}

	err = json.Unmarshal(data, &services)
	if err != nil {
		if log.Enabled() {
			log.Info("Error parsing services json returned from podIP " + podIP + " URL " + url + "\n" + string(data) + "\n" + err.Error())
		}
		return false, err
	}

	for _, service := range services.Items {
		if service["type"] == "DistributedCache" {
			url = fmt.Sprintf(partitionFormat, podIP, service["name"])
			response, err := cl.Get(url)
			if err == nil {
				if response.StatusCode == 200 {
					data, _ = ioutil.ReadAll(response.Body)
					fields := &PartitionData{}
					err = json.Unmarshal(data, &fields)
					if err == nil {
						// we must have more than one service member and backups > 0 to event think about being HA.
						if fields.BackupCount > 0 && fields.ServiceNodeCount > 1 {
							if fields.HAStatusCode <= 1 || fields.RemainingDistributionCount != 0 {
								// we're not HA
								return false, nil
							}
						}
					} else {
						if log.Enabled() {
							log.Info("Error checking StatusHA on podIP " + podIP + "\n" + err.Error())
						}
						return false, err
					}
				}
			} else {
				log.Info("Error accessing podIP " + podIP + " URL " + url + "\n" + err.Error())
			}
		}
	}

	return true, nil
}

const (
	// The URL pattern for Coherence management services query.
	servicesFormat = "http://%s:30000/management/coherence/cluster/services"
	// The URL pattern for Coherence management partition assignment query.
	partitionFormat = "http://%s:30000/management/coherence/cluster/services/%s/partition"
)

// A struct to use to hold the results of a Coherence management ReST query.
type RestData struct {
	Links []map[string]interface{}
	Items []map[string]interface{}
}

// A struct to use to hold the results of a Coherence management ReST partition assignment query.
type PartitionData struct {
	HAStatus                   string `json:"HAStatus"`
	HAStatusCode               int    `json:"HAStatusCode"`
	RemainingDistributionCount int    `json:"remainingDistributionCount"`
	BackupCount                int    `json:"backupCount"`
	ServiceNodeCount           int    `json:"serviceNodeCount"`
}