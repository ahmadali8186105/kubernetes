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

	apiflowcontrolv1beta1 "k8s.io/api/flowcontrol/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsflowcontrolv1beta1 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

// FlowSchemasGetter has a method to return a FlowSchemaInterface.
// A group's client should implement this interface.
type FlowSchemasGetter interface {
	FlowSchemas() FlowSchemaInterface
}

// FlowSchemaInterface has methods to work with FlowSchema resources.
type FlowSchemaInterface interface {
	Create(ctx context.Context, flowSchema *apiflowcontrolv1beta1.FlowSchema, opts apismetav1.CreateOptions) (*apiflowcontrolv1beta1.FlowSchema, error)
	Update(ctx context.Context, flowSchema *apiflowcontrolv1beta1.FlowSchema, opts apismetav1.UpdateOptions) (*apiflowcontrolv1beta1.FlowSchema, error)
	UpdateStatus(ctx context.Context, flowSchema *apiflowcontrolv1beta1.FlowSchema, opts apismetav1.UpdateOptions) (*apiflowcontrolv1beta1.FlowSchema, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apiflowcontrolv1beta1.FlowSchema, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apiflowcontrolv1beta1.FlowSchemaList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiflowcontrolv1beta1.FlowSchema, err error)
	Apply(ctx context.Context, flowSchema *applyconfigurationsflowcontrolv1beta1.FlowSchemaApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error)
	ApplyStatus(ctx context.Context, flowSchema *applyconfigurationsflowcontrolv1beta1.FlowSchemaApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error)
	FlowSchemaExpansion
}

// flowSchemas implements FlowSchemaInterface
type flowSchemas struct {
	client clientgorest.Interface
}

// newFlowSchemas returns a FlowSchemas
func newFlowSchemas(c *FlowcontrolV1beta1Client) *flowSchemas {
	return &flowSchemas{
		client: c.RESTClient(),
	}
}

// Get takes name of the flowSchema, and returns the corresponding flowSchema object, and an error if there is any.
func (c *flowSchemas) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error) {
	result = &apiflowcontrolv1beta1.FlowSchema{}
	err = c.client.Get().
		Resource("flowschemas").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FlowSchemas that match those selectors.
func (c *flowSchemas) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiflowcontrolv1beta1.FlowSchemaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apiflowcontrolv1beta1.FlowSchemaList{}
	err = c.client.Get().
		Resource("flowschemas").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested flowSchemas.
func (c *flowSchemas) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("flowschemas").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a flowSchema and creates it.  Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *flowSchemas) Create(ctx context.Context, flowSchema *apiflowcontrolv1beta1.FlowSchema, opts apismetav1.CreateOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error) {
	result = &apiflowcontrolv1beta1.FlowSchema{}
	err = c.client.Post().
		Resource("flowschemas").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(flowSchema).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a flowSchema and updates it. Returns the server's representation of the flowSchema, and an error, if there is any.
func (c *flowSchemas) Update(ctx context.Context, flowSchema *apiflowcontrolv1beta1.FlowSchema, opts apismetav1.UpdateOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error) {
	result = &apiflowcontrolv1beta1.FlowSchema{}
	err = c.client.Put().
		Resource("flowschemas").
		Name(flowSchema.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(flowSchema).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *flowSchemas) UpdateStatus(ctx context.Context, flowSchema *apiflowcontrolv1beta1.FlowSchema, opts apismetav1.UpdateOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error) {
	result = &apiflowcontrolv1beta1.FlowSchema{}
	err = c.client.Put().
		Resource("flowschemas").
		Name(flowSchema.Name).
		SubResource("status").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(flowSchema).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the flowSchema and deletes it. Returns an error if one occurs.
func (c *flowSchemas) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("flowschemas").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flowSchemas) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("flowschemas").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched flowSchema.
func (c *flowSchemas) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiflowcontrolv1beta1.FlowSchema, err error) {
	result = &apiflowcontrolv1beta1.FlowSchema{}
	err = c.client.Patch(pt).
		Resource("flowschemas").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied flowSchema.
func (c *flowSchemas) Apply(ctx context.Context, flowSchema *applyconfigurationsflowcontrolv1beta1.FlowSchemaApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error) {
	if flowSchema == nil {
		return nil, fmt.Errorf("flowSchema provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(flowSchema)
	if err != nil {
		return nil, err
	}
	name := flowSchema.Name
	if name == nil {
		return nil, fmt.Errorf("flowSchema.Name must be provided to Apply")
	}
	result = &apiflowcontrolv1beta1.FlowSchema{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Resource("flowschemas").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *flowSchemas) ApplyStatus(ctx context.Context, flowSchema *applyconfigurationsflowcontrolv1beta1.FlowSchemaApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiflowcontrolv1beta1.FlowSchema, err error) {
	if flowSchema == nil {
		return nil, fmt.Errorf("flowSchema provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(flowSchema)
	if err != nil {
		return nil, err
	}

	name := flowSchema.Name
	if name == nil {
		return nil, fmt.Errorf("flowSchema.Name must be provided to Apply")
	}

	result = &apiflowcontrolv1beta1.FlowSchema{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Resource("flowschemas").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
