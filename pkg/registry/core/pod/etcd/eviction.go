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

package etcd

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/policy"
	policyclient "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic/registry"
	"k8s.io/kubernetes/pkg/runtime"
)

func newEvictionStorage(store *registry.Store, podDisruptionBudgetClient policyclient.PodDisruptionBudgetsGetter) *EvictionREST {
	return &EvictionREST{store: store, podDisruptionBudgetClient: podDisruptionBudgetClient}
}

// EvictionREST implements the REST endpoint for evicting pods from nodes when etcd is in use.
type EvictionREST struct {
	store                     *registry.Store
	podDisruptionBudgetClient policyclient.PodDisruptionBudgetsGetter
}

var _ = rest.Creater(&EvictionREST{})

// New creates a new eviction resource
func (r *EvictionREST) New() runtime.Object {
	return &policy.Eviction{}
}

// Create attempts to create a new eviction.  That is, it tries to evict a pod.
func (r *EvictionREST) Create(ctx api.Context, obj runtime.Object) (runtime.Object, error) {
	eviction := obj.(*policy.Eviction)

	obj, err := r.store.Get(ctx, eviction.Name)
	if err != nil {
		return nil, err
	}
	pod := obj.(*api.Pod)
	pdbs, err := r.getPodDisruptionBudgets(ctx, pod)
	if err != nil {
		return nil, err
	}

	if len(pdbs) > 1 {
		return &unversioned.Status{
			Status:  unversioned.StatusFailure,
			Message: "This pod has more than one PodDisruptionBudget, which the eviction subresource does not support.",
			Code:    500,
		}, nil
	} else if len(pdbs) == 1 {
		pdb := pdbs[0]
		// Try to verify-and-decrement

		// If it was false already, or if it becomes false during the course of our retries,
		// raise an error marked as a 429.
		ok, err := r.checkAndDecrement(pod.Namespace, pdb)
		if err != nil {
			return nil, err
		}

		if !ok {
			return &unversioned.Status{
				Status: unversioned.StatusFailure,
				// TODO(mml): Include some more details about why the eviction is disallowed.
				// Ideally any such text is generated by the DisruptionController (offline).
				Message: "Cannot evict pod as it would violate the pod's disruption budget.",
				Code:    429,
				// TODO(mml): Add a Retry-After header.  Once there are time-based
				// budgets, we can sometimes compute a sensible suggested value.  But
				// even without that, we can give a suggestion (10 minutes?) that
				// prevents well-behaved clients from hammering us.
			}, nil
		}
	}

	// At this point there was either no PDB or we succeded in decrementing

	// Try the delete
	_, err = r.store.Delete(ctx, eviction.Name, eviction.DeleteOptions)
	if err != nil {
		return nil, err
	}

	// Success!
	return &unversioned.Status{Status: unversioned.StatusSuccess}, nil
}

func (r *EvictionREST) checkAndDecrement(namespace string, pdb policy.PodDisruptionBudget) (ok bool, err error) {
	if !pdb.Status.PodDisruptionAllowed {
		return false, nil
	}

	pdb.Status.PodDisruptionAllowed = false
	if _, err := r.podDisruptionBudgetClient.PodDisruptionBudgets(namespace).Update(&pdb); err != nil {
		return false, err
	}

	return true, nil
}

// Returns any PDBs that match the pod.
// err is set if there's an error.
func (r *EvictionREST) getPodDisruptionBudgets(ctx api.Context, pod *api.Pod) (pdbs []policy.PodDisruptionBudget, err error) {
	if len(pod.Labels) == 0 {
		return
	}

	pdbList, err := r.podDisruptionBudgetClient.PodDisruptionBudgets(pod.Namespace).List(api.ListOptions{})
	if err != nil {
		return
	}

	for _, pdb := range pdbList.Items {
		if pdb.Namespace != pod.Namespace {
			continue
		}
		selector, err := unversioned.LabelSelectorAsSelector(pdb.Spec.Selector)
		if err != nil {
			continue
		}
		// If a PDB with a nil or empty selector creeps in, it should match nothing, not everything.
		if selector.Empty() || !selector.Matches(labels.Set(pod.Labels)) {
			continue
		}

		pdbs = append(pdbs, pdb)
	}

	return pdbs, nil
}
