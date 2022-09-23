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

package v1beta1

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	apistoragev1beta1 "k8s.io/api/storage/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

// CSIStorageCapacitiesGetter has a method to return a CSIStorageCapacityInterface.
// A group's client should implement this interface.
type CSIStorageCapacitiesGetter interface {
	CSIStorageCapacities(namespace string) CSIStorageCapacityInterface
}

// CSIStorageCapacityInterface has methods to work with CSIStorageCapacity resources.
type CSIStorageCapacityInterface interface {
	Create(ctx context.Context, cSIStorageCapacity *apistoragev1beta1.CSIStorageCapacity, opts apismetav1.CreateOptions) (*apistoragev1beta1.CSIStorageCapacity, error)
	Update(ctx context.Context, cSIStorageCapacity *apistoragev1beta1.CSIStorageCapacity, opts apismetav1.UpdateOptions) (*apistoragev1beta1.CSIStorageCapacity, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apistoragev1beta1.CSIStorageCapacity, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apistoragev1beta1.CSIStorageCapacityList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apistoragev1beta1.CSIStorageCapacity, err error)
	Apply(ctx context.Context, cSIStorageCapacity *applyconfigurationsstoragev1beta1.CSIStorageCapacityApplyConfiguration, opts apismetav1.ApplyOptions) (result *apistoragev1beta1.CSIStorageCapacity, err error)
	CSIStorageCapacityExpansion
}

// cSIStorageCapacities implements CSIStorageCapacityInterface
type cSIStorageCapacities struct {
	client clientgorest.Interface
	ns     string
}

// newCSIStorageCapacities returns a CSIStorageCapacities
func newCSIStorageCapacities(c *StorageV1beta1Client, namespace string) *cSIStorageCapacities {
	return &cSIStorageCapacities{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cSIStorageCapacity, and returns the corresponding cSIStorageCapacity object, and an error if there is any.
func (c *cSIStorageCapacities) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apistoragev1beta1.CSIStorageCapacity, err error) {
	result = &apistoragev1beta1.CSIStorageCapacity{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("csistoragecapacities").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CSIStorageCapacities that match those selectors.
func (c *cSIStorageCapacities) List(ctx context.Context, opts apismetav1.ListOptions) (result *apistoragev1beta1.CSIStorageCapacityList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apistoragev1beta1.CSIStorageCapacityList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("csistoragecapacities").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested cSIStorageCapacities.
func (c *cSIStorageCapacities) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("csistoragecapacities").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cSIStorageCapacity and creates it.  Returns the server's representation of the cSIStorageCapacity, and an error, if there is any.
func (c *cSIStorageCapacities) Create(ctx context.Context, cSIStorageCapacity *apistoragev1beta1.CSIStorageCapacity, opts apismetav1.CreateOptions) (result *apistoragev1beta1.CSIStorageCapacity, err error) {
	result = &apistoragev1beta1.CSIStorageCapacity{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("csistoragecapacities").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(cSIStorageCapacity).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cSIStorageCapacity and updates it. Returns the server's representation of the cSIStorageCapacity, and an error, if there is any.
func (c *cSIStorageCapacities) Update(ctx context.Context, cSIStorageCapacity *apistoragev1beta1.CSIStorageCapacity, opts apismetav1.UpdateOptions) (result *apistoragev1beta1.CSIStorageCapacity, err error) {
	result = &apistoragev1beta1.CSIStorageCapacity{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("csistoragecapacities").
		Name(cSIStorageCapacity.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(cSIStorageCapacity).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cSIStorageCapacity and deletes it. Returns an error if one occurs.
func (c *cSIStorageCapacities) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("csistoragecapacities").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cSIStorageCapacities) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("csistoragecapacities").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cSIStorageCapacity.
func (c *cSIStorageCapacities) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apistoragev1beta1.CSIStorageCapacity, err error) {
	result = &apistoragev1beta1.CSIStorageCapacity{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("csistoragecapacities").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied cSIStorageCapacity.
func (c *cSIStorageCapacities) Apply(ctx context.Context, cSIStorageCapacity *applyconfigurationsstoragev1beta1.CSIStorageCapacityApplyConfiguration, opts apismetav1.ApplyOptions) (result *apistoragev1beta1.CSIStorageCapacity, err error) {
	if cSIStorageCapacity == nil {
		return nil, fmt.Errorf("cSIStorageCapacity provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(cSIStorageCapacity)
	if err != nil {
		return nil, err
	}
	name := cSIStorageCapacity.Name
	if name == nil {
		return nil, fmt.Errorf("cSIStorageCapacity.Name must be provided to Apply")
	}
	result = &apistoragev1beta1.CSIStorageCapacity{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Namespace(c.ns).
		Resource("csistoragecapacities").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
