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

package resource

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Quantity) DeepCopyInto(out *Quantity) {
	*out = in.DeepCopy()
	return
}
