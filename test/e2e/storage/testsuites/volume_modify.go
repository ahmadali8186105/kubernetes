/*
Copyright 2024 The Kubernetes Authors.

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

package testsuites

import (
	"context"
	"encoding/json"
	"time"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	storagev1alpha1 "k8s.io/api/storage/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/errors"
	clientset "k8s.io/client-go/kubernetes"
	"k8s.io/kubernetes/pkg/features"
	"k8s.io/kubernetes/test/e2e/framework"
	e2epod "k8s.io/kubernetes/test/e2e/framework/pod"
	e2eskipper "k8s.io/kubernetes/test/e2e/framework/skipper"
	e2evolume "k8s.io/kubernetes/test/e2e/framework/volume"
	storageframework "k8s.io/kubernetes/test/e2e/storage/framework"
	admissionapi "k8s.io/pod-security-admission/api"
)

const (
	modifyPollInterval               = 2 * time.Second
	setVACWaitPeriod                 = 30 * time.Second
	modifyingConditionSyncWaitPeriod = 2 * time.Minute
	modifyVolumeWaitPeriod           = 10 * time.Minute
	vacCleanupWaitPeriod             = 30 * time.Second
)

type volumeModifyTestSuite struct {
	tsInfo storageframework.TestSuiteInfo
}

// InitCustomVolumeModifyTestSuite returns volumeModifyTestSuite that implements TestSuite interface
// using custom test patterns
func InitCustomVolumeModifyTestSuite(patterns []storageframework.TestPattern) storageframework.TestSuite {
	return &volumeModifyTestSuite{
		tsInfo: storageframework.TestSuiteInfo{
			Name:         "volume-modify",
			TestPatterns: patterns,
			SupportedSizeRange: e2evolume.SizeRange{
				Min: "1Gi",
			},
			TestTags: []interface{}{framework.WithFeatureGate(features.VolumeAttributesClass)},
		},
	}
}

// InitVolumeModifyTestSuite returns volumeModifyTestSuite that implements TestSuite interface
// using testsuite default patterns
func InitVolumeModifyTestSuite() storageframework.TestSuite {
	patterns := []storageframework.TestPattern{
		storageframework.DefaultFsDynamicPV,
		storageframework.BlockVolModeDynamicPV,
		storageframework.NtfsDynamicPV,
	}
	return InitCustomVolumeModifyTestSuite(patterns)
}

func (v *volumeModifyTestSuite) GetTestSuiteInfo() storageframework.TestSuiteInfo {
	return v.tsInfo
}

func (v *volumeModifyTestSuite) SkipUnsupportedTests(driver storageframework.TestDriver, pattern storageframework.TestPattern) {
	_, ok := driver.(storageframework.VolumeAttributesClassTestDriver)
	if !ok {
		e2eskipper.Skipf("Driver %q does not support VolumeAttributesClass tests - skipping", driver.GetDriverInfo().Name)
	}
	// Skip block storage tests if the driver we are testing against does not support block volumes
	// TODO: This should be made generic so that it doesn't have to be re-written for every test that uses the 	BlockVolModeDynamicPV testcase
	if !driver.GetDriverInfo().Capabilities[storageframework.CapBlock] && pattern.VolMode == v1.PersistentVolumeBlock {
		e2eskipper.Skipf("Driver %q does not support block volume mode - skipping", driver.GetDriverInfo().Name)
	}
}

func (v *volumeModifyTestSuite) DefineTests(driver storageframework.TestDriver, pattern storageframework.TestPattern) {
	type local struct {
		config *storageframework.PerTestConfig

		resource *storageframework.VolumeResource
		vac      *storagev1alpha1.VolumeAttributesClass
	}
	var l local

	// Beware that it also registers an AfterEach which renders f unusable. Any code using
	// f must run inside an It or Context callback.
	f := framework.NewFrameworkWithCustomTimeouts("volume-modify", storageframework.GetDriverTimeouts(driver))
	f.NamespacePodSecurityLevel = admissionapi.LevelPrivileged

	init := func(ctx context.Context, createVolumeWithVAC bool) {
		l = local{}

		l.config = driver.PrepareTest(ctx, f)
		vacDriver, _ := driver.(storageframework.VolumeAttributesClassTestDriver)
		l.vac = vacDriver.GetVolumeAttributesClass(ctx, l.config)

		if l.vac == nil {
			e2eskipper.Skipf("Driver %q returned nil VolumeAttributesClass - skipping", driver.GetDriverInfo().Name)
		}

		ginkgo.By("Creating VolumeAttributesClass")
		_, err := f.ClientSet.StorageV1alpha1().VolumeAttributesClasses().Create(ctx, l.vac, metav1.CreateOptions{})
		framework.ExpectNoError(err, "While creating VolumeAttributesClass")

		ginkgo.By("Creating volume")
		testVolumeSizeRange := v.GetTestSuiteInfo().SupportedSizeRange
		if createVolumeWithVAC {
			l.resource = storageframework.CreateVolumeResourceWithVAC(ctx, driver, l.config, pattern, testVolumeSizeRange, &l.vac.Name)
		} else {
			l.resource = storageframework.CreateVolumeResource(ctx, driver, l.config, pattern, testVolumeSizeRange)
		}
	}

	cleanup := func(ctx context.Context) {
		var errs []error
		if l.resource != nil {
			ginkgo.By("Deleting VolumeResource")
			errs = append(errs, l.resource.CleanupResource(ctx))
			l.resource = nil
		}

		if l.vac != nil {
			ginkgo.By("Deleting VAC")
			CleanupVAC(ctx, l.vac, f.ClientSet, vacCleanupWaitPeriod)
			l.vac = nil
		}

		framework.ExpectNoError(errors.NewAggregate(errs), "While cleaning up")
	}

	ginkgo.It("should create a volume with VAC", func(ctx context.Context) {
		init(ctx, true /* volume created with VAC */)
		ginkgo.DeferCleanup(cleanup)

		ginkgo.By("Creating a pod with dynamically provisioned volume")
		podConfig := e2epod.Config{
			NS:            f.Namespace.Name,
			PVCs:          []*v1.PersistentVolumeClaim{l.resource.Pvc},
			SeLinuxLabel:  e2epod.GetLinuxLabel(),
			NodeSelection: l.config.ClientNodeSelection,
			ImageID:       e2epod.GetDefaultTestImageID(),
		}
		pod, err := e2epod.CreateSecPodWithNodeSelection(ctx, f.ClientSet, &podConfig, f.Timeouts.PodStart)
		ginkgo.DeferCleanup(e2epod.DeletePodWithWait, f.ClientSet, pod)
		framework.ExpectNoError(err, "While creating test pod with VAC")

		createdPVC, err := f.ClientSet.CoreV1().PersistentVolumeClaims(f.Namespace.Name).Get(ctx, l.resource.Pvc.Name, metav1.GetOptions{})
		framework.ExpectNoError(err, "While getting created PVC")
		// Check VAC matches on created PVC, but not current VAC in status
		gomega.Expect(vacMatches(createdPVC, l.vac.Name, false)).To(gomega.BeTrueBecause("Created PVC should match expected VAC"))
	})

	ginkgo.It("should modify volume with no VAC", func(ctx context.Context) {
		init(ctx, false /* volume created without VAC */)
		ginkgo.DeferCleanup(cleanup)

		var err error
		ginkgo.By("Creating a pod with dynamically provisioned volume")
		podConfig := e2epod.Config{
			NS:            f.Namespace.Name,
			PVCs:          []*v1.PersistentVolumeClaim{l.resource.Pvc},
			SeLinuxLabel:  e2epod.GetLinuxLabel(),
			NodeSelection: l.config.ClientNodeSelection,
			ImageID:       e2epod.GetDefaultTestImageID(),
		}
		pod, err := e2epod.CreateSecPodWithNodeSelection(ctx, f.ClientSet, &podConfig, f.Timeouts.PodStart)
		ginkgo.DeferCleanup(e2epod.DeletePodWithWait, f.ClientSet, pod)
		framework.ExpectNoError(err, "While creating pod for modifying")

		ginkgo.By("Modifying PVC via VAC")
		newPVC := SetPVCVACName(ctx, l.resource.Pvc, l.vac.Name, f.ClientSet, setVACWaitPeriod)
		l.resource.Pvc = newPVC
		gomega.Expect(l.resource.Pvc).NotTo(gomega.BeNil())

		ginkgo.By("Waiting for modification to finish")
		WaitForVolumeModification(ctx, l.resource.Pvc, f.ClientSet, modifyVolumeWaitPeriod)

		pvcConditions := l.resource.Pvc.Status.Conditions
		gomega.Expect(pvcConditions).To(gomega.BeEmpty(), "PVC should not have conditions")
	})

	ginkgo.It("should modify volume that already has a VAC", func(ctx context.Context) {
		init(ctx, true /* volume created with VAC */)
		ginkgo.DeferCleanup(cleanup)

		vacDriver, _ := driver.(storageframework.VolumeAttributesClassTestDriver)
		newVAC := vacDriver.GetVolumeAttributesClass(ctx, l.config)
		gomega.Expect(newVAC).NotTo(gomega.BeNil())
		_, err := f.ClientSet.StorageV1alpha1().VolumeAttributesClasses().Create(ctx, newVAC, metav1.CreateOptions{})
		framework.ExpectNoError(err, "While creating new VolumeAttributesClass")
		ginkgo.DeferCleanup(CleanupVAC, newVAC, f.ClientSet, vacCleanupWaitPeriod)

		ginkgo.By("Creating a pod with dynamically provisioned volume")
		podConfig := e2epod.Config{
			NS:            f.Namespace.Name,
			PVCs:          []*v1.PersistentVolumeClaim{l.resource.Pvc},
			SeLinuxLabel:  e2epod.GetLinuxLabel(),
			NodeSelection: l.config.ClientNodeSelection,
			ImageID:       e2epod.GetDefaultTestImageID(),
		}
		pod, err := e2epod.CreateSecPodWithNodeSelection(ctx, f.ClientSet, &podConfig, f.Timeouts.PodStart)
		ginkgo.DeferCleanup(e2epod.DeletePodWithWait, f.ClientSet, pod)
		framework.ExpectNoError(err, "While creating pod for modifying")

		ginkgo.By("Modifying PVC via VAC")
		newPVC := SetPVCVACName(ctx, l.resource.Pvc, newVAC.Name, f.ClientSet, setVACWaitPeriod)
		l.resource.Pvc = newPVC
		gomega.Expect(l.resource.Pvc).NotTo(gomega.BeNil())

		ginkgo.By("Waiting for modification to finish")
		WaitForVolumeModification(ctx, l.resource.Pvc, f.ClientSet, modifyVolumeWaitPeriod)

		pvcConditions := l.resource.Pvc.Status.Conditions
		gomega.Expect(pvcConditions).To(gomega.BeEmpty(), "PVC should not have conditions")
	})
}

// SetPVCVACName sets the VolumeAttributesClassName on a PVC object
func SetPVCVACName(ctx context.Context, origPVC *v1.PersistentVolumeClaim, name string, c clientset.Interface, timeout time.Duration) *v1.PersistentVolumeClaim {
	pvcName := origPVC.Name
	var patchedPVC *v1.PersistentVolumeClaim

	gomega.Eventually(ctx, func(g gomega.Gomega) {
		var err error
		patch := []map[string]interface{}{{"op": "replace", "path": "/spec/volumeAttributesClassName", "value": name}}
		patchBytes, _ := json.Marshal(patch)

		patchedPVC, err = c.CoreV1().PersistentVolumeClaims(origPVC.Namespace).Patch(ctx, pvcName, types.JSONPatchType, patchBytes, metav1.PatchOptions{})
		framework.ExpectNoError(err, "While patching PVC to add VAC name")
	}, timeout, modifyPollInterval).Should(gomega.Succeed())

	return patchedPVC
}

// WaitForVolumeModification waits for the volume to be modified
// The input PVC is assumed to have a VolumeAttributesClassName set
func WaitForVolumeModification(ctx context.Context, pvc *v1.PersistentVolumeClaim, c clientset.Interface, timeout time.Duration) {
	pvName := pvc.Spec.VolumeName
	gomega.Eventually(ctx, func(g gomega.Gomega) {
		pv, err := c.CoreV1().PersistentVolumes().Get(ctx, pvName, metav1.GetOptions{})
		framework.ExpectNoError(err, "While getting existing PV")
		g.Expect(pv.Spec.VolumeAttributesClassName).NotTo(gomega.BeNil())
		newPVC, err := c.CoreV1().PersistentVolumeClaims(pvc.Namespace).Get(ctx, pvc.Name, metav1.GetOptions{})
		framework.ExpectNoError(err, "While getting new PVC")
		g.Expect(vacMatches(newPVC, *pv.Spec.VolumeAttributesClassName, true)).To(gomega.BeTrueBecause("Modified PVC should match expected VAC"))
	}, timeout, modifyPollInterval).Should(gomega.Succeed())
}

func CleanupVAC(ctx context.Context, vac *storagev1alpha1.VolumeAttributesClass, c clientset.Interface, timeout time.Duration) {
	gomega.Eventually(ctx, func() error {
		return c.StorageV1alpha1().VolumeAttributesClasses().Delete(ctx, vac.Name, metav1.DeleteOptions{})
	}, timeout, modifyPollInterval).Should(gomega.BeNil())
}

func vacMatches(pvc *v1.PersistentVolumeClaim, expectedVac string, checkStatusCurrentVac bool) bool {
	// Check the following to ensure the VAC matches and that all pending modifications are complete:
	// 1. VAC Name matches Expected
	// 2. PVC Modify Volume status is either nil or has an empty status string
	// 3. PVC Status Current VAC Matches Expected (only if checkStatusCurrentVac is true)
	// (3) is only expected to be true after a VAC is modified, but not when a VAC is used to create a volume
	if pvc.Spec.VolumeAttributesClassName == nil || *pvc.Spec.VolumeAttributesClassName != expectedVac {
		return false
	}
	if pvc.Status.ModifyVolumeStatus != nil && (pvc.Status.ModifyVolumeStatus.Status != "" || pvc.Status.ModifyVolumeStatus.TargetVolumeAttributesClassName != expectedVac) {
		return false
	}
	if checkStatusCurrentVac {
		if pvc.Status.CurrentVolumeAttributesClassName == nil || *pvc.Status.CurrentVolumeAttributesClassName != expectedVac {
			return false
		}
	}

	return true
}
