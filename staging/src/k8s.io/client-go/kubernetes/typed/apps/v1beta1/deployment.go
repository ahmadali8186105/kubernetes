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

	apiappsv1beta1 "k8s.io/api/apps/v1beta1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apimachinerypkgtypes "k8s.io/apimachinery/pkg/types"
	apimachinerypkgwatch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsappsv1beta1 "k8s.io/client-go/applyconfigurations/apps/v1beta1"
	clientgokubernetesscheme "k8s.io/client-go/kubernetes/scheme"
	clientgorest "k8s.io/client-go/rest"
)

// DeploymentsGetter has a method to return a DeploymentInterface.
// A group's client should implement this interface.
type DeploymentsGetter interface {
	Deployments(namespace string) DeploymentInterface
}

// DeploymentInterface has methods to work with Deployment resources.
type DeploymentInterface interface {
	Create(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.CreateOptions) (*apiappsv1beta1.Deployment, error)
	Update(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.UpdateOptions) (*apiappsv1beta1.Deployment, error)
	UpdateStatus(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.UpdateOptions) (*apiappsv1beta1.Deployment, error)
	Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error
	Get(ctx context.Context, name string, opts apismetav1.GetOptions) (*apiappsv1beta1.Deployment, error)
	List(ctx context.Context, opts apismetav1.ListOptions) (*apiappsv1beta1.DeploymentList, error)
	Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error)
	Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiappsv1beta1.Deployment, err error)
	Apply(ctx context.Context, deployment *applyconfigurationsappsv1beta1.DeploymentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.Deployment, err error)
	ApplyStatus(ctx context.Context, deployment *applyconfigurationsappsv1beta1.DeploymentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.Deployment, err error)
	DeploymentExpansion
}

// deployments implements DeploymentInterface
type deployments struct {
	client clientgorest.Interface
	ns     string
}

// newDeployments returns a Deployments
func newDeployments(c *AppsV1beta1Client, namespace string) *deployments {
	return &deployments{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deployment, and returns the corresponding deployment object, and an error if there is any.
func (c *deployments) Get(ctx context.Context, name string, options apismetav1.GetOptions) (result *apiappsv1beta1.Deployment, err error) {
	result = &apiappsv1beta1.Deployment{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		VersionedParams(&options, clientgokubernetesscheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deployments) List(ctx context.Context, opts apismetav1.ListOptions) (result *apiappsv1beta1.DeploymentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &apiappsv1beta1.DeploymentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a apimachinerypkgwatch.Interface that watches the requested deployments.
func (c *deployments) Watch(ctx context.Context, opts apismetav1.ListOptions) (apimachinerypkgwatch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a deployment and creates it.  Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Create(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.CreateOptions) (result *apiappsv1beta1.Deployment, err error) {
	result = &apiappsv1beta1.Deployment{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(deployment).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a deployment and updates it. Returns the server's representation of the deployment, and an error, if there is any.
func (c *deployments) Update(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.UpdateOptions) (result *apiappsv1beta1.Deployment, err error) {
	result = &apiappsv1beta1.Deployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deployment.Name).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(deployment).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *deployments) UpdateStatus(ctx context.Context, deployment *apiappsv1beta1.Deployment, opts apismetav1.UpdateOptions) (result *apiappsv1beta1.Deployment, err error) {
	result = &apiappsv1beta1.Deployment{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("deployments").
		Name(deployment.Name).
		SubResource("status").
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(deployment).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the deployment and deletes it. Returns an error if one occurs.
func (c *deployments) Delete(ctx context.Context, name string, opts apismetav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deployments) DeleteCollection(ctx context.Context, opts apismetav1.DeleteOptions, listOpts apismetav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("deployments").
		VersionedParams(&listOpts, clientgokubernetesscheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched deployment.
func (c *deployments) Patch(ctx context.Context, name string, pt apimachinerypkgtypes.PatchType, data []byte, opts apismetav1.PatchOptions, subresources ...string) (result *apiappsv1beta1.Deployment, err error) {
	result = &apiappsv1beta1.Deployment{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("deployments").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied deployment.
func (c *deployments) Apply(ctx context.Context, deployment *applyconfigurationsappsv1beta1.DeploymentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.Deployment, err error) {
	if deployment == nil {
		return nil, fmt.Errorf("deployment provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}
	name := deployment.Name
	if name == nil {
		return nil, fmt.Errorf("deployment.Name must be provided to Apply")
	}
	result = &apiappsv1beta1.Deployment{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Namespace(c.ns).
		Resource("deployments").
		Name(*name).
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *deployments) ApplyStatus(ctx context.Context, deployment *applyconfigurationsappsv1beta1.DeploymentApplyConfiguration, opts apismetav1.ApplyOptions) (result *apiappsv1beta1.Deployment, err error) {
	if deployment == nil {
		return nil, fmt.Errorf("deployment provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(deployment)
	if err != nil {
		return nil, err
	}

	name := deployment.Name
	if name == nil {
		return nil, fmt.Errorf("deployment.Name must be provided to Apply")
	}

	result = &apiappsv1beta1.Deployment{}
	err = c.client.Patch(apimachinerypkgtypes.ApplyPatchType).
		Namespace(c.ns).
		Resource("deployments").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, clientgokubernetesscheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
