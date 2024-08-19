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

	v1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	storagev1 "k8s.io/client-go/applyconfigurations/storage/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CSINodesGetter has a method to return a CSINodeInterface.
// A group's client should implement this interface.
type CSINodesGetter interface {
	CSINodes() CSINodeInterface
}

// CSINodeInterface has methods to work with CSINode resources.
type CSINodeInterface interface {
	Create(ctx context.Context, cSINode *v1.CSINode, opts metav1.CreateOptions) (*v1.CSINode, error)
	Update(ctx context.Context, cSINode *v1.CSINode, opts metav1.UpdateOptions) (*v1.CSINode, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CSINode, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CSINodeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CSINode, err error)
	Apply(ctx context.Context, cSINode *storagev1.CSINodeApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CSINode, err error)
	CSINodeExpansion
}

// cSINodes implements CSINodeInterface
type cSINodes struct {
	*gentype.ClientWithListAndApply[*v1.CSINode, *v1.CSINodeList, *storagev1.CSINodeApplyConfiguration]
}

// newCSINodes returns a CSINodes
func newCSINodes(c *StorageV1Client) *cSINodes {
	return &cSINodes{
		gentype.NewClientWithListAndApply[*v1.CSINode, *v1.CSINodeList, *storagev1.CSINodeApplyConfiguration](
			"csinodes",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.CSINode { return &v1.CSINode{} },
			func() *v1.CSINodeList { return &v1.CSINodeList{} },
			true,
		),
	}
}
