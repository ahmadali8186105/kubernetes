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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/client-gen --output-base /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/vendor --clientset-path=k8s.io/client-go --clientset-name=kubernetes --input-base=k8s.io/kubernetes/vendor/k8s.io/api --input=core/v1,admissionregistration/v1alpha1,admissionregistration/v1beta1,apps/v1beta1,apps/v1beta2,apps/v1,authentication/v1,authentication/v1beta1,authorization/v1,authorization/v1beta1,autoscaling/v1,autoscaling/v2beta1,batch/v1,batch/v1beta1,batch/v2alpha1,certificates/v1beta1,extensions/v1beta1,events/v1beta1,networking/v1,policy/v1beta1,rbac/v1,rbac/v1beta1,rbac/v1alpha1,scheduling/v1alpha1,settings/v1alpha1,storage/v1beta1,storage/v1,storage/v1alpha1

package v1

import (
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	scheme "k8s.io/client-go/kubernetes/scheme"
	rest "k8s.io/client-go/rest"
)

// NamespacesGetter has a method to return a NamespaceInterface.
// A group's client should implement this interface.
type NamespacesGetter interface {
	Namespaces() NamespaceInterface
}

// NamespaceInterface has methods to work with Namespace resources.
type NamespaceInterface interface {
	Create(*v1.Namespace) (*v1.Namespace, error)
	Update(*v1.Namespace) (*v1.Namespace, error)
	UpdateStatus(*v1.Namespace) (*v1.Namespace, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Namespace, error)
	List(opts meta_v1.ListOptions) (*v1.NamespaceList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Namespace, err error)
	NamespaceExpansion
}

// namespaces implements NamespaceInterface
type namespaces struct {
	client rest.Interface
}

// newNamespaces returns a Namespaces
func newNamespaces(c *CoreV1Client) *namespaces {
	return &namespaces{
		client: c.RESTClient(),
	}
}

// Get takes name of the namespace, and returns the corresponding namespace object, and an error if there is any.
func (c *namespaces) Get(name string, options meta_v1.GetOptions) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Get().
		Resource("namespaces").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Namespaces that match those selectors.
func (c *namespaces) List(opts meta_v1.ListOptions) (result *v1.NamespaceList, err error) {
	result = &v1.NamespaceList{}
	err = c.client.Get().
		Resource("namespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested namespaces.
func (c *namespaces) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("namespaces").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a namespace and creates it.  Returns the server's representation of the namespace, and an error, if there is any.
func (c *namespaces) Create(namespace *v1.Namespace) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Post().
		Resource("namespaces").
		Body(namespace).
		Do().
		Into(result)
	return
}

// Update takes the representation of a namespace and updates it. Returns the server's representation of the namespace, and an error, if there is any.
func (c *namespaces) Update(namespace *v1.Namespace) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Put().
		Resource("namespaces").
		Name(namespace.Name).
		Body(namespace).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *namespaces) UpdateStatus(namespace *v1.Namespace) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Put().
		Resource("namespaces").
		Name(namespace.Name).
		SubResource("status").
		Body(namespace).
		Do().
		Into(result)
	return
}

// Delete takes name of the namespace and deletes it. Returns an error if one occurs.
func (c *namespaces) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("namespaces").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *namespaces) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("namespaces").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched namespace.
func (c *namespaces) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Namespace, err error) {
	result = &v1.Namespace{}
	err = c.client.Patch(pt).
		Resource("namespaces").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
