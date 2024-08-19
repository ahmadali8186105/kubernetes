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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"

	v1alpha1 "k8s.io/api/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rbacv1alpha1 "k8s.io/client-go/applyconfigurations/rbac/v1alpha1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// RolesGetter has a method to return a RoleInterface.
// A group's client should implement this interface.
type RolesGetter interface {
	Roles(namespace string) RoleInterface
}

// RoleInterface has methods to work with Role resources.
type RoleInterface interface {
	Create(ctx context.Context, role *v1alpha1.Role, opts v1.CreateOptions) (*v1alpha1.Role, error)
	Update(ctx context.Context, role *v1alpha1.Role, opts v1.UpdateOptions) (*v1alpha1.Role, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Role, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RoleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Role, err error)
	Apply(ctx context.Context, role *rbacv1alpha1.RoleApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.Role, err error)
	RoleExpansion
}

// roles implements RoleInterface
type roles struct {
	*gentype.ClientWithListAndApply[*v1alpha1.Role, *v1alpha1.RoleList, *rbacv1alpha1.RoleApplyConfiguration]
}

// newRoles returns a Roles
func newRoles(c *RbacV1alpha1Client, namespace string) *roles {
	return &roles{
		gentype.NewClientWithListAndApply[*v1alpha1.Role, *v1alpha1.RoleList, *rbacv1alpha1.RoleApplyConfiguration](
			"roles",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.Role { return &v1alpha1.Role{} },
			func() *v1alpha1.RoleList { return &v1alpha1.RoleList{} },
			true,
		),
	}
}
