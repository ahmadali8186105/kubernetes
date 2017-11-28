// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was automatically generated by: _output/bin/deepcopy-gen --v 1 --logtostderr -i k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm,k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/cmd/kubeadm/app/phases/etcd/spec,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission,k8s.io/kubernetes/pkg/apis/admissionregistration,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/networking,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/scheduling,k8s.io/kubernetes/pkg/apis/settings,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/controller/garbagecollector/metaonly,k8s.io/kubernetes/pkg/kubectl/cmd/testing,k8s.io/kubernetes/pkg/kubectl/testing,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/pkg/registry/rbac/reconciliation,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/plugin/pkg/scheduler/api,k8s.io/kubernetes/plugin/pkg/scheduler/api/v1,k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/admissionregistration/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/admissionregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1beta2,k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1,k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/authorization/v1,k8s.io/kubernetes/vendor/k8s.io/api/authorization/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/autoscaling/v1,k8s.io/kubernetes/vendor/k8s.io/api/autoscaling/v2beta1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v2alpha1,k8s.io/kubernetes/vendor/k8s.io/api/certificates/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/core/v1,k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/extensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/imagepolicy/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/networking/v1,k8s.io/kubernetes/vendor/k8s.io/api/policy/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/settings/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/labels,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/testing,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/testing,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/test,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/watch,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/endpoints/openapi/testing,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/endpoints/testing,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/registry/rest,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/storage/testing,k8s.io/kubernetes/vendor/k8s.io/client-go/rest,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme,k8s.io/kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api,k8s.io/kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/crd/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/crd/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1 --bounding-dirs k8s.io/kubernetes,k8s.io/api -O zz_generated.deepcopy

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthInfo) DeepCopyInto(out *AuthInfo) {
	*out = *in
	if in.ClientCertificateData != nil {
		in, out := &in.ClientCertificateData, &out.ClientCertificateData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ClientKeyData != nil {
		in, out := &in.ClientKeyData, &out.ClientKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateGroups != nil {
		in, out := &in.ImpersonateGroups, &out.ImpersonateGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateUserExtra != nil {
		in, out := &in.ImpersonateUserExtra, &out.ImpersonateUserExtra
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = make([]string, len(val))
				copy((*out)[key], val)
			}
		}
	}
	if in.AuthProvider != nil {
		in, out := &in.AuthProvider, &out.AuthProvider
		if *in == nil {
			*out = nil
		} else {
			*out = new(AuthProviderConfig)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthInfo.
func (in *AuthInfo) DeepCopy() *AuthInfo {
	if in == nil {
		return nil
	}
	out := new(AuthInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProviderConfig) DeepCopyInto(out *AuthProviderConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProviderConfig.
func (in *AuthProviderConfig) DeepCopy() *AuthProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AuthProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cluster) DeepCopyInto(out *Cluster) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cluster.
func (in *Cluster) DeepCopy() *Cluster {
	if in == nil {
		return nil
	}
	out := new(Cluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	in.Preferences.DeepCopyInto(&out.Preferences)
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]NamedCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AuthInfos != nil {
		in, out := &in.AuthInfos, &out.AuthInfos
		*out = make([]NamedAuthInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Contexts != nil {
		in, out := &in.Contexts, &out.Contexts
		*out = make([]NamedContext, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Config) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedAuthInfo) DeepCopyInto(out *NamedAuthInfo) {
	*out = *in
	in.AuthInfo.DeepCopyInto(&out.AuthInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedAuthInfo.
func (in *NamedAuthInfo) DeepCopy() *NamedAuthInfo {
	if in == nil {
		return nil
	}
	out := new(NamedAuthInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedCluster) DeepCopyInto(out *NamedCluster) {
	*out = *in
	in.Cluster.DeepCopyInto(&out.Cluster)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedCluster.
func (in *NamedCluster) DeepCopy() *NamedCluster {
	if in == nil {
		return nil
	}
	out := new(NamedCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedContext) DeepCopyInto(out *NamedContext) {
	*out = *in
	in.Context.DeepCopyInto(&out.Context)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedContext.
func (in *NamedContext) DeepCopy() *NamedContext {
	if in == nil {
		return nil
	}
	out := new(NamedContext)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedExtension) DeepCopyInto(out *NamedExtension) {
	*out = *in
	in.Extension.DeepCopyInto(&out.Extension)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedExtension.
func (in *NamedExtension) DeepCopy() *NamedExtension {
	if in == nil {
		return nil
	}
	out := new(NamedExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Preferences) DeepCopyInto(out *Preferences) {
	*out = *in
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = make([]NamedExtension, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Preferences.
func (in *Preferences) DeepCopy() *Preferences {
	if in == nil {
		return nil
	}
	out := new(Preferences)
	in.DeepCopyInto(out)
	return out
}
