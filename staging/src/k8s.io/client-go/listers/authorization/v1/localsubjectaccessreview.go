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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/lister-gen --output-base /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/vendor --output-package k8s.io/client-go/listers --input-dirs k8s.io/api/admission/v1beta1,k8s.io/api/admissionregistration/v1alpha1,k8s.io/api/admissionregistration/v1beta1,k8s.io/api/apps/v1,k8s.io/api/apps/v1beta1,k8s.io/api/apps/v1beta2,k8s.io/api/authentication/v1,k8s.io/api/authentication/v1beta1,k8s.io/api/authorization/v1,k8s.io/api/authorization/v1beta1,k8s.io/api/autoscaling/v1,k8s.io/api/autoscaling/v2beta1,k8s.io/api/batch/v1,k8s.io/api/batch/v1beta1,k8s.io/api/batch/v2alpha1,k8s.io/api/certificates/v1beta1,k8s.io/api/core/v1,k8s.io/api/events/v1beta1,k8s.io/api/extensions/v1beta1,k8s.io/api/imagepolicy/v1alpha1,k8s.io/api/networking/v1,k8s.io/api/policy/v1beta1,k8s.io/api/rbac/v1,k8s.io/api/rbac/v1alpha1,k8s.io/api/rbac/v1beta1,k8s.io/api/scheduling/v1alpha1,k8s.io/api/settings/v1alpha1,k8s.io/api/storage/v1,k8s.io/api/storage/v1alpha1,k8s.io/api/storage/v1beta1

// This file was automatically generated by lister-gen

package v1

import (
	v1 "k8s.io/api/authorization/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LocalSubjectAccessReviewLister helps list LocalSubjectAccessReviews.
type LocalSubjectAccessReviewLister interface {
	// List lists all LocalSubjectAccessReviews in the indexer.
	List(selector labels.Selector) (ret []*v1.LocalSubjectAccessReview, err error)
	// LocalSubjectAccessReviews returns an object that can list and get LocalSubjectAccessReviews.
	LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewNamespaceLister
	LocalSubjectAccessReviewListerExpansion
}

// localSubjectAccessReviewLister implements the LocalSubjectAccessReviewLister interface.
type localSubjectAccessReviewLister struct {
	indexer cache.Indexer
}

// NewLocalSubjectAccessReviewLister returns a new LocalSubjectAccessReviewLister.
func NewLocalSubjectAccessReviewLister(indexer cache.Indexer) LocalSubjectAccessReviewLister {
	return &localSubjectAccessReviewLister{indexer: indexer}
}

// List lists all LocalSubjectAccessReviews in the indexer.
func (s *localSubjectAccessReviewLister) List(selector labels.Selector) (ret []*v1.LocalSubjectAccessReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.LocalSubjectAccessReview))
	})
	return ret, err
}

// LocalSubjectAccessReviews returns an object that can list and get LocalSubjectAccessReviews.
func (s *localSubjectAccessReviewLister) LocalSubjectAccessReviews(namespace string) LocalSubjectAccessReviewNamespaceLister {
	return localSubjectAccessReviewNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// LocalSubjectAccessReviewNamespaceLister helps list and get LocalSubjectAccessReviews.
type LocalSubjectAccessReviewNamespaceLister interface {
	// List lists all LocalSubjectAccessReviews in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.LocalSubjectAccessReview, err error)
	// Get retrieves the LocalSubjectAccessReview from the indexer for a given namespace and name.
	Get(name string) (*v1.LocalSubjectAccessReview, error)
	LocalSubjectAccessReviewNamespaceListerExpansion
}

// localSubjectAccessReviewNamespaceLister implements the LocalSubjectAccessReviewNamespaceLister
// interface.
type localSubjectAccessReviewNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all LocalSubjectAccessReviews in the indexer for a given namespace.
func (s localSubjectAccessReviewNamespaceLister) List(selector labels.Selector) (ret []*v1.LocalSubjectAccessReview, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.LocalSubjectAccessReview))
	})
	return ret, err
}

// Get retrieves the LocalSubjectAccessReview from the indexer for a given namespace and name.
func (s localSubjectAccessReviewNamespaceLister) Get(name string) (*v1.LocalSubjectAccessReview, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("localsubjectaccessreview"), name)
	}
	return obj.(*v1.LocalSubjectAccessReview), nil
}
