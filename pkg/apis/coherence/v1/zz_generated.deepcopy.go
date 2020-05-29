// +build !ignore_autogenerated

/*
 * Copyright (c) 2020, 2020 Oracle and/or its affiliates. All rights reserved.
 * Licensed under the Universal Permissive License v 1.0 as shown at
 * http://oss.oracle.com/licenses/upl.
 */

// Code generated by operator-sdk. DO NOT EDIT.

package v1

import (
	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	status "github.com/operator-framework/operator-sdk/pkg/status"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationSpec) DeepCopyInto(out *ApplicationSpec) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Main != nil {
		in, out := &in.Main, &out.Main
		*out = new(string)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.WorkingDir != nil {
		in, out := &in.WorkingDir, &out.WorkingDir
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationSpec.
func (in *ApplicationSpec) DeepCopy() *ApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(ApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Coherence) DeepCopyInto(out *Coherence) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Coherence.
func (in *Coherence) DeepCopy() *Coherence {
	if in == nil {
		return nil
	}
	out := new(Coherence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Coherence) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceList) DeepCopyInto(out *CoherenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Coherence, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceList.
func (in *CoherenceList) DeepCopy() *CoherenceList {
	if in == nil {
		return nil
	}
	out := new(CoherenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoherenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceResourceSpec) DeepCopyInto(out *CoherenceResourceSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Cluster != nil {
		in, out := &in.Cluster, &out.Cluster
		*out = new(string)
		**out = **in
	}
	if in.Coherence != nil {
		in, out := &in.Coherence, &out.Coherence
		*out = new(CoherenceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Application != nil {
		in, out := &in.Application, &out.Application
		*out = new(ApplicationSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JVM != nil {
		in, out := &in.JVM, &out.JVM
		*out = new(JVMSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]NamedPortSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Scaling != nil {
		in, out := &in.Scaling, &out.Scaling
		*out = new(ScalingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.StartQuorum != nil {
		in, out := &in.StartQuorum, &out.StartQuorum
		*out = make([]StartQuorum, len(*in))
		copy(*out, *in)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SideCars != nil {
		in, out := &in.SideCars, &out.SideCars
		*out = make([]corev1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ConfigMapVolumes != nil {
		in, out := &in.ConfigMapVolumes, &out.ConfigMapVolumes
		*out = make([]ConfigMapVolumeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecretVolumes != nil {
		in, out := &in.SecretVolumes, &out.SecretVolumes
		*out = make([]SecretVolumeSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]corev1.PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HealthPort != nil {
		in, out := &in.HealthPort, &out.HealthPort
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(ReadinessProbeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(ReadinessProbeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ShareProcessNamespace != nil {
		in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
		*out = new(bool)
		**out = **in
	}
	if in.HostIPC != nil {
		in, out := &in.HostIPC, &out.HostIPC
		*out = new(bool)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(NetworkSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CoherenceUtils != nil {
		in, out := &in.CoherenceUtils, &out.CoherenceUtils
		*out = new(ImageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.AutomountServiceAccountToken != nil {
		in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
		*out = new(bool)
		**out = **in
	}
	if in.OperatorRequestTimeout != nil {
		in, out := &in.OperatorRequestTimeout, &out.OperatorRequestTimeout
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceResourceSpec.
func (in *CoherenceResourceSpec) DeepCopy() *CoherenceResourceSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceResourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceResourceStatus) DeepCopyInto(out *CoherenceResourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(status.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceResourceStatus.
func (in *CoherenceResourceStatus) DeepCopy() *CoherenceResourceStatus {
	if in == nil {
		return nil
	}
	out := new(CoherenceResourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceSpec) DeepCopyInto(out *CoherenceSpec) {
	*out = *in
	if in.CacheConfig != nil {
		in, out := &in.CacheConfig, &out.CacheConfig
		*out = new(string)
		**out = **in
	}
	if in.OverrideConfig != nil {
		in, out := &in.OverrideConfig, &out.OverrideConfig
		*out = new(string)
		**out = **in
	}
	if in.StorageEnabled != nil {
		in, out := &in.StorageEnabled, &out.StorageEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Persistence != nil {
		in, out := &in.Persistence, &out.Persistence
		*out = new(PersistenceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(int32)
		**out = **in
	}
	if in.Management != nil {
		in, out := &in.Management, &out.Management
		*out = new(PortSpecWithSSL)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = new(PortSpecWithSSL)
		(*in).DeepCopyInto(*out)
	}
	if in.Tracing != nil {
		in, out := &in.Tracing, &out.Tracing
		*out = new(CoherenceTracingSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ExcludeFromWKA != nil {
		in, out := &in.ExcludeFromWKA, &out.ExcludeFromWKA
		*out = new(bool)
		**out = **in
	}
	if in.SkipVersionCheck != nil {
		in, out := &in.SkipVersionCheck, &out.SkipVersionCheck
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceSpec.
func (in *CoherenceSpec) DeepCopy() *CoherenceSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoherenceTracingSpec) DeepCopyInto(out *CoherenceTracingSpec) {
	*out = *in
	if in.Ratio != nil {
		in, out := &in.Ratio, &out.Ratio
		x := (*in).DeepCopy()
		*out = &x
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoherenceTracingSpec.
func (in *CoherenceTracingSpec) DeepCopy() *CoherenceTracingSpec {
	if in == nil {
		return nil
	}
	out := new(CoherenceTracingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapVolumeSpec) DeepCopyInto(out *ConfigMapVolumeSpec) {
	*out = *in
	if in.MountPropagation != nil {
		in, out := &in.MountPropagation, &out.MountPropagation
		*out = new(corev1.MountPropagationMode)
		**out = **in
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.KeyToPath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultMode != nil {
		in, out := &in.DefaultMode, &out.DefaultMode
		*out = new(int32)
		**out = **in
	}
	if in.Optional != nil {
		in, out := &in.Optional, &out.Optional
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapVolumeSpec.
func (in *ConfigMapVolumeSpec) DeepCopy() *ConfigMapVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigMapVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.ImagePullPolicy != nil {
		in, out := &in.ImagePullPolicy, &out.ImagePullPolicy
		*out = new(corev1.PullPolicy)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JVMSpec) DeepCopyInto(out *JVMSpec) {
	*out = *in
	if in.Classpath != nil {
		in, out := &in.Classpath, &out.Classpath
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.LoggingConfig != nil {
		in, out := &in.LoggingConfig, &out.LoggingConfig
		*out = new(string)
		**out = **in
	}
	if in.Debug != nil {
		in, out := &in.Debug, &out.Debug
		*out = new(JvmDebugSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UseContainerLimits != nil {
		in, out := &in.UseContainerLimits, &out.UseContainerLimits
		*out = new(bool)
		**out = **in
	}
	if in.Gc != nil {
		in, out := &in.Gc, &out.Gc
		*out = new(JvmGarbageCollectorSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DiagnosticsVolume != nil {
		in, out := &in.DiagnosticsVolume, &out.DiagnosticsVolume
		*out = new(corev1.VolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(JvmMemorySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Jmxmp != nil {
		in, out := &in.Jmxmp, &out.Jmxmp
		*out = new(JvmJmxmpSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.UseJibClasspath != nil {
		in, out := &in.UseJibClasspath, &out.UseJibClasspath
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JVMSpec.
func (in *JVMSpec) DeepCopy() *JVMSpec {
	if in == nil {
		return nil
	}
	out := new(JVMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmDebugSpec) DeepCopyInto(out *JvmDebugSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		*out = new(bool)
		**out = **in
	}
	if in.Attach != nil {
		in, out := &in.Attach, &out.Attach
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmDebugSpec.
func (in *JvmDebugSpec) DeepCopy() *JvmDebugSpec {
	if in == nil {
		return nil
	}
	out := new(JvmDebugSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmGarbageCollectorSpec) DeepCopyInto(out *JvmGarbageCollectorSpec) {
	*out = *in
	if in.Collector != nil {
		in, out := &in.Collector, &out.Collector
		*out = new(string)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmGarbageCollectorSpec.
func (in *JvmGarbageCollectorSpec) DeepCopy() *JvmGarbageCollectorSpec {
	if in == nil {
		return nil
	}
	out := new(JvmGarbageCollectorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmJmxmpSpec) DeepCopyInto(out *JvmJmxmpSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmJmxmpSpec.
func (in *JvmJmxmpSpec) DeepCopy() *JvmJmxmpSpec {
	if in == nil {
		return nil
	}
	out := new(JvmJmxmpSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmMemorySpec) DeepCopyInto(out *JvmMemorySpec) {
	*out = *in
	if in.HeapSize != nil {
		in, out := &in.HeapSize, &out.HeapSize
		*out = new(string)
		**out = **in
	}
	if in.StackSize != nil {
		in, out := &in.StackSize, &out.StackSize
		*out = new(string)
		**out = **in
	}
	if in.MetaspaceSize != nil {
		in, out := &in.MetaspaceSize, &out.MetaspaceSize
		*out = new(string)
		**out = **in
	}
	if in.DirectMemorySize != nil {
		in, out := &in.DirectMemorySize, &out.DirectMemorySize
		*out = new(string)
		**out = **in
	}
	if in.NativeMemoryTracking != nil {
		in, out := &in.NativeMemoryTracking, &out.NativeMemoryTracking
		*out = new(string)
		**out = **in
	}
	if in.OnOutOfMemory != nil {
		in, out := &in.OnOutOfMemory, &out.OnOutOfMemory
		*out = new(JvmOutOfMemorySpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmMemorySpec.
func (in *JvmMemorySpec) DeepCopy() *JvmMemorySpec {
	if in == nil {
		return nil
	}
	out := new(JvmMemorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JvmOutOfMemorySpec) DeepCopyInto(out *JvmOutOfMemorySpec) {
	*out = *in
	if in.Exit != nil {
		in, out := &in.Exit, &out.Exit
		*out = new(bool)
		**out = **in
	}
	if in.HeapDump != nil {
		in, out := &in.HeapDump, &out.HeapDump
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JvmOutOfMemorySpec.
func (in *JvmOutOfMemorySpec) DeepCopy() *JvmOutOfMemorySpec {
	if in == nil {
		return nil
	}
	out := new(JvmOutOfMemorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectReference) DeepCopyInto(out *LocalObjectReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectReference.
func (in *LocalObjectReference) DeepCopy() *LocalObjectReference {
	if in == nil {
		return nil
	}
	out := new(LocalObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedPortSpec) DeepCopyInto(out *NamedPortSpec) {
	*out = *in
	if in.Protocol != nil {
		in, out := &in.Protocol, &out.Protocol
		*out = new(corev1.Protocol)
		**out = **in
	}
	if in.NodePort != nil {
		in, out := &in.NodePort, &out.NodePort
		*out = new(int32)
		**out = **in
	}
	if in.HostPort != nil {
		in, out := &in.HostPort, &out.HostPort
		*out = new(int32)
		**out = **in
	}
	if in.HostIP != nil {
		in, out := &in.HostIP, &out.HostIP
		*out = new(string)
		**out = **in
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceMonitor != nil {
		in, out := &in.ServiceMonitor, &out.ServiceMonitor
		*out = new(ServiceMonitorSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedPortSpec.
func (in *NamedPortSpec) DeepCopy() *NamedPortSpec {
	if in == nil {
		return nil
	}
	out := new(NamedPortSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkSpec) DeepCopyInto(out *NetworkSpec) {
	*out = *in
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = new(PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DNSPolicy != nil {
		in, out := &in.DNSPolicy, &out.DNSPolicy
		*out = new(corev1.DNSPolicy)
		**out = **in
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]corev1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.Hostname != nil {
		in, out := &in.Hostname, &out.Hostname
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkSpec.
func (in *NetworkSpec) DeepCopy() *NetworkSpec {
	if in == nil {
		return nil
	}
	out := new(NetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceSpec) DeepCopyInto(out *PersistenceSpec) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(string)
		**out = **in
	}
	in.PersistentStorageSpec.DeepCopyInto(&out.PersistentStorageSpec)
	if in.Snapshots != nil {
		in, out := &in.Snapshots, &out.Snapshots
		*out = new(PersistentStorageSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceSpec.
func (in *PersistenceSpec) DeepCopy() *PersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(PersistenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentStorageSpec) DeepCopyInto(out *PersistentStorageSpec) {
	*out = *in
	if in.PersistentVolumeClaim != nil {
		in, out := &in.PersistentVolumeClaim, &out.PersistentVolumeClaim
		*out = new(corev1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(corev1.VolumeSource)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentStorageSpec.
func (in *PersistentStorageSpec) DeepCopy() *PersistentStorageSpec {
	if in == nil {
		return nil
	}
	out := new(PersistentStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDNSConfig) DeepCopyInto(out *PodDNSConfig) {
	*out = *in
	if in.Nameservers != nil {
		in, out := &in.Nameservers, &out.Nameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Searches != nil {
		in, out := &in.Searches, &out.Searches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make([]corev1.PodDNSConfigOption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDNSConfig.
func (in *PodDNSConfig) DeepCopy() *PodDNSConfig {
	if in == nil {
		return nil
	}
	out := new(PodDNSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PortSpecWithSSL) DeepCopyInto(out *PortSpecWithSSL) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.SSL != nil {
		in, out := &in.SSL, &out.SSL
		*out = new(SSLSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PortSpecWithSSL.
func (in *PortSpecWithSSL) DeepCopy() *PortSpecWithSSL {
	if in == nil {
		return nil
	}
	out := new(PortSpecWithSSL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeHandler) DeepCopyInto(out *ProbeHandler) {
	*out = *in
	if in.Exec != nil {
		in, out := &in.Exec, &out.Exec
		*out = new(corev1.ExecAction)
		(*in).DeepCopyInto(*out)
	}
	if in.HTTPGet != nil {
		in, out := &in.HTTPGet, &out.HTTPGet
		*out = new(corev1.HTTPGetAction)
		(*in).DeepCopyInto(*out)
	}
	if in.TCPSocket != nil {
		in, out := &in.TCPSocket, &out.TCPSocket
		*out = new(corev1.TCPSocketAction)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeHandler.
func (in *ProbeHandler) DeepCopy() *ProbeHandler {
	if in == nil {
		return nil
	}
	out := new(ProbeHandler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadinessProbeSpec) DeepCopyInto(out *ReadinessProbeSpec) {
	*out = *in
	in.ProbeHandler.DeepCopyInto(&out.ProbeHandler)
	if in.InitialDelaySeconds != nil {
		in, out := &in.InitialDelaySeconds, &out.InitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	if in.PeriodSeconds != nil {
		in, out := &in.PeriodSeconds, &out.PeriodSeconds
		*out = new(int32)
		**out = **in
	}
	if in.SuccessThreshold != nil {
		in, out := &in.SuccessThreshold, &out.SuccessThreshold
		*out = new(int32)
		**out = **in
	}
	if in.FailureThreshold != nil {
		in, out := &in.FailureThreshold, &out.FailureThreshold
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadinessProbeSpec.
func (in *ReadinessProbeSpec) DeepCopy() *ReadinessProbeSpec {
	if in == nil {
		return nil
	}
	out := new(ReadinessProbeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	if in.Spec != nil {
		out.Spec = in.Spec.DeepCopyObject()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resources) DeepCopyInto(out *Resources) {
	*out = *in
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Resource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resources.
func (in *Resources) DeepCopy() *Resources {
	if in == nil {
		return nil
	}
	out := new(Resources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLSpec) DeepCopyInto(out *SSLSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = new(string)
		**out = **in
	}
	if in.KeyStore != nil {
		in, out := &in.KeyStore, &out.KeyStore
		*out = new(string)
		**out = **in
	}
	if in.KeyStorePasswordFile != nil {
		in, out := &in.KeyStorePasswordFile, &out.KeyStorePasswordFile
		*out = new(string)
		**out = **in
	}
	if in.KeyPasswordFile != nil {
		in, out := &in.KeyPasswordFile, &out.KeyPasswordFile
		*out = new(string)
		**out = **in
	}
	if in.KeyStoreAlgorithm != nil {
		in, out := &in.KeyStoreAlgorithm, &out.KeyStoreAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.KeyStoreProvider != nil {
		in, out := &in.KeyStoreProvider, &out.KeyStoreProvider
		*out = new(string)
		**out = **in
	}
	if in.KeyStoreType != nil {
		in, out := &in.KeyStoreType, &out.KeyStoreType
		*out = new(string)
		**out = **in
	}
	if in.TrustStore != nil {
		in, out := &in.TrustStore, &out.TrustStore
		*out = new(string)
		**out = **in
	}
	if in.TrustStorePasswordFile != nil {
		in, out := &in.TrustStorePasswordFile, &out.TrustStorePasswordFile
		*out = new(string)
		**out = **in
	}
	if in.TrustStoreAlgorithm != nil {
		in, out := &in.TrustStoreAlgorithm, &out.TrustStoreAlgorithm
		*out = new(string)
		**out = **in
	}
	if in.TrustStoreProvider != nil {
		in, out := &in.TrustStoreProvider, &out.TrustStoreProvider
		*out = new(string)
		**out = **in
	}
	if in.TrustStoreType != nil {
		in, out := &in.TrustStoreType, &out.TrustStoreType
		*out = new(string)
		**out = **in
	}
	if in.RequireClientCert != nil {
		in, out := &in.RequireClientCert, &out.RequireClientCert
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLSpec.
func (in *SSLSpec) DeepCopy() *SSLSpec {
	if in == nil {
		return nil
	}
	out := new(SSLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingProbe) DeepCopyInto(out *ScalingProbe) {
	*out = *in
	in.Handler.DeepCopyInto(&out.Handler)
	if in.TimeoutSeconds != nil {
		in, out := &in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingProbe.
func (in *ScalingProbe) DeepCopy() *ScalingProbe {
	if in == nil {
		return nil
	}
	out := new(ScalingProbe)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingSpec) DeepCopyInto(out *ScalingSpec) {
	*out = *in
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(ScalingPolicy)
		**out = **in
	}
	if in.Probe != nil {
		in, out := &in.Probe, &out.Probe
		*out = new(ScalingProbe)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingSpec.
func (in *ScalingSpec) DeepCopy() *ScalingSpec {
	if in == nil {
		return nil
	}
	out := new(ScalingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretVolumeSpec) DeepCopyInto(out *SecretVolumeSpec) {
	*out = *in
	if in.MountPropagation != nil {
		in, out := &in.MountPropagation, &out.MountPropagation
		*out = new(corev1.MountPropagationMode)
		**out = **in
	}
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]corev1.KeyToPath, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultMode != nil {
		in, out := &in.DefaultMode, &out.DefaultMode
		*out = new(int32)
		**out = **in
	}
	if in.Optional != nil {
		in, out := &in.Optional, &out.Optional
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretVolumeSpec.
func (in *SecretVolumeSpec) DeepCopy() *SecretVolumeSpec {
	if in == nil {
		return nil
	}
	out := new(SecretVolumeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceMonitorSpec) DeepCopyInto(out *ServiceMonitorSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetLabels != nil {
		in, out := &in.TargetLabels, &out.TargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodTargetLabels != nil {
		in, out := &in.PodTargetLabels, &out.PodTargetLabels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Params != nil {
		in, out := &in.Params, &out.Params
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.TLSConfig != nil {
		in, out := &in.TLSConfig, &out.TLSConfig
		*out = new(monitoringv1.TLSConfig)
		(*in).DeepCopyInto(*out)
	}
	in.BearerTokenSecret.DeepCopyInto(&out.BearerTokenSecret)
	if in.HonorTimestamps != nil {
		in, out := &in.HonorTimestamps, &out.HonorTimestamps
		*out = new(bool)
		**out = **in
	}
	if in.BasicAuth != nil {
		in, out := &in.BasicAuth, &out.BasicAuth
		*out = new(monitoringv1.BasicAuth)
		(*in).DeepCopyInto(*out)
	}
	if in.MetricRelabelings != nil {
		in, out := &in.MetricRelabelings, &out.MetricRelabelings
		*out = make([]*monitoringv1.RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(monitoringv1.RelabelConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.Relabelings != nil {
		in, out := &in.Relabelings, &out.Relabelings
		*out = make([]*monitoringv1.RelabelConfig, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(monitoringv1.RelabelConfig)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.ProxyURL != nil {
		in, out := &in.ProxyURL, &out.ProxyURL
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceMonitorSpec.
func (in *ServiceMonitorSpec) DeepCopy() *ServiceMonitorSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceMonitorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(corev1.ServiceType)
		**out = **in
	}
	if in.ExternalIPs != nil {
		in, out := &in.ExternalIPs, &out.ExternalIPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterIP != nil {
		in, out := &in.ClusterIP, &out.ClusterIP
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerIP != nil {
		in, out := &in.LoadBalancerIP, &out.LoadBalancerIP
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SessionAffinity != nil {
		in, out := &in.SessionAffinity, &out.SessionAffinity
		*out = new(corev1.ServiceAffinity)
		**out = **in
	}
	if in.LoadBalancerSourceRanges != nil {
		in, out := &in.LoadBalancerSourceRanges, &out.LoadBalancerSourceRanges
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalName != nil {
		in, out := &in.ExternalName, &out.ExternalName
		*out = new(string)
		**out = **in
	}
	if in.ExternalTrafficPolicy != nil {
		in, out := &in.ExternalTrafficPolicy, &out.ExternalTrafficPolicy
		*out = new(corev1.ServiceExternalTrafficPolicyType)
		**out = **in
	}
	if in.HealthCheckNodePort != nil {
		in, out := &in.HealthCheckNodePort, &out.HealthCheckNodePort
		*out = new(int32)
		**out = **in
	}
	if in.PublishNotReadyAddresses != nil {
		in, out := &in.PublishNotReadyAddresses, &out.PublishNotReadyAddresses
		*out = new(bool)
		**out = **in
	}
	if in.SessionAffinityConfig != nil {
		in, out := &in.SessionAffinityConfig, &out.SessionAffinityConfig
		*out = new(corev1.SessionAffinityConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.IPFamily != nil {
		in, out := &in.IPFamily, &out.IPFamily
		*out = new(corev1.IPFamily)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartQuorum) DeepCopyInto(out *StartQuorum) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartQuorum.
func (in *StartQuorum) DeepCopy() *StartQuorum {
	if in == nil {
		return nil
	}
	out := new(StartQuorum)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StartQuorumStatus) DeepCopyInto(out *StartQuorumStatus) {
	*out = *in
	out.StartQuorum = in.StartQuorum
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StartQuorumStatus.
func (in *StartQuorumStatus) DeepCopy() *StartQuorumStatus {
	if in == nil {
		return nil
	}
	out := new(StartQuorumStatus)
	in.DeepCopyInto(out)
	return out
}
