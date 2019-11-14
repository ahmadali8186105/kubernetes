/*
Copyright 2019 The Kubernetes Authors.

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

package flowcontrol

import (
	"strings"

	rmtypesv1alpha1 "k8s.io/api/flowcontrol/v1alpha1"
	"k8s.io/apiserver/pkg/authentication/serviceaccount"
)

func matchesFlowSchema(digest RequestDigest, flowSchema *rmtypesv1alpha1.FlowSchema) bool {
	for _, policyRule := range flowSchema.Spec.Rules {
		subjectMatches := false
		for _, subject := range policyRule.Subjects {
			if matchesSubject(digest, subject) {
				subjectMatches = true
				break
			}
		}
		if !subjectMatches {
			continue
		}

		for _, rr := range policyRule.ResourceRules {
			if matchesResourcePolicyRule(digest, rr) {
				return true
			}
		}
		for _, rr := range policyRule.NonResourceRules {
			if matchesNonResourcePolicyRule(digest, rr) {
				return true
			}
		}
	}
	return false
}

func matchesSubject(digest RequestDigest, subject rmtypesv1alpha1.Subject) bool {
	user := digest.User
	switch subject.Kind {
	case rmtypesv1alpha1.SubjectKindUser:
		return subject.User.Name == rmtypesv1alpha1.NameAll || subject.User.Name == user.GetName()
	case rmtypesv1alpha1.SubjectKindGroup:
		return containsString(subject.Group.Name, user.GetGroups(), rmtypesv1alpha1.NameAll)
	case rmtypesv1alpha1.SubjectKindServiceAccount:
		if subject.ServiceAccount.Name == rmtypesv1alpha1.NameAll {
			return serviceAccountMatchesNamespace(subject.ServiceAccount.Namespace, user.GetName())
		}
		return serviceaccount.MatchesUsername(subject.ServiceAccount.Namespace, subject.ServiceAccount.Name, user.GetName())
	default:
		return false
	}
}

// serviceAccountMatchesNamespace checks whether the provided service account username matches the namespace, without
// allocating. Use this when checking a service account namespace against a known string.
// This is copied from `k8s.io/apiserver/pkg/authentication/serviceaccount::MatchesUsername` and simplified to not check the name part.
func serviceAccountMatchesNamespace(namespace string, username string) bool {
	const (
		ServiceAccountUsernamePrefix    = "system:serviceaccount:"
		ServiceAccountUsernameSeparator = ":"
	)
	if !strings.HasPrefix(username, ServiceAccountUsernamePrefix) {
		return false
	}
	username = username[len(ServiceAccountUsernamePrefix):]

	if !strings.HasPrefix(username, namespace) {
		return false
	}
	username = username[len(namespace):]

	return strings.HasPrefix(username, ServiceAccountUsernameSeparator)
}

func matchesResourcePolicyRule(digest RequestDigest, policyRule rmtypesv1alpha1.ResourcePolicyRule) bool {
	if !matchPolicyRuleVerb(policyRule.Verbs, digest.RequestInfo.Verb) {
		return false
	}
	if !digest.RequestInfo.IsResourceRequest {
		return false
	}
	if !matchPolicyRuleResource(policyRule.Resources, digest.RequestInfo.Resource) {
		return false
	}
	if !matchPolicyRuleAPIGroup(policyRule.APIGroups, digest.RequestInfo.APIGroup) {
		return false
	}
	if len(digest.RequestInfo.Namespace) == 0 {
		return policyRule.ClusterScope
	}
	return containsString(digest.RequestInfo.Namespace, policyRule.Namespaces, rmtypesv1alpha1.NamespaceEvery)
}

func matchesNonResourcePolicyRule(digest RequestDigest, policyRule rmtypesv1alpha1.NonResourcePolicyRule) bool {
	if !matchPolicyRuleVerb(policyRule.Verbs, digest.RequestInfo.Verb) {
		return false
	}
	if digest.RequestInfo.IsResourceRequest {
		return false
	}
	return matchPolicyRuleNonResourceURL(policyRule.NonResourceURLs, digest.RequestInfo.Path)
}

func matchPolicyRuleVerb(policyRuleVerbs []string, requestVerb string) bool {
	return containsString(requestVerb, policyRuleVerbs, rmtypesv1alpha1.VerbAll)
}

func matchPolicyRuleNonResourceURL(policyRuleRequestURLs []string, requestPath string) bool {
	return containsString(requestPath, policyRuleRequestURLs, rmtypesv1alpha1.NonResourceAll)
}

func matchPolicyRuleAPIGroup(policyRuleAPIGroups []string, requestAPIGroup string) bool {
	return containsString(requestAPIGroup, policyRuleAPIGroups, rmtypesv1alpha1.APIGroupAll)
}

func matchPolicyRuleResource(policyRuleRequestResources []string, requestResource string) bool {
	return containsString(requestResource, policyRuleRequestResources, rmtypesv1alpha1.ResourceAll)
}

func containsString(x string, list []string, wildcard string) bool {
	if len(list) == 1 && list[0] == wildcard {
		return true
	}
	for _, y := range list {
		if x == y {
			return true
		}
	}
	return false
}
