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

// This file was automatically generated by: _output/bin/deepcopy-gen --v 1 --logtostderr -i k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm,k8s.io/kubernetes/cmd/kubeadm/app/apis/kubeadm/v1alpha1,k8s.io/kubernetes/cmd/kubeadm/app/phases/etcd/spec,k8s.io/kubernetes/pkg/apis/abac,k8s.io/kubernetes/pkg/apis/abac/v0,k8s.io/kubernetes/pkg/apis/abac/v1beta1,k8s.io/kubernetes/pkg/apis/admission,k8s.io/kubernetes/pkg/apis/admissionregistration,k8s.io/kubernetes/pkg/apis/apps,k8s.io/kubernetes/pkg/apis/authentication,k8s.io/kubernetes/pkg/apis/authorization,k8s.io/kubernetes/pkg/apis/autoscaling,k8s.io/kubernetes/pkg/apis/batch,k8s.io/kubernetes/pkg/apis/certificates,k8s.io/kubernetes/pkg/apis/componentconfig,k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1,k8s.io/kubernetes/pkg/apis/core,k8s.io/kubernetes/pkg/apis/extensions,k8s.io/kubernetes/pkg/apis/imagepolicy,k8s.io/kubernetes/pkg/apis/networking,k8s.io/kubernetes/pkg/apis/policy,k8s.io/kubernetes/pkg/apis/rbac,k8s.io/kubernetes/pkg/apis/scheduling,k8s.io/kubernetes/pkg/apis/settings,k8s.io/kubernetes/pkg/apis/storage,k8s.io/kubernetes/pkg/controller/garbagecollector/metaonly,k8s.io/kubernetes/pkg/kubectl/cmd/testing,k8s.io/kubernetes/pkg/kubectl/testing,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig,k8s.io/kubernetes/pkg/kubelet/apis/kubeletconfig/v1alpha1,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig,k8s.io/kubernetes/pkg/proxy/apis/kubeproxyconfig/v1alpha1,k8s.io/kubernetes/pkg/registry/rbac/reconciliation,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit,k8s.io/kubernetes/plugin/pkg/admission/eventratelimit/apis/eventratelimit/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction,k8s.io/kubernetes/plugin/pkg/admission/podtolerationrestriction/apis/podtolerationrestriction/v1alpha1,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota,k8s.io/kubernetes/plugin/pkg/admission/resourcequota/apis/resourcequota/v1alpha1,k8s.io/kubernetes/plugin/pkg/scheduler/api,k8s.io/kubernetes/plugin/pkg/scheduler/api/v1,k8s.io/kubernetes/vendor/k8s.io/api/admission/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/admissionregistration/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/admissionregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/apps/v1beta2,k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1,k8s.io/kubernetes/vendor/k8s.io/api/authentication/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/authorization/v1,k8s.io/kubernetes/vendor/k8s.io/api/authorization/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/autoscaling/v1,k8s.io/kubernetes/vendor/k8s.io/api/autoscaling/v2beta1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/batch/v2alpha1,k8s.io/kubernetes/vendor/k8s.io/api/certificates/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/core/v1,k8s.io/kubernetes/vendor/k8s.io/api/events/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/extensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/imagepolicy/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/networking/v1,k8s.io/kubernetes/vendor/k8s.io/api/policy/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/rbac/v1beta1,k8s.io/kubernetes/vendor/k8s.io/api/scheduling/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/settings/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/api/storage/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/examples/client-go/pkg/apis/cr/v1,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions,k8s.io/kubernetes/vendor/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/api/resource,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/internalversion,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1/unstructured,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/meta/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/apis/testapigroup/v1,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/labels,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/serializer/testing,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/runtime/testing,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/test,k8s.io/kubernetes/vendor/k8s.io/apimachinery/pkg/watch,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/admission/plugin/webhook/config/apis/webhookadmission/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/apiserver/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/audit/v1beta1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/endpoints/openapi/testing,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/endpoints/testing,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/registry/rest,k8s.io/kubernetes/vendor/k8s.io/apiserver/pkg/storage/testing,k8s.io/kubernetes/vendor/k8s.io/client-go/rest,k8s.io/kubernetes/vendor/k8s.io/client-go/scale/scheme,k8s.io/kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api,k8s.io/kubernetes/vendor/k8s.io/client-go/tools/clientcmd/api/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/apiserver/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/crd/apis/example/v1,k8s.io/kubernetes/vendor/k8s.io/code-generator/_examples/crd/apis/example2/v1,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration,k8s.io/kubernetes/vendor/k8s.io/kube-aggregator/pkg/apis/apiregistration/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/custom_metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/metrics/pkg/apis/metrics/v1beta1,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle,k8s.io/kubernetes/vendor/k8s.io/sample-apiserver/pkg/apis/wardle/v1alpha1,k8s.io/kubernetes/vendor/k8s.io/sample-controller/pkg/apis/samplecontroller/v1alpha1 --bounding-dirs k8s.io/kubernetes,k8s.io/api -O zz_generated.deepcopy

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package batch

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/kubernetes/pkg/apis/core"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJob) DeepCopyInto(out *CronJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJob.
func (in *CronJob) DeepCopy() *CronJob {
	if in == nil {
		return nil
	}
	out := new(CronJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobList) DeepCopyInto(out *CronJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CronJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobList.
func (in *CronJobList) DeepCopy() *CronJobList {
	if in == nil {
		return nil
	}
	out := new(CronJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CronJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobSpec) DeepCopyInto(out *CronJobSpec) {
	*out = *in
	if in.StartingDeadlineSeconds != nil {
		in, out := &in.StartingDeadlineSeconds, &out.StartingDeadlineSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.Suspend != nil {
		in, out := &in.Suspend, &out.Suspend
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	in.JobTemplate.DeepCopyInto(&out.JobTemplate)
	if in.SuccessfulJobsHistoryLimit != nil {
		in, out := &in.SuccessfulJobsHistoryLimit, &out.SuccessfulJobsHistoryLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.FailedJobsHistoryLimit != nil {
		in, out := &in.FailedJobsHistoryLimit, &out.FailedJobsHistoryLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobSpec.
func (in *CronJobSpec) DeepCopy() *CronJobSpec {
	if in == nil {
		return nil
	}
	out := new(CronJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CronJobStatus) DeepCopyInto(out *CronJobStatus) {
	*out = *in
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]core.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CronJobStatus.
func (in *CronJobStatus) DeepCopy() *CronJobStatus {
	if in == nil {
		return nil
	}
	out := new(CronJobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Job) DeepCopyInto(out *Job) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Job.
func (in *Job) DeepCopy() *Job {
	if in == nil {
		return nil
	}
	out := new(Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Job) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobCondition) DeepCopyInto(out *JobCondition) {
	*out = *in
	in.LastProbeTime.DeepCopyInto(&out.LastProbeTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobCondition.
func (in *JobCondition) DeepCopy() *JobCondition {
	if in == nil {
		return nil
	}
	out := new(JobCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobList) DeepCopyInto(out *JobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Job, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobList.
func (in *JobList) DeepCopy() *JobList {
	if in == nil {
		return nil
	}
	out := new(JobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobSpec) DeepCopyInto(out *JobSpec) {
	*out = *in
	if in.Parallelism != nil {
		in, out := &in.Parallelism, &out.Parallelism
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Completions != nil {
		in, out := &in.Completions, &out.Completions
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		if *in == nil {
			*out = nil
		} else {
			*out = new(int64)
			**out = **in
		}
	}
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.LabelSelector)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ManualSelector != nil {
		in, out := &in.ManualSelector, &out.ManualSelector
		if *in == nil {
			*out = nil
		} else {
			*out = new(bool)
			**out = **in
		}
	}
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobSpec.
func (in *JobSpec) DeepCopy() *JobSpec {
	if in == nil {
		return nil
	}
	out := new(JobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobStatus) DeepCopyInto(out *JobStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]JobCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobStatus.
func (in *JobStatus) DeepCopy() *JobStatus {
	if in == nil {
		return nil
	}
	out := new(JobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplate) DeepCopyInto(out *JobTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplate.
func (in *JobTemplate) DeepCopy() *JobTemplate {
	if in == nil {
		return nil
	}
	out := new(JobTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JobTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobTemplateSpec) DeepCopyInto(out *JobTemplateSpec) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobTemplateSpec.
func (in *JobTemplateSpec) DeepCopy() *JobTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(JobTemplateSpec)
	in.DeepCopyInto(out)
	return out
}
