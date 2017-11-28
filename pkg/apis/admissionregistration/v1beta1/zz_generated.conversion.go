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
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	admissionregistration "k8s.io/kubernetes/pkg/apis/admissionregistration"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration,
		Convert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration,
		Convert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList,
		Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList,
		Convert_v1beta1_Rule_To_admissionregistration_Rule,
		Convert_admissionregistration_Rule_To_v1beta1_Rule,
		Convert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations,
		Convert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations,
		Convert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference,
		Convert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference,
		Convert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration,
		Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration,
		Convert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList,
		Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList,
		Convert_v1beta1_Webhook_To_admissionregistration_Webhook,
		Convert_admissionregistration_Webhook_To_v1beta1_Webhook,
		Convert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig,
		Convert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig,
	)
}

func autoConvert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in *v1beta1.MutatingWebhookConfiguration, out *admissionregistration.MutatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Webhooks = *(*[]admissionregistration.Webhook)(unsafe.Pointer(&in.Webhooks))
	return nil
}

// Convert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration is an autogenerated conversion function.
func Convert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in *v1beta1.MutatingWebhookConfiguration, out *admissionregistration.MutatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_MutatingWebhookConfiguration_To_admissionregistration_MutatingWebhookConfiguration(in, out, s)
}

func autoConvert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration(in *admissionregistration.MutatingWebhookConfiguration, out *v1beta1.MutatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Webhooks = *(*[]v1beta1.Webhook)(unsafe.Pointer(&in.Webhooks))
	return nil
}

// Convert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration is an autogenerated conversion function.
func Convert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration(in *admissionregistration.MutatingWebhookConfiguration, out *v1beta1.MutatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_admissionregistration_MutatingWebhookConfiguration_To_v1beta1_MutatingWebhookConfiguration(in, out, s)
}

func autoConvert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in *v1beta1.MutatingWebhookConfigurationList, out *admissionregistration.MutatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]admissionregistration.MutatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in *v1beta1.MutatingWebhookConfigurationList, out *admissionregistration.MutatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_v1beta1_MutatingWebhookConfigurationList_To_admissionregistration_MutatingWebhookConfigurationList(in, out, s)
}

func autoConvert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList(in *admissionregistration.MutatingWebhookConfigurationList, out *v1beta1.MutatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.MutatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList(in *admissionregistration.MutatingWebhookConfigurationList, out *v1beta1.MutatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_admissionregistration_MutatingWebhookConfigurationList_To_v1beta1_MutatingWebhookConfigurationList(in, out, s)
}

func autoConvert_v1beta1_Rule_To_admissionregistration_Rule(in *v1beta1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_v1beta1_Rule_To_admissionregistration_Rule is an autogenerated conversion function.
func Convert_v1beta1_Rule_To_admissionregistration_Rule(in *v1beta1.Rule, out *admissionregistration.Rule, s conversion.Scope) error {
	return autoConvert_v1beta1_Rule_To_admissionregistration_Rule(in, out, s)
}

func autoConvert_admissionregistration_Rule_To_v1beta1_Rule(in *admissionregistration.Rule, out *v1beta1.Rule, s conversion.Scope) error {
	out.APIGroups = *(*[]string)(unsafe.Pointer(&in.APIGroups))
	out.APIVersions = *(*[]string)(unsafe.Pointer(&in.APIVersions))
	out.Resources = *(*[]string)(unsafe.Pointer(&in.Resources))
	return nil
}

// Convert_admissionregistration_Rule_To_v1beta1_Rule is an autogenerated conversion function.
func Convert_admissionregistration_Rule_To_v1beta1_Rule(in *admissionregistration.Rule, out *v1beta1.Rule, s conversion.Scope) error {
	return autoConvert_admissionregistration_Rule_To_v1beta1_Rule(in, out, s)
}

func autoConvert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *v1beta1.RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
	out.Operations = *(*[]admissionregistration.OperationType)(unsafe.Pointer(&in.Operations))
	if err := Convert_v1beta1_Rule_To_admissionregistration_Rule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations is an autogenerated conversion function.
func Convert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in *v1beta1.RuleWithOperations, out *admissionregistration.RuleWithOperations, s conversion.Scope) error {
	return autoConvert_v1beta1_RuleWithOperations_To_admissionregistration_RuleWithOperations(in, out, s)
}

func autoConvert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *v1beta1.RuleWithOperations, s conversion.Scope) error {
	out.Operations = *(*[]v1beta1.OperationType)(unsafe.Pointer(&in.Operations))
	if err := Convert_admissionregistration_Rule_To_v1beta1_Rule(&in.Rule, &out.Rule, s); err != nil {
		return err
	}
	return nil
}

// Convert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations is an autogenerated conversion function.
func Convert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations(in *admissionregistration.RuleWithOperations, out *v1beta1.RuleWithOperations, s conversion.Scope) error {
	return autoConvert_admissionregistration_RuleWithOperations_To_v1beta1_RuleWithOperations(in, out, s)
}

func autoConvert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference(in *v1beta1.ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Path = (*string)(unsafe.Pointer(in.Path))
	return nil
}

// Convert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference is an autogenerated conversion function.
func Convert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference(in *v1beta1.ServiceReference, out *admissionregistration.ServiceReference, s conversion.Scope) error {
	return autoConvert_v1beta1_ServiceReference_To_admissionregistration_ServiceReference(in, out, s)
}

func autoConvert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference(in *admissionregistration.ServiceReference, out *v1beta1.ServiceReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	out.Path = (*string)(unsafe.Pointer(in.Path))
	return nil
}

// Convert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference is an autogenerated conversion function.
func Convert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference(in *admissionregistration.ServiceReference, out *v1beta1.ServiceReference, s conversion.Scope) error {
	return autoConvert_admissionregistration_ServiceReference_To_v1beta1_ServiceReference(in, out, s)
}

func autoConvert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in *v1beta1.ValidatingWebhookConfiguration, out *admissionregistration.ValidatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Webhooks = *(*[]admissionregistration.Webhook)(unsafe.Pointer(&in.Webhooks))
	return nil
}

// Convert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration is an autogenerated conversion function.
func Convert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in *v1beta1.ValidatingWebhookConfiguration, out *admissionregistration.ValidatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_v1beta1_ValidatingWebhookConfiguration_To_admissionregistration_ValidatingWebhookConfiguration(in, out, s)
}

func autoConvert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration(in *admissionregistration.ValidatingWebhookConfiguration, out *v1beta1.ValidatingWebhookConfiguration, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Webhooks = *(*[]v1beta1.Webhook)(unsafe.Pointer(&in.Webhooks))
	return nil
}

// Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration(in *admissionregistration.ValidatingWebhookConfiguration, out *v1beta1.ValidatingWebhookConfiguration, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingWebhookConfiguration_To_v1beta1_ValidatingWebhookConfiguration(in, out, s)
}

func autoConvert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in *v1beta1.ValidatingWebhookConfigurationList, out *admissionregistration.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]admissionregistration.ValidatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in *v1beta1.ValidatingWebhookConfigurationList, out *admissionregistration.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_v1beta1_ValidatingWebhookConfigurationList_To_admissionregistration_ValidatingWebhookConfigurationList(in, out, s)
}

func autoConvert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList(in *admissionregistration.ValidatingWebhookConfigurationList, out *v1beta1.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.ValidatingWebhookConfiguration)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList is an autogenerated conversion function.
func Convert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList(in *admissionregistration.ValidatingWebhookConfigurationList, out *v1beta1.ValidatingWebhookConfigurationList, s conversion.Scope) error {
	return autoConvert_admissionregistration_ValidatingWebhookConfigurationList_To_v1beta1_ValidatingWebhookConfigurationList(in, out, s)
}

func autoConvert_v1beta1_Webhook_To_admissionregistration_Webhook(in *v1beta1.Webhook, out *admissionregistration.Webhook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]admissionregistration.RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*admissionregistration.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.NamespaceSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	return nil
}

// Convert_v1beta1_Webhook_To_admissionregistration_Webhook is an autogenerated conversion function.
func Convert_v1beta1_Webhook_To_admissionregistration_Webhook(in *v1beta1.Webhook, out *admissionregistration.Webhook, s conversion.Scope) error {
	return autoConvert_v1beta1_Webhook_To_admissionregistration_Webhook(in, out, s)
}

func autoConvert_admissionregistration_Webhook_To_v1beta1_Webhook(in *admissionregistration.Webhook, out *v1beta1.Webhook, s conversion.Scope) error {
	out.Name = in.Name
	if err := Convert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(&in.ClientConfig, &out.ClientConfig, s); err != nil {
		return err
	}
	out.Rules = *(*[]v1beta1.RuleWithOperations)(unsafe.Pointer(&in.Rules))
	out.FailurePolicy = (*v1beta1.FailurePolicyType)(unsafe.Pointer(in.FailurePolicy))
	out.NamespaceSelector = (*v1.LabelSelector)(unsafe.Pointer(in.NamespaceSelector))
	return nil
}

// Convert_admissionregistration_Webhook_To_v1beta1_Webhook is an autogenerated conversion function.
func Convert_admissionregistration_Webhook_To_v1beta1_Webhook(in *admissionregistration.Webhook, out *v1beta1.Webhook, s conversion.Scope) error {
	return autoConvert_admissionregistration_Webhook_To_v1beta1_Webhook(in, out, s)
}

func autoConvert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in *v1beta1.WebhookClientConfig, out *admissionregistration.WebhookClientConfig, s conversion.Scope) error {
	out.URL = (*string)(unsafe.Pointer(in.URL))
	out.Service = (*admissionregistration.ServiceReference)(unsafe.Pointer(in.Service))
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	return nil
}

// Convert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig is an autogenerated conversion function.
func Convert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in *v1beta1.WebhookClientConfig, out *admissionregistration.WebhookClientConfig, s conversion.Scope) error {
	return autoConvert_v1beta1_WebhookClientConfig_To_admissionregistration_WebhookClientConfig(in, out, s)
}

func autoConvert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(in *admissionregistration.WebhookClientConfig, out *v1beta1.WebhookClientConfig, s conversion.Scope) error {
	out.URL = (*string)(unsafe.Pointer(in.URL))
	out.Service = (*v1beta1.ServiceReference)(unsafe.Pointer(in.Service))
	out.CABundle = *(*[]byte)(unsafe.Pointer(&in.CABundle))
	return nil
}

// Convert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig is an autogenerated conversion function.
func Convert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(in *admissionregistration.WebhookClientConfig, out *v1beta1.WebhookClientConfig, s conversion.Scope) error {
	return autoConvert_admissionregistration_WebhookClientConfig_To_v1beta1_WebhookClientConfig(in, out, s)
}
