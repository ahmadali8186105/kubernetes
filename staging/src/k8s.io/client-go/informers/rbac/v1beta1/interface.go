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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/informer-gen --output-base /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/vendor --output-package k8s.io/client-go/informers --single-directory --input-dirs k8s.io/api/admission/v1beta1,k8s.io/api/admissionregistration/v1alpha1,k8s.io/api/admissionregistration/v1beta1,k8s.io/api/apps/v1,k8s.io/api/apps/v1beta1,k8s.io/api/apps/v1beta2,k8s.io/api/authentication/v1,k8s.io/api/authentication/v1beta1,k8s.io/api/authorization/v1,k8s.io/api/authorization/v1beta1,k8s.io/api/autoscaling/v1,k8s.io/api/autoscaling/v2beta1,k8s.io/api/batch/v1,k8s.io/api/batch/v1beta1,k8s.io/api/batch/v2alpha1,k8s.io/api/certificates/v1beta1,k8s.io/api/core/v1,k8s.io/api/events/v1beta1,k8s.io/api/extensions/v1beta1,k8s.io/api/imagepolicy/v1alpha1,k8s.io/api/networking/v1,k8s.io/api/policy/v1beta1,k8s.io/api/rbac/v1,k8s.io/api/rbac/v1alpha1,k8s.io/api/rbac/v1beta1,k8s.io/api/scheduling/v1alpha1,k8s.io/api/settings/v1alpha1,k8s.io/api/storage/v1,k8s.io/api/storage/v1alpha1,k8s.io/api/storage/v1beta1 --versioned-clientset-package k8s.io/client-go/kubernetes --listers-package k8s.io/client-go/listers

// This file was automatically generated by informer-gen

package v1beta1

import (
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterRoles returns a ClusterRoleInformer.
	ClusterRoles() ClusterRoleInformer
	// ClusterRoleBindings returns a ClusterRoleBindingInformer.
	ClusterRoleBindings() ClusterRoleBindingInformer
	// Roles returns a RoleInformer.
	Roles() RoleInformer
	// RoleBindings returns a RoleBindingInformer.
	RoleBindings() RoleBindingInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterRoles returns a ClusterRoleInformer.
func (v *version) ClusterRoles() ClusterRoleInformer {
	return &clusterRoleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterRoleBindings returns a ClusterRoleBindingInformer.
func (v *version) ClusterRoleBindings() ClusterRoleBindingInformer {
	return &clusterRoleBindingInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Roles returns a RoleInformer.
func (v *version) Roles() RoleInformer {
	return &roleInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RoleBindings returns a RoleBindingInformer.
func (v *version) RoleBindings() RoleBindingInformer {
	return &roleBindingInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
