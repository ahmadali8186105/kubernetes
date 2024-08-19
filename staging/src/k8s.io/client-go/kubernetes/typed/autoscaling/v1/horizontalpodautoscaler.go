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

package v1

import (
	"context"

	v1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	autoscalingv1 "k8s.io/client-go/applyconfigurations/autoscaling/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// HorizontalPodAutoscalersGetter has a method to return a HorizontalPodAutoscalerInterface.
// A group's client should implement this interface.
type HorizontalPodAutoscalersGetter interface {
	HorizontalPodAutoscalers(namespace string) HorizontalPodAutoscalerInterface
}

// HorizontalPodAutoscalerInterface has methods to work with HorizontalPodAutoscaler resources.
type HorizontalPodAutoscalerInterface interface {
	Create(ctx context.Context, horizontalPodAutoscaler *v1.HorizontalPodAutoscaler, opts metav1.CreateOptions) (*v1.HorizontalPodAutoscaler, error)
	Update(ctx context.Context, horizontalPodAutoscaler *v1.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*v1.HorizontalPodAutoscaler, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, horizontalPodAutoscaler *v1.HorizontalPodAutoscaler, opts metav1.UpdateOptions) (*v1.HorizontalPodAutoscaler, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.HorizontalPodAutoscaler, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.HorizontalPodAutoscalerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.HorizontalPodAutoscaler, err error)
	Apply(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HorizontalPodAutoscaler, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, horizontalPodAutoscaler *autoscalingv1.HorizontalPodAutoscalerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.HorizontalPodAutoscaler, err error)
	HorizontalPodAutoscalerExpansion
}

// horizontalPodAutoscalers implements HorizontalPodAutoscalerInterface
type horizontalPodAutoscalers struct {
	*gentype.ClientWithListAndApply[*v1.HorizontalPodAutoscaler, *v1.HorizontalPodAutoscalerList, *autoscalingv1.HorizontalPodAutoscalerApplyConfiguration]
}

// newHorizontalPodAutoscalers returns a HorizontalPodAutoscalers
func newHorizontalPodAutoscalers(c *AutoscalingV1Client, namespace string) *horizontalPodAutoscalers {
	return &horizontalPodAutoscalers{
		gentype.NewClientWithListAndApply[*v1.HorizontalPodAutoscaler, *v1.HorizontalPodAutoscalerList, *autoscalingv1.HorizontalPodAutoscalerApplyConfiguration](
			"horizontalpodautoscalers",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.HorizontalPodAutoscaler { return &v1.HorizontalPodAutoscaler{} },
			func() *v1.HorizontalPodAutoscalerList { return &v1.HorizontalPodAutoscalerList{} },
			true,
		),
	}
}
