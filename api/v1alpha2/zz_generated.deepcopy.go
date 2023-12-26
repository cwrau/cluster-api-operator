//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/component-base/config/v1alpha1"
	"sigs.k8s.io/cluster-api/api/v1beta1"
	timex "time"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonProvider) DeepCopyInto(out *AddonProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonProvider.
func (in *AddonProvider) DeepCopy() *AddonProvider {
	if in == nil {
		return nil
	}
	out := new(AddonProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonProviderList) DeepCopyInto(out *AddonProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AddonProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonProviderList.
func (in *AddonProviderList) DeepCopy() *AddonProviderList {
	if in == nil {
		return nil
	}
	out := new(AddonProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonProviderSpec) DeepCopyInto(out *AddonProviderSpec) {
	*out = *in
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonProviderSpec.
func (in *AddonProviderSpec) DeepCopy() *AddonProviderSpec {
	if in == nil {
		return nil
	}
	out := new(AddonProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonProviderStatus) DeepCopyInto(out *AddonProviderStatus) {
	*out = *in
	in.ProviderStatus.DeepCopyInto(&out.ProviderStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonProviderStatus.
func (in *AddonProviderStatus) DeepCopy() *AddonProviderStatus {
	if in == nil {
		return nil
	}
	out := new(AddonProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapProvider) DeepCopyInto(out *BootstrapProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapProvider.
func (in *BootstrapProvider) DeepCopy() *BootstrapProvider {
	if in == nil {
		return nil
	}
	out := new(BootstrapProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BootstrapProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapProviderList) DeepCopyInto(out *BootstrapProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BootstrapProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapProviderList.
func (in *BootstrapProviderList) DeepCopy() *BootstrapProviderList {
	if in == nil {
		return nil
	}
	out := new(BootstrapProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BootstrapProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapProviderSpec) DeepCopyInto(out *BootstrapProviderSpec) {
	*out = *in
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapProviderSpec.
func (in *BootstrapProviderSpec) DeepCopy() *BootstrapProviderSpec {
	if in == nil {
		return nil
	}
	out := new(BootstrapProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootstrapProviderStatus) DeepCopyInto(out *BootstrapProviderStatus) {
	*out = *in
	in.ProviderStatus.DeepCopyInto(&out.ProviderStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootstrapProviderStatus.
func (in *BootstrapProviderStatus) DeepCopy() *BootstrapProviderStatus {
	if in == nil {
		return nil
	}
	out := new(BootstrapProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigmapReference) DeepCopyInto(out *ConfigmapReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigmapReference.
func (in *ConfigmapReference) DeepCopy() *ConfigmapReference {
	if in == nil {
		return nil
	}
	out := new(ConfigmapReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerSpec) DeepCopyInto(out *ContainerSpec) {
	*out = *in
	if in.ImageURL != nil {
		in, out := &in.ImageURL, &out.ImageURL
		*out = new(string)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerSpec.
func (in *ContainerSpec) DeepCopy() *ContainerSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneProvider) DeepCopyInto(out *ControlPlaneProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneProvider.
func (in *ControlPlaneProvider) DeepCopy() *ControlPlaneProvider {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneProviderList) DeepCopyInto(out *ControlPlaneProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ControlPlaneProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneProviderList.
func (in *ControlPlaneProviderList) DeepCopy() *ControlPlaneProviderList {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControlPlaneProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneProviderSpec) DeepCopyInto(out *ControlPlaneProviderSpec) {
	*out = *in
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneProviderSpec.
func (in *ControlPlaneProviderSpec) DeepCopy() *ControlPlaneProviderSpec {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControlPlaneProviderStatus) DeepCopyInto(out *ControlPlaneProviderStatus) {
	*out = *in
	in.ProviderStatus.DeepCopyInto(&out.ProviderStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControlPlaneProviderStatus.
func (in *ControlPlaneProviderStatus) DeepCopy() *ControlPlaneProviderStatus {
	if in == nil {
		return nil
	}
	out := new(ControlPlaneProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfigurationSpec) DeepCopyInto(out *ControllerConfigurationSpec) {
	*out = *in
	if in.GroupKindConcurrency != nil {
		in, out := &in.GroupKindConcurrency, &out.GroupKindConcurrency
		*out = make(map[string]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CacheSyncTimeout != nil {
		in, out := &in.CacheSyncTimeout, &out.CacheSyncTimeout
		*out = new(timex.Duration)
		**out = **in
	}
	if in.RecoverPanic != nil {
		in, out := &in.RecoverPanic, &out.RecoverPanic
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfigurationSpec.
func (in *ControllerConfigurationSpec) DeepCopy() *ControllerConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerHealth) DeepCopyInto(out *ControllerHealth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerHealth.
func (in *ControllerHealth) DeepCopy() *ControllerHealth {
	if in == nil {
		return nil
	}
	out := new(ControllerHealth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerConfiguration) DeepCopyInto(out *ControllerManagerConfiguration) {
	*out = *in
	if in.SyncPeriod != nil {
		in, out := &in.SyncPeriod, &out.SyncPeriod
		*out = new(v1.Duration)
		**out = **in
	}
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(v1alpha1.LeaderElectionConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.GracefulShutdownTimeout != nil {
		in, out := &in.GracefulShutdownTimeout, &out.GracefulShutdownTimeout
		*out = new(v1.Duration)
		**out = **in
	}
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		*out = new(ControllerConfigurationSpec)
		(*in).DeepCopyInto(*out)
	}
	out.Metrics = in.Metrics
	out.Health = in.Health
	in.Webhook.DeepCopyInto(&out.Webhook)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerConfiguration.
func (in *ControllerManagerConfiguration) DeepCopy() *ControllerManagerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerMetrics) DeepCopyInto(out *ControllerMetrics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerMetrics.
func (in *ControllerMetrics) DeepCopy() *ControllerMetrics {
	if in == nil {
		return nil
	}
	out := new(ControllerMetrics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerWebhook) DeepCopyInto(out *ControllerWebhook) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerWebhook.
func (in *ControllerWebhook) DeepCopy() *ControllerWebhook {
	if in == nil {
		return nil
	}
	out := new(ControllerWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreProvider) DeepCopyInto(out *CoreProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreProvider.
func (in *CoreProvider) DeepCopy() *CoreProvider {
	if in == nil {
		return nil
	}
	out := new(CoreProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoreProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreProviderList) DeepCopyInto(out *CoreProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CoreProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreProviderList.
func (in *CoreProviderList) DeepCopy() *CoreProviderList {
	if in == nil {
		return nil
	}
	out := new(CoreProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CoreProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreProviderSpec) DeepCopyInto(out *CoreProviderSpec) {
	*out = *in
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreProviderSpec.
func (in *CoreProviderSpec) DeepCopy() *CoreProviderSpec {
	if in == nil {
		return nil
	}
	out := new(CoreProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreProviderStatus) DeepCopyInto(out *CoreProviderStatus) {
	*out = *in
	in.ProviderStatus.DeepCopyInto(&out.ProviderStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreProviderStatus.
func (in *CoreProviderStatus) DeepCopy() *CoreProviderStatus {
	if in == nil {
		return nil
	}
	out := new(CoreProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentSpec) DeepCopyInto(out *DeploymentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
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
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]ContainerSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]corev1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentSpec.
func (in *DeploymentSpec) DeepCopy() *DeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(DeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FetchConfiguration) DeepCopyInto(out *FetchConfiguration) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FetchConfiguration.
func (in *FetchConfiguration) DeepCopy() *FetchConfiguration {
	if in == nil {
		return nil
	}
	out := new(FetchConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMProvider) DeepCopyInto(out *IPAMProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMProvider.
func (in *IPAMProvider) DeepCopy() *IPAMProvider {
	if in == nil {
		return nil
	}
	out := new(IPAMProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMProviderList) DeepCopyInto(out *IPAMProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IPAMProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMProviderList.
func (in *IPAMProviderList) DeepCopy() *IPAMProviderList {
	if in == nil {
		return nil
	}
	out := new(IPAMProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IPAMProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMProviderSpec) DeepCopyInto(out *IPAMProviderSpec) {
	*out = *in
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMProviderSpec.
func (in *IPAMProviderSpec) DeepCopy() *IPAMProviderSpec {
	if in == nil {
		return nil
	}
	out := new(IPAMProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPAMProviderStatus) DeepCopyInto(out *IPAMProviderStatus) {
	*out = *in
	in.ProviderStatus.DeepCopyInto(&out.ProviderStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPAMProviderStatus.
func (in *IPAMProviderStatus) DeepCopy() *IPAMProviderStatus {
	if in == nil {
		return nil
	}
	out := new(IPAMProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureProvider) DeepCopyInto(out *InfrastructureProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureProvider.
func (in *InfrastructureProvider) DeepCopy() *InfrastructureProvider {
	if in == nil {
		return nil
	}
	out := new(InfrastructureProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureProviderList) DeepCopyInto(out *InfrastructureProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfrastructureProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureProviderList.
func (in *InfrastructureProviderList) DeepCopy() *InfrastructureProviderList {
	if in == nil {
		return nil
	}
	out := new(InfrastructureProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfrastructureProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureProviderSpec) DeepCopyInto(out *InfrastructureProviderSpec) {
	*out = *in
	in.ProviderSpec.DeepCopyInto(&out.ProviderSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureProviderSpec.
func (in *InfrastructureProviderSpec) DeepCopy() *InfrastructureProviderSpec {
	if in == nil {
		return nil
	}
	out := new(InfrastructureProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfrastructureProviderStatus) DeepCopyInto(out *InfrastructureProviderStatus) {
	*out = *in
	in.ProviderStatus.DeepCopyInto(&out.ProviderStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfrastructureProviderStatus.
func (in *InfrastructureProviderStatus) DeepCopy() *InfrastructureProviderStatus {
	if in == nil {
		return nil
	}
	out := new(InfrastructureProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagerSpec) DeepCopyInto(out *ManagerSpec) {
	*out = *in
	in.ControllerManagerConfiguration.DeepCopyInto(&out.ControllerManagerConfiguration)
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagerSpec.
func (in *ManagerSpec) DeepCopy() *ManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderSpec) DeepCopyInto(out *ProviderSpec) {
	*out = *in
	if in.Manager != nil {
		in, out := &in.Manager, &out.Manager
		*out = new(ManagerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(DeploymentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigSecret != nil {
		in, out := &in.ConfigSecret, &out.ConfigSecret
		*out = new(SecretReference)
		**out = **in
	}
	if in.FetchConfig != nil {
		in, out := &in.FetchConfig, &out.FetchConfig
		*out = new(FetchConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalManifestsRef != nil {
		in, out := &in.AdditionalManifestsRef, &out.AdditionalManifestsRef
		*out = new(ConfigmapReference)
		**out = **in
	}
	if in.ManifestPatches != nil {
		in, out := &in.ManifestPatches, &out.ManifestPatches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderSpec.
func (in *ProviderSpec) DeepCopy() *ProviderSpec {
	if in == nil {
		return nil
	}
	out := new(ProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProviderStatus) DeepCopyInto(out *ProviderStatus) {
	*out = *in
	if in.Contract != nil {
		in, out := &in.Contract, &out.Contract
		*out = new(string)
		**out = **in
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(v1beta1.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstalledVersion != nil {
		in, out := &in.InstalledVersion, &out.InstalledVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProviderStatus.
func (in *ProviderStatus) DeepCopy() *ProviderStatus {
	if in == nil {
		return nil
	}
	out := new(ProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}
