// +build !ignore_autogenerated

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

// This file was automatically generated by: _output/bin/defaulter-gen --v 1 --logtostderr -i k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission/v1beta1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1alpha1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1beta2,k8s.io/kubernetes/pkg/apis/authentication/v1,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization/v1,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/autoscaling/v2beta1,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v1beta1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates/v1beta1,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/kubernetes/pkg/apis/events/v1beta1,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/networking/v1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac/v1,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac/v1beta1,k8s.io/kubernetes/pkg/apis/scheduling/v1alpha1,k8s.io/kubernetes/pkg/apis/settings/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1,k8s.io/kubernetes/pkg/apis/storage/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1beta1,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1 --extra-peer-dirs k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission/v1beta1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1alpha1,k8s.io/kubernetes/pkg/apis/admissionregistration/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1,k8s.io/kubernetes/pkg/apis/apps/v1beta1,k8s.io/kubernetes/pkg/apis/apps/v1beta2,k8s.io/kubernetes/pkg/apis/authentication/v1,k8s.io/kubernetes/pkg/apis/authentication/v1beta1,k8s.io/kubernetes/pkg/apis/authorization/v1,k8s.io/kubernetes/pkg/apis/authorization/v1beta1,k8s.io/kubernetes/pkg/apis/autoscaling/v1,k8s.io/kubernetes/pkg/apis/autoscaling/v2beta1,k8s.io/kubernetes/pkg/apis/batch/v1,k8s.io/kubernetes/pkg/apis/batch/v1beta1,k8s.io/kubernetes/pkg/apis/batch/v2alpha1,k8s.io/kubernetes/pkg/apis/certificates/v1beta1,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core/v1,k8s.io/kubernetes/pkg/apis/events/v1beta1,k8s.io/kubernetes/pkg/apis/extensions/v1beta1,k8s.io/kubernetes/pkg/apis/imagepolicy/v1alpha1,k8s.io/kubernetes/pkg/apis/networking/v1,k8s.io/kubernetes/pkg/apis/policy/v1beta1,k8s.io/kubernetes/pkg/apis/rbac/v1,k8s.io/kubernetes/pkg/apis/rbac/v1alpha1,k8s.io/kubernetes/pkg/apis/rbac/v1beta1,k8s.io/kubernetes/pkg/apis/scheduling/v1alpha1,k8s.io/kubernetes/pkg/apis/settings/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1,k8s.io/kubernetes/pkg/apis/storage/v1alpha1,k8s.io/kubernetes/pkg/apis/storage/v1beta1,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1 -O zz_generated.defaults

// This file was autogenerated by defaulter-gen. Do not edit it manually!

package v1

import (
	v1 "k8s.io/api/batch/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core_v1 "k8s.io/kubernetes/pkg/apis/core/v1"
)

// RegisterDefaults adds defaulters functions to the given scheme.
// Public to allow building arbitrary schemes.
// All generated defaulters are covering - they call all nested defaulters.
func RegisterDefaults(scheme *runtime.Scheme) error {
	scheme.AddTypeDefaultingFunc(&v1.Job{}, func(obj interface{}) { SetObjectDefaults_Job(obj.(*v1.Job)) })
	scheme.AddTypeDefaultingFunc(&v1.JobList{}, func(obj interface{}) { SetObjectDefaults_JobList(obj.(*v1.JobList)) })
	return nil
}

func SetObjectDefaults_Job(in *v1.Job) {
	SetDefaults_Job(in)
	core_v1.SetDefaults_PodSpec(&in.Spec.Template.Spec)
	for i := range in.Spec.Template.Spec.Volumes {
		a := &in.Spec.Template.Spec.Volumes[i]
		core_v1.SetDefaults_Volume(a)
		if a.VolumeSource.HostPath != nil {
			core_v1.SetDefaults_HostPathVolumeSource(a.VolumeSource.HostPath)
		}
		if a.VolumeSource.Secret != nil {
			core_v1.SetDefaults_SecretVolumeSource(a.VolumeSource.Secret)
		}
		if a.VolumeSource.ISCSI != nil {
			core_v1.SetDefaults_ISCSIVolumeSource(a.VolumeSource.ISCSI)
		}
		if a.VolumeSource.RBD != nil {
			core_v1.SetDefaults_RBDVolumeSource(a.VolumeSource.RBD)
		}
		if a.VolumeSource.DownwardAPI != nil {
			core_v1.SetDefaults_DownwardAPIVolumeSource(a.VolumeSource.DownwardAPI)
			for j := range a.VolumeSource.DownwardAPI.Items {
				b := &a.VolumeSource.DownwardAPI.Items[j]
				if b.FieldRef != nil {
					core_v1.SetDefaults_ObjectFieldSelector(b.FieldRef)
				}
			}
		}
		if a.VolumeSource.ConfigMap != nil {
			core_v1.SetDefaults_ConfigMapVolumeSource(a.VolumeSource.ConfigMap)
		}
		if a.VolumeSource.AzureDisk != nil {
			core_v1.SetDefaults_AzureDiskVolumeSource(a.VolumeSource.AzureDisk)
		}
		if a.VolumeSource.Projected != nil {
			core_v1.SetDefaults_ProjectedVolumeSource(a.VolumeSource.Projected)
			for j := range a.VolumeSource.Projected.Sources {
				b := &a.VolumeSource.Projected.Sources[j]
				if b.DownwardAPI != nil {
					for k := range b.DownwardAPI.Items {
						c := &b.DownwardAPI.Items[k]
						if c.FieldRef != nil {
							core_v1.SetDefaults_ObjectFieldSelector(c.FieldRef)
						}
					}
				}
			}
		}
		if a.VolumeSource.ScaleIO != nil {
			core_v1.SetDefaults_ScaleIOVolumeSource(a.VolumeSource.ScaleIO)
		}
	}
	for i := range in.Spec.Template.Spec.InitContainers {
		a := &in.Spec.Template.Spec.InitContainers[i]
		core_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			core_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					core_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		core_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		core_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			core_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				core_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			core_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				core_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					core_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					core_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
	for i := range in.Spec.Template.Spec.Containers {
		a := &in.Spec.Template.Spec.Containers[i]
		core_v1.SetDefaults_Container(a)
		for j := range a.Ports {
			b := &a.Ports[j]
			core_v1.SetDefaults_ContainerPort(b)
		}
		for j := range a.Env {
			b := &a.Env[j]
			if b.ValueFrom != nil {
				if b.ValueFrom.FieldRef != nil {
					core_v1.SetDefaults_ObjectFieldSelector(b.ValueFrom.FieldRef)
				}
			}
		}
		core_v1.SetDefaults_ResourceList(&a.Resources.Limits)
		core_v1.SetDefaults_ResourceList(&a.Resources.Requests)
		if a.LivenessProbe != nil {
			core_v1.SetDefaults_Probe(a.LivenessProbe)
			if a.LivenessProbe.Handler.HTTPGet != nil {
				core_v1.SetDefaults_HTTPGetAction(a.LivenessProbe.Handler.HTTPGet)
			}
		}
		if a.ReadinessProbe != nil {
			core_v1.SetDefaults_Probe(a.ReadinessProbe)
			if a.ReadinessProbe.Handler.HTTPGet != nil {
				core_v1.SetDefaults_HTTPGetAction(a.ReadinessProbe.Handler.HTTPGet)
			}
		}
		if a.Lifecycle != nil {
			if a.Lifecycle.PostStart != nil {
				if a.Lifecycle.PostStart.HTTPGet != nil {
					core_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PostStart.HTTPGet)
				}
			}
			if a.Lifecycle.PreStop != nil {
				if a.Lifecycle.PreStop.HTTPGet != nil {
					core_v1.SetDefaults_HTTPGetAction(a.Lifecycle.PreStop.HTTPGet)
				}
			}
		}
	}
}

func SetObjectDefaults_JobList(in *v1.JobList) {
	for i := range in.Items {
		a := &in.Items[i]
		SetObjectDefaults_Job(a)
	}
}
