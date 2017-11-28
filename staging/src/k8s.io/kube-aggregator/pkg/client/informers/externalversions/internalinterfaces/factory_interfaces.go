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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/local/go/src/k8s.io/kubernetes/informer-gen --output-base vendor/k8s.io/kube-aggregator/hack/../../.. --input-dirs k8s.io/kube-aggregator/pkg/apis/apiregistration --input-dirs k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1 --versioned-clientset-package k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset --internal-clientset-package k8s.io/kube-aggregator/pkg/client/clientset_generated/internalclientset --listers-package k8s.io/kube-aggregator/pkg/client/listers --output-package k8s.io/kube-aggregator/pkg/client/informers

// This file was automatically generated by informer-gen

package internalinterfaces

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	cache "k8s.io/client-go/tools/cache"
	clientset "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset"
	time "time"
)

type NewInformerFunc func(clientset.Interface, time.Duration) cache.SharedIndexInformer

// SharedInformerFactory a small interface to allow for adding an informer without an import cycle
type SharedInformerFactory interface {
	Start(stopCh <-chan struct{})
	InformerFor(obj runtime.Object, newFunc NewInformerFunc) cache.SharedIndexInformer
}

type TweakListOptionsFunc func(*v1.ListOptions)
