package helper

import (
	"errors"
	"github.com/ghodss/yaml"
	coh "github.com/oracle/coherence-operator/pkg/apis/coherence/v1"
	"io/ioutil"
	coreV1 "k8s.io/api/core/v1"
	"os"
)

type OperatorValues struct {
	// The name to use for the service account to use when RBAC is enabled
	// The role bindings must already have been created as this chart does not create them it just
	// sets the serviceAccountName value in the Pod spec.
	// +optional
	ServiceAccountName string `json:"serviceAccountName,omitempty"`
	// The secrets to be used when pulling images. Secrets must be manually created in the target namespace.
	// ref: https://kubernetes.io/docs/tasks/configure-pod-container/pull-image-private-registry/
	// +optional
	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`
	// Affinity controls Pod scheduling preferences.
	// ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
	// +optional
	Affinity *coreV1.Affinity `json:"affinity,omitempty"`
	// NodeSelector is the Node labels for pod assignment
	// ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#nodeselector
	// +optional
	NodeSelector *coreV1.NodeSelector `json:"nodeSelector,omitempty"`
	// If specified, the pod's tolerations.
	// +optional
	Tolerations *[]coreV1.Toleration `json:"tolerations,omitempty"`
	// Service groups the values used to configure the internal K8s service.
	// +optional
	Service *OperatorServiceSpec `json:"service,omitempty"`
	// Image groups the values used to configure the image used
	// +optional
	CoherenceOperator *coh.ImageSpec `json:"coherenceOperator,omitempty"`
	// Controls whether or not log capture via EFK stack is enabled.
	// +optional
	LogCaptureEnabled bool `json:"logCaptureEnabled,omitempty"`
	// Specify the docker image containing Elasticsearch.
	// These parameters are ignored if 'logCaptureEnabled' is false
	// or elasticsearchEndpoinit is set.
	// +optional
	Elasticsearch *coh.ImageSpec `json:"elasticsearch,omitempty"`
	// Specify the fluentd image
	// These parameters are ignored if 'logCaptureEnabled' is false.
	// +optional
	Fluentd *coh.ImageSpec `json:"fluentd,omitempty"`
	// Specify the kibana image
	// These parameters are ignored if 'logCaptureEnabled' is false.
	// +optional
	Kibana *coh.ImageSpec `json:"kibana,omitempty"`
	// Specifies values for Kibana Dashboard Imports if logCaptureEnabled is true
	// +optional
	DashboardImport *DashboardImportSpec `json:"dashboardImport,omitempty"`
}

// LoadFromYaml loads the data from the specified YAML file into this OperatorValues
func (v *OperatorValues) LoadFromYaml(file string) error {
	if v == nil {
		return errors.New("attempted to load yaml into a nil OperatorValues reference")
	}
	_, err := os.Stat(file)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, v)
}

type DashboardImportSpec struct {
	Timeout   int32          `json:"timeout,omitempty"`
	Xpackauth *XpackAuthSpec `json:"xpackauth,omitempty"`
}

type XpackAuthSpec struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type OperatorServiceSpec struct {
	// The name of the service. It must be unique among services.
	// If not set the operator will use the service name "coherence-operator-service".
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	Type coreV1.ServiceType `json:"type,omitempty"`
	// The external domain name.
	// +optional
	Domain string `json:"domain,omitempty"`
	// Only applies to Service Type: LoadBalancer
	// LoadBalancer will get created with the IP specified in this field.
	// This feature depends on whether the underlying cloud-provider supports specifying
	// the loadBalancerIP when a load balancer is created.
	// This field will be ignored if the cloud-provider does not support the feature.
	// +optional
	LoadBalancerIP string `json:"loadBalancerIP,omitempty" protobuf:"bytes,8,opt,name=loadBalancerIP"`
	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty"`
}