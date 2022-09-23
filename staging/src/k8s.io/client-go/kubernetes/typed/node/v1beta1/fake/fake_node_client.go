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

package fake

import (
	v1beta1 "k8s.io/client-go/kubernetes/typed/node/v1beta1"
	clientgorest "k8s.io/client-go/rest"
	clientgotesting "k8s.io/client-go/testing"
)

type FakeNodeV1beta1 struct {
	*clientgotesting.Fake
}

func (c *FakeNodeV1beta1) RuntimeClasses() v1beta1.RuntimeClassInterface {
	return &FakeRuntimeClasses{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeNodeV1beta1) RESTClient() clientgorest.Interface {
	var ret *clientgorest.RESTClient
	return ret
}
