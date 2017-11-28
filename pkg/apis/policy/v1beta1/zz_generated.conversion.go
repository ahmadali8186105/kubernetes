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

package v1beta1

import (
	v1beta1 "k8s.io/api/policy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
	policy "k8s.io/kubernetes/pkg/apis/policy"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_Eviction_To_policy_Eviction,
		Convert_policy_Eviction_To_v1beta1_Eviction,
		Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget,
		Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget,
		Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList,
		Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList,
		Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec,
		Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec,
		Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus,
		Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus,
	)
}

func autoConvert_v1beta1_Eviction_To_policy_Eviction(in *v1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

// Convert_v1beta1_Eviction_To_policy_Eviction is an autogenerated conversion function.
func Convert_v1beta1_Eviction_To_policy_Eviction(in *v1beta1.Eviction, out *policy.Eviction, s conversion.Scope) error {
	return autoConvert_v1beta1_Eviction_To_policy_Eviction(in, out, s)
}

func autoConvert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *v1beta1.Eviction, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.DeleteOptions = (*v1.DeleteOptions)(unsafe.Pointer(in.DeleteOptions))
	return nil
}

// Convert_policy_Eviction_To_v1beta1_Eviction is an autogenerated conversion function.
func Convert_policy_Eviction_To_v1beta1_Eviction(in *policy.Eviction, out *v1beta1.Eviction, s conversion.Scope) error {
	return autoConvert_policy_Eviction_To_v1beta1_Eviction(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *v1beta1.PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in *v1beta1.PodDisruptionBudget, out *policy.PodDisruptionBudget, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudget_To_policy_PodDisruptionBudget(in, out, s)
}

func autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *v1beta1.PodDisruptionBudget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in *policy.PodDisruptionBudget, out *v1beta1.PodDisruptionBudget, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudget_To_v1beta1_PodDisruptionBudget(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *v1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]policy.PodDisruptionBudget)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in *v1beta1.PodDisruptionBudgetList, out *policy.PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetList_To_policy_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *v1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.PodDisruptionBudget)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in *policy.PodDisruptionBudgetList, out *v1beta1.PodDisruptionBudgetList, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetList_To_v1beta1_PodDisruptionBudgetList(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *v1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in *v1beta1.PodDisruptionBudgetSpec, out *policy.PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetSpec_To_policy_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *v1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
	out.MinAvailable = (*intstr.IntOrString)(unsafe.Pointer(in.MinAvailable))
	out.Selector = (*v1.LabelSelector)(unsafe.Pointer(in.Selector))
	out.MaxUnavailable = (*intstr.IntOrString)(unsafe.Pointer(in.MaxUnavailable))
	return nil
}

// Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in *policy.PodDisruptionBudgetSpec, out *v1beta1.PodDisruptionBudgetSpec, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetSpec_To_v1beta1_PodDisruptionBudgetSpec(in, out, s)
}

func autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *v1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	return nil
}

// Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus is an autogenerated conversion function.
func Convert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in *v1beta1.PodDisruptionBudgetStatus, out *policy.PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_PodDisruptionBudgetStatus_To_policy_PodDisruptionBudgetStatus(in, out, s)
}

func autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *v1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.DisruptedPods = *(*map[string]v1.Time)(unsafe.Pointer(&in.DisruptedPods))
	out.PodDisruptionsAllowed = in.PodDisruptionsAllowed
	out.CurrentHealthy = in.CurrentHealthy
	out.DesiredHealthy = in.DesiredHealthy
	out.ExpectedPods = in.ExpectedPods
	return nil
}

// Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus is an autogenerated conversion function.
func Convert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in *policy.PodDisruptionBudgetStatus, out *v1beta1.PodDisruptionBudgetStatus, s conversion.Scope) error {
	return autoConvert_policy_PodDisruptionBudgetStatus_To_v1beta1_PodDisruptionBudgetStatus(in, out, s)
}
