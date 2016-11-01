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

package petset

import (
	"fmt"
	"strconv"

	"k8s.io/kubernetes/pkg/api"
	apiErrors "k8s.io/kubernetes/pkg/api/errors"
	"k8s.io/kubernetes/pkg/apis/apps"
	"k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset"
	"k8s.io/kubernetes/pkg/client/record"
	"k8s.io/kubernetes/pkg/runtime"

	"github.com/golang/glog"
	"errors"
)

// petLifeCycleEvent is used to communicate high level actions the controller
// needs to take on a given pet. It's recorded in the pcb. The recognized values
// are listed below.
type petLifeCycleEvent string

const (
	syncPet   petLifeCycleEvent = "sync"
	deletePet petLifeCycleEvent = "delete"
	// updateRetries is the number of Get/Update cycles we perform when an
	// update fails.
	updateRetries = 3
	// StatefulSetInitAnnotation is an annotation which when set, indicates that the
	// pet has finished initializing itself.
	// TODO: Replace this with init container status.
	StatefulSetInitAnnotation = "pod.alpha.kubernetes.io/initialized"
	//StatefulSetReplicasAnnotation is used to communicate the current number of replicas
	//configured for the Stateful Set to the StatefulSet's Pods
	StatefulSetReplicasAnnotation = "pod.alpha.kubernetes.io/replicas"
)

// pcb is the control block used to transmit all updates about a single pet.
// It serves as the manifest for a single pet. Users must populate the pod
// and parent fields to pass it around safely.
type pcb struct {
	// pod is the desired pet pod.
	pod *api.Pod
	// pvcs is a list of desired persistent volume claims for the pet pod.
	pvcs []api.PersistentVolumeClaim
	// event is the lifecycle event associated with this update.
	event petLifeCycleEvent
	// id is the identity index of this pet.
	id string
	// parent is a pointer to the parent statefulset.
	parent *apps.StatefulSet
}

//Gets the number of replicas stored in pod's StatefulSetReplicasAnnotation. If an the annotation
//is not present, or if it contains an invalid value, -1 is returned. Otherwise, a 32 bit integer
//representation of the annotation's value is returned.
func getReplicas(pod *api.Pod) int32 {
	if pod == nil {
		return -1
	} else if replicaStr,exists := pod.Annotations[StatefulSetReplicasAnnotation]; !exists{
		glog.Errorf("Pod %s has no annotation for %s",pod.Name,StatefulSetReplicasAnnotation)
		return -1
	} else if replicas64, err := strconv.ParseInt(replicaStr,10,32); err != nil {
		glog.Errorf("Pod %s has invalid value %s for annotation %s",
			pod.Name,replicaStr,StatefulSetReplicasAnnotation)
		return -1
	} else {
		return int32(replicas64)
	}
}

//Sets pod's StatefulSetReplicasAnnotation to the base 10 string value of replicas.
func setReplicas(p *api.Pod, replicas int32) {
	p.Annotations[StatefulSetReplicasAnnotation] = strconv.FormatInt(int64(replicas),10)
}

//Updates pod to be equivalent to desired, and to contain the number same number of replicas as set.
//By equivalent, we mean that all of the IdentityMappers for StatefulSet produce equivalent values
//for both pod and desired. The returned bool is true if pod has been mutated or false if no
//mutation occurs. If the returned error is not nil, the returned bool is false.
func update (set *apps.StatefulSet, id string, pod *api.Pod, desired *api.Pod) (bool,error) {
	//preconditions
	if set == nil {
		return false,errors.New("nil StatefulSet parameter")
	} else if (pod == nil){
		return false,errors.New("nil source supplied")
	} else if (desired == nil ) {
		return false,errors.New("nil target supplied")
	}

	//collect mappers that violate the equality condition
	mappers := newIdentityMappers(set)
	violated := make([]identityMapper,0,len(mappers))
	for _,mapper := range(mappers){
		if mapper.Identity(pod) != mapper.Identity(desired) {
			violated = append(violated,mapper)
		}
	}

	//return false if no mutations are necessary
	if len(violated) == 0 && getReplicas(pod) == set.Spec.Replicas {
		return false, nil
	}

	for _, mapper := range (violated) {
		mapper.SetIdentity(id, pod)
	}
	setReplicas(pod,set.Spec.Replicas)
	return true, nil

}

// pvcClient is a client for managing persistent volume claims.
type pvcClient interface {
	// DeletePVCs deletes the pvcs in the given pcb.
	DeletePVCs(*pcb) error
	// SyncPVCs creates/updates pvcs in the given pcb.
	SyncPVCs(*pcb) error
}

// petSyncer syncs a single pet.
type petSyncer struct {
	petClient

	// blockingPet is an unhealthy pet either from this iteration or a previous
	// iteration, either because it is not yet Running, or being Deleted, that
	// prevents other creates/deletions.
	blockingPet *pcb
}

// errUnhealthyPet is returned when a we either know for sure a pet is unhealthy,
// or don't know its state but assume it is unhealthy. It's used as a signal to the caller for further operations like updating status.replicas.
// This is not a fatal error.
type errUnhealthyPet string

func (e errUnhealthyPet) Error() string {
	return string(e)
}

// Sync syncs the given pet.
func (p *petSyncer) Sync(pet *pcb) error {
	if pet == nil {
		return nil
	}
	realPet, exists, err := p.Get(pet)
	if err != nil {
		return err
	}
	// There is not constraint except quota on the number of pvcs created.
	// This is done per pet so we get a working cluster ASAP, even if user
	// runs out of quota.
	if err := p.SyncPVCs(pet); err != nil {
		return err
	}
	// if pet failed - we need to remove old one because of consistent naming
	if exists && realPet.pod.Status.Phase == api.PodFailed {
		glog.V(4).Infof("Delete evicted pod %v", realPet.pod.Name)
		if err := p.petClient.Delete(realPet); err != nil {
			return err
		}
	} else if exists {
		if !p.isHealthy(realPet.pod) {
			glog.Infof("StatefulSet %v waiting on unhealthy pet %v", pet.parent.Name, realPet.pod.Name)
		}
		return p.Update(realPet, pet)
	}
	if p.blockingPet != nil {
		message := errUnhealthyPet(fmt.Sprintf("Create of %v in StatefulSet %v blocked by unhealthy pet %v", pet.pod.Name, pet.parent.Name, p.blockingPet.pod.Name))
		glog.Info(message)
		return message
	}
	// This is counted as a create, even if it fails. We can't skip indices
	// because some pets might allocate a special role to earlier indices.
	// The returned error will force a requeue.
	// TODO: What's the desired behavior if pet-0 is deleted while pet-1 is
	// not yet healthy? currently pet-0 will wait till pet-1 is healthy,
	// this feels safer, but might lead to deadlock.
	p.blockingPet = pet
	if err := p.Create(pet); err != nil {
		return err
	}
	return nil
}

// Delete deletes the given pet, if no other pet in the statefulset is blocking a
// scale event.
func (p *petSyncer) Delete(pet *pcb) error {
	if pet == nil {
		return nil
	}
	realPet, exists, err := p.Get(pet)
	if err != nil {
		return err
	}
	if !exists {
		return nil
	}
	if p.blockingPet != nil {
		glog.Infof("Delete of %v in StatefulSet %v blocked by unhealthy pet %v", realPet.pod.Name, pet.parent.Name, p.blockingPet.pod.Name)
		return nil
	}
	// This is counted as a delete, even if it fails.
	// The returned error will force a requeue.
	p.blockingPet = realPet
	if !p.isDying(realPet.pod) {
		glog.Infof("StatefulSet %v deleting pet %v", pet.parent.Name, pet.pod.Name)
		return p.petClient.Delete(pet)
	}
	glog.Infof("StatefulSet %v waiting on pet %v to die in %v", pet.parent.Name, realPet.pod.Name, realPet.pod.DeletionTimestamp)
	return nil
}

// petClient is a client for managing pets.
type petClient interface {
	pvcClient
	petHealthChecker
	Delete(*pcb) error
	Get(*pcb) (*pcb, bool, error)
	Create(*pcb) error
	Update(*pcb, *pcb) error
}

// apiServerPetClient is a statefulset aware Kubernetes client.
type apiServerPetClient struct {
	c        internalclientset.Interface
	recorder record.EventRecorder
	petHealthChecker
}

// Get gets the pet in the pcb from the apiserver.
func (p *apiServerPetClient) Get(pet *pcb) (*pcb, bool, error) {
	ns := pet.parent.Namespace
	pod, err := p.c.Core().Pods(ns).Get(pet.pod.Name)
	if apiErrors.IsNotFound(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}
	realPet := *pet
	realPet.pod = pod
	return &realPet, true, nil
}

// Delete deletes the pet in the pcb from the apiserver.
func (p *apiServerPetClient) Delete(pet *pcb) error {
	err := p.c.Core().Pods(pet.parent.Namespace).Delete(pet.pod.Name, nil)
	if apiErrors.IsNotFound(err) {
		err = nil
	}
	p.event(pet.parent, "Delete", fmt.Sprintf("pet: %v", pet.pod.Name), err)
	return err
}

// Create creates the pet in the pcb.
func (p *apiServerPetClient) Create(pet *pcb) error {
	_, err := p.c.Core().Pods(pet.parent.Namespace).Create(pet.pod)
	p.event(pet.parent, "Create", fmt.Sprintf("pet: %v", pet.pod.Name), err)
	return err
}

// Update updates the pet in the 'pet' pcb to match the pet in the 'expectedPet' pcb.
// If the pod object of a pet which to be updated has been changed in server side, we
// will get the actual value and set pet identity before retries.
func (p *apiServerPetClient) Update(pet *pcb, expectedPet *pcb) error {
	pc := p.c.Core().Pods(pet.parent.Namespace)

	for i := 0; ; i++ {
		mutated, err := update(pet.parent,pet.id,pet.pod,expectedPet.pod)

		if err != nil || !mutated {
			return err
		}
		glog.Infof("Resetting pet %v/%v to match StatefulSet %v spec", pet.pod.Namespace, pet.pod.Name, pet.parent.Name)
		_, err = pc.Update(pet.pod)
		if err == nil || i >= updateRetries {
			return err
		}
		pod, getErr := pc.Get(pet.pod.Name)
		if getErr != nil {
			return getErr
		}
		pet.pod = pod
	}
}

// DeletePVCs should delete PVCs, when implemented.
func (p *apiServerPetClient) DeletePVCs(pet *pcb) error {
	// TODO: Implement this when we delete pvcs.
	return nil
}

func (p *apiServerPetClient) getPVC(pvcName, pvcNamespace string) (*api.PersistentVolumeClaim, error) {
	pvc, err := p.c.Core().PersistentVolumeClaims(pvcNamespace).Get(pvcName)
	return pvc, err
}

func (p *apiServerPetClient) createPVC(pvc *api.PersistentVolumeClaim) error {
	_, err := p.c.Core().PersistentVolumeClaims(pvc.Namespace).Create(pvc)
	return err
}

// SyncPVCs syncs pvcs in the given pcb.
func (p *apiServerPetClient) SyncPVCs(pet *pcb) error {
	errmsg := ""
	// Create new claims.
	for i, pvc := range pet.pvcs {
		_, err := p.getPVC(pvc.Name, pet.parent.Namespace)
		if err != nil {
			if apiErrors.IsNotFound(err) {
				var err error
				if err = p.createPVC(&pet.pvcs[i]); err != nil {
					errmsg += fmt.Sprintf("Failed to create %v: %v", pvc.Name, err)
				}
				p.event(pet.parent, "Create", fmt.Sprintf("pvc: %v", pvc.Name), err)
			} else {
				errmsg += fmt.Sprintf("Error trying to get pvc %v, %v.", pvc.Name, err)
			}
		}
		// TODO: Check resource requirements and accessmodes, update if necessary
	}
	if len(errmsg) != 0 {
		return fmt.Errorf("%v", errmsg)
	}
	return nil
}

// event formats an event for the given runtime object.
func (p *apiServerPetClient) event(obj runtime.Object, reason, msg string, err error) {
	if err != nil {
		p.recorder.Eventf(obj, api.EventTypeWarning, fmt.Sprintf("Failed%v", reason), fmt.Sprintf("%v, error: %v", msg, err))
	} else {
		p.recorder.Eventf(obj, api.EventTypeNormal, fmt.Sprintf("Successful%v", reason), msg)
	}
}

// petHealthChecker is an interface to check pet health. It makes a boolean
// decision based on the given pod.
type petHealthChecker interface {
	isHealthy(*api.Pod) bool
	isDying(*api.Pod) bool
}

// defaultPetHealthChecks does basic health checking.
// It doesn't update, probe or get the pod.
type defaultPetHealthChecker struct{}

// isHealthy returns true if the pod is running and has the
// "pod.alpha.kubernetes.io/initialized" set to "true".
func (d *defaultPetHealthChecker) isHealthy(pod *api.Pod) bool {
	if pod == nil || pod.Status.Phase != api.PodRunning {
		return false
	}
	initialized, ok := pod.Annotations[StatefulSetInitAnnotation]
	if !ok {
		glog.Infof("StatefulSet pod %v in %v, waiting on annotation %v", api.PodRunning, pod.Name, StatefulSetInitAnnotation)
		return false
	}
	b, err := strconv.ParseBool(initialized)
	if err != nil {
		return false
	}
	return b && api.IsPodReady(pod)
}

// isDying returns true if the pod has a non-nil deletion timestamp. Since the
// timestamp can only decrease, once this method returns true for a given pet, it
// will never return false.
func (d *defaultPetHealthChecker) isDying(pod *api.Pod) bool {
	return pod != nil && pod.DeletionTimestamp != nil
}
