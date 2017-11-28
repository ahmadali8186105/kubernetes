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

// This file was automatically generated by: _output/bin/conversion-gen --extra-peer-dirs k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/api/core/v1 --v 1 --logtostderr -i k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission/v1beta1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1alpha1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1beta2,k8s.io/kubernetes/pkg/apis/authentication/v1,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization/v1,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/autoscaling/v2beta1,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v1beta1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates/v1beta1,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/kubernetes/pkg/apis/events/v1beta1,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/networking/v1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac/v1,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac/v1beta1,k8s.io/kubernetes/pkg/apis/scheduling/v1alpha1,k8s.io/kubernetes/pkg/apis/settings/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1,k8s.io/kubernetes/pkg/apis/storage/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1beta1,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/appsv1beta1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/appsv1beta2,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/autoscalingv1,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme/extensionsv1beta1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1 -O zz_generated.conversion

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	v1alpha1 "k8s.io/api/scheduling/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	scheduling "k8s.io/kubernetes/pkg/apis/scheduling"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_PriorityClass_To_scheduling_PriorityClass,
		Convert_scheduling_PriorityClass_To_v1alpha1_PriorityClass,
		Convert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList,
		Convert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList,
	)
}

func autoConvert_v1alpha1_PriorityClass_To_scheduling_PriorityClass(in *v1alpha1.PriorityClass, out *scheduling.PriorityClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Value = in.Value
	out.GlobalDefault = in.GlobalDefault
	out.Description = in.Description
	return nil
}

// Convert_v1alpha1_PriorityClass_To_scheduling_PriorityClass is an autogenerated conversion function.
func Convert_v1alpha1_PriorityClass_To_scheduling_PriorityClass(in *v1alpha1.PriorityClass, out *scheduling.PriorityClass, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityClass_To_scheduling_PriorityClass(in, out, s)
}

func autoConvert_scheduling_PriorityClass_To_v1alpha1_PriorityClass(in *scheduling.PriorityClass, out *v1alpha1.PriorityClass, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Value = in.Value
	out.GlobalDefault = in.GlobalDefault
	out.Description = in.Description
	return nil
}

// Convert_scheduling_PriorityClass_To_v1alpha1_PriorityClass is an autogenerated conversion function.
func Convert_scheduling_PriorityClass_To_v1alpha1_PriorityClass(in *scheduling.PriorityClass, out *v1alpha1.PriorityClass, s conversion.Scope) error {
	return autoConvert_scheduling_PriorityClass_To_v1alpha1_PriorityClass(in, out, s)
}

func autoConvert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList(in *v1alpha1.PriorityClassList, out *scheduling.PriorityClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]scheduling.PriorityClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList is an autogenerated conversion function.
func Convert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList(in *v1alpha1.PriorityClassList, out *scheduling.PriorityClassList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PriorityClassList_To_scheduling_PriorityClassList(in, out, s)
}

func autoConvert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList(in *scheduling.PriorityClassList, out *v1alpha1.PriorityClassList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1alpha1.PriorityClass)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList is an autogenerated conversion function.
func Convert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList(in *scheduling.PriorityClassList, out *v1alpha1.PriorityClassList, s conversion.Scope) error {
	return autoConvert_scheduling_PriorityClassList_To_v1alpha1_PriorityClassList(in, out, s)
}
