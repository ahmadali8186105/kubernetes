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

package v1beta3

import (
	"context"

	v1beta3 "k8s.io/api/flowcontrol/v1beta3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	flowcontrolv1beta3 "k8s.io/client-go/applyconfigurations/flowcontrol/v1beta3"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// FlowSchemasGetter has a method to return a FlowSchemaInterface.
// A group's client should implement this interface.
type FlowSchemasGetter interface {
	FlowSchemas() FlowSchemaInterface
}

// FlowSchemaInterface has methods to work with FlowSchema resources.
type FlowSchemaInterface interface {
	Create(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.CreateOptions) (*v1beta3.FlowSchema, error)
	Update(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (*v1beta3.FlowSchema, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, flowSchema *v1beta3.FlowSchema, opts v1.UpdateOptions) (*v1beta3.FlowSchema, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta3.FlowSchema, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta3.FlowSchemaList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta3.FlowSchema, err error)
	Apply(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, flowSchema *flowcontrolv1beta3.FlowSchemaApplyConfiguration, opts v1.ApplyOptions) (result *v1beta3.FlowSchema, err error)
	FlowSchemaExpansion
}

// flowSchemas implements FlowSchemaInterface
type flowSchemas struct {
	*gentype.ClientWithListAndApply[*v1beta3.FlowSchema, *v1beta3.FlowSchemaList, *flowcontrolv1beta3.FlowSchemaApplyConfiguration]
}

// newFlowSchemas returns a FlowSchemas
func newFlowSchemas(c *FlowcontrolV1beta3Client) *flowSchemas {
	return &flowSchemas{
		gentype.NewClientWithListAndApply[*v1beta3.FlowSchema, *v1beta3.FlowSchemaList, *flowcontrolv1beta3.FlowSchemaApplyConfiguration](
			"flowschemas",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1beta3.FlowSchema { return &v1beta3.FlowSchema{} },
			func() *v1beta3.FlowSchemaList { return &v1beta3.FlowSchemaList{} },
			true,
		),
	}
}
