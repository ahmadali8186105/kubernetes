/*
Copyright 2019 The Kubernetes Authors.

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

package v1

import (
	admissionregistrationv1 "k8s.io/api/admissionregistration/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	utilpointer "k8s.io/utils/pointer"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaultsValidatingWebhook sets defaults for ValidatingWebhook
func SetDefaultsValidatingWebhook(obj *admissionregistrationv1.ValidatingWebhook) {
	if obj.FailurePolicy == nil {
		policy := admissionregistrationv1.Fail
		obj.FailurePolicy = &policy
	}
	if obj.MatchPolicy == nil {
		policy := admissionregistrationv1.Equivalent
		obj.MatchPolicy = &policy
	}
	if obj.NamespaceSelector == nil {
		selector := metav1.LabelSelector{}
		obj.NamespaceSelector = &selector
	}
	if obj.ObjectSelector == nil {
		selector := metav1.LabelSelector{}
		obj.ObjectSelector = &selector
	}
	if obj.TimeoutSeconds == nil {
		obj.TimeoutSeconds = new(int32)
		*obj.TimeoutSeconds = 10
	}
}

// SetDefaultsMutatingWebhook sets defaults for ValidatingMutatingWebhook
func SetDefaultsMutatingWebhook(obj *admissionregistrationv1.MutatingWebhook) {
	if obj.FailurePolicy == nil {
		policy := admissionregistrationv1.Fail
		obj.FailurePolicy = &policy
	}
	if obj.MatchPolicy == nil {
		policy := admissionregistrationv1.Equivalent
		obj.MatchPolicy = &policy
	}
	if obj.NamespaceSelector == nil {
		selector := metav1.LabelSelector{}
		obj.NamespaceSelector = &selector
	}
	if obj.ObjectSelector == nil {
		selector := metav1.LabelSelector{}
		obj.ObjectSelector = &selector
	}
	if obj.TimeoutSeconds == nil {
		obj.TimeoutSeconds = new(int32)
		*obj.TimeoutSeconds = 10
	}
	if obj.ReinvocationPolicy == nil {
		never := admissionregistrationv1.NeverReinvocationPolicy
		obj.ReinvocationPolicy = &never
	}
}

// SetDefaultsRule sets defaults for Webhook's Rule
func SetDefaultsRule(obj *admissionregistrationv1.Rule) {
	if obj.Scope == nil {
		s := admissionregistrationv1.AllScopes
		obj.Scope = &s
	}
}

// SetDefaultsServiceReference sets defaults for Webhook's ServiceReference
func SetDefaultsServiceReference(obj *admissionregistrationv1.ServiceReference) {
	if obj.Port == nil {
		obj.Port = utilpointer.Int32Ptr(443)
	}
}
