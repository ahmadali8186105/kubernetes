/*
Copyright 2016 The Kubernetes Authors.

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

package install

import (
	clientset "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/controller/informers"
	"k8s.io/kubernetes/pkg/quota"
	"k8s.io/kubernetes/pkg/quota/evaluator/core"
	"k8s.io/kubernetes/pkg/quota/evaluator/extensions"
)

// NewRegistry returns a registry of quota evaluators.
// If a shared informer factory is provided, it is used by evaluators rather than performing direct queries.
func NewRegistry(kubeClient clientset.Interface, f informers.SharedInformerFactory) quota.Registry {
	return &quota.UnionRegistry{
		core.NewRegistry(kubeClient, f),
		extensions.NewRegistry(kubeClient, f),
	}
}
