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
	v1beta1 "k8s.io/api/certificates/v1beta1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	certificates "k8s.io/kubernetes/pkg/apis/certificates"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest,
		Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest,
		Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition,
		Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition,
		Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList,
		Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList,
		Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec,
		Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec,
		Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus,
		Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus,
	)
}

func autoConvert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *v1beta1.CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in *v1beta1.CertificateSigningRequest, out *certificates.CertificateSigningRequest, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequest_To_certificates_CertificateSigningRequest(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *v1beta1.CertificateSigningRequest, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in *certificates.CertificateSigningRequest, out *v1beta1.CertificateSigningRequest, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequest_To_v1beta1_CertificateSigningRequest(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *v1beta1.CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
	out.Type = certificates.RequestConditionType(in.Type)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

// Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in *v1beta1.CertificateSigningRequestCondition, out *certificates.CertificateSigningRequestCondition, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestCondition_To_certificates_CertificateSigningRequestCondition(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *v1beta1.CertificateSigningRequestCondition, s conversion.Scope) error {
	out.Type = v1beta1.RequestConditionType(in.Type)
	out.Reason = in.Reason
	out.Message = in.Message
	out.LastUpdateTime = in.LastUpdateTime
	return nil
}

// Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in *certificates.CertificateSigningRequestCondition, out *v1beta1.CertificateSigningRequestCondition, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestCondition_To_v1beta1_CertificateSigningRequestCondition(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *v1beta1.CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]certificates.CertificateSigningRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in *v1beta1.CertificateSigningRequestList, out *certificates.CertificateSigningRequestList, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestList_To_certificates_CertificateSigningRequestList(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *v1beta1.CertificateSigningRequestList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]v1beta1.CertificateSigningRequest)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in *certificates.CertificateSigningRequestList, out *v1beta1.CertificateSigningRequestList, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestList_To_v1beta1_CertificateSigningRequestList(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *v1beta1.CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
	out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
	out.Usages = *(*[]certificates.KeyUsage)(unsafe.Pointer(&in.Usages))
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]certificates.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in *v1beta1.CertificateSigningRequestSpec, out *certificates.CertificateSigningRequestSpec, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestSpec_To_certificates_CertificateSigningRequestSpec(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *v1beta1.CertificateSigningRequestSpec, s conversion.Scope) error {
	out.Request = *(*[]byte)(unsafe.Pointer(&in.Request))
	out.Usages = *(*[]v1beta1.KeyUsage)(unsafe.Pointer(&in.Usages))
	out.Username = in.Username
	out.UID = in.UID
	out.Groups = *(*[]string)(unsafe.Pointer(&in.Groups))
	out.Extra = *(*map[string]v1beta1.ExtraValue)(unsafe.Pointer(&in.Extra))
	return nil
}

// Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in *certificates.CertificateSigningRequestSpec, out *v1beta1.CertificateSigningRequestSpec, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestSpec_To_v1beta1_CertificateSigningRequestSpec(in, out, s)
}

func autoConvert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *v1beta1.CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
	out.Conditions = *(*[]certificates.CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	return nil
}

// Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus is an autogenerated conversion function.
func Convert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in *v1beta1.CertificateSigningRequestStatus, out *certificates.CertificateSigningRequestStatus, s conversion.Scope) error {
	return autoConvert_v1beta1_CertificateSigningRequestStatus_To_certificates_CertificateSigningRequestStatus(in, out, s)
}

func autoConvert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *v1beta1.CertificateSigningRequestStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1beta1.CertificateSigningRequestCondition)(unsafe.Pointer(&in.Conditions))
	out.Certificate = *(*[]byte)(unsafe.Pointer(&in.Certificate))
	return nil
}

// Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus is an autogenerated conversion function.
func Convert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in *certificates.CertificateSigningRequestStatus, out *v1beta1.CertificateSigningRequestStatus, s conversion.Scope) error {
	return autoConvert_certificates_CertificateSigningRequestStatus_To_v1beta1_CertificateSigningRequestStatus(in, out, s)
}
