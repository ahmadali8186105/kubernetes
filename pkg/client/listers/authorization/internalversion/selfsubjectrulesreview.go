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

// This file was automatically generated by: /Users/sts/Quellen/kubernetes/src/k8s.io/kubernetes/_output/bin/lister-gen --input-dirs k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/admission,k8s.io/kubernetes/pkg/apis/admissionregistration,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/networking,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/scheduling,k8s.io/kubernetes/pkg/apis/settings,k8s.io/kubernetes/pkg/apis/storage

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	authorization "k8s.io/kubernetes/pkg/apis/authorization"
)

// SelfSubjectRulesReviewLister helps list SelfSubjectRulesReviews.
type SelfSubjectRulesReviewLister interface {
	// List lists all SelfSubjectRulesReviews in the indexer.
	List(selector labels.Selector) (ret []*authorization.SelfSubjectRulesReview, err error)
	// Get retrieves the SelfSubjectRulesReview from the index for a given name.
	Get(name string) (*authorization.SelfSubjectRulesReview, error)
	SelfSubjectRulesReviewListerExpansion
}

// selfSubjectRulesReviewLister implements the SelfSubjectRulesReviewLister interface.
type selfSubjectRulesReviewLister struct {
	indexer cache.Indexer
}

// NewSelfSubjectRulesReviewLister returns a new SelfSubjectRulesReviewLister.
func NewSelfSubjectRulesReviewLister(indexer cache.Indexer) SelfSubjectRulesReviewLister {
	return &selfSubjectRulesReviewLister{indexer: indexer}
}

// List lists all SelfSubjectRulesReviews in the indexer.
func (s *selfSubjectRulesReviewLister) List(selector labels.Selector) (ret []*authorization.SelfSubjectRulesReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*authorization.SelfSubjectRulesReview))
	})
	return ret, err
}

// Get retrieves the SelfSubjectRulesReview from the index for a given name.
func (s *selfSubjectRulesReviewLister) Get(name string) (*authorization.SelfSubjectRulesReview, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(authorization.Resource("selfsubjectrulesreview"), name)
	}
	return obj.(*authorization.SelfSubjectRulesReview), nil
}
