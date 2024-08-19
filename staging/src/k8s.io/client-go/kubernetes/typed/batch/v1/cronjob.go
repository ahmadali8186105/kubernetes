/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	batchv1 "k8s.io/client-go/applyconfigurations/batch/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CronJobsGetter has a method to return a CronJobInterface.
// A group's client should implement this interface.
type CronJobsGetter interface {
	CronJobs(namespace string) CronJobInterface
}

// CronJobInterface has methods to work with CronJob resources.
type CronJobInterface interface {
	Create(ctx context.Context, cronJob *v1.CronJob, opts metav1.CreateOptions) (*v1.CronJob, error)
	Update(ctx context.Context, cronJob *v1.CronJob, opts metav1.UpdateOptions) (*v1.CronJob, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, cronJob *v1.CronJob, opts metav1.UpdateOptions) (*v1.CronJob, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.CronJob, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.CronJobList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.CronJob, err error)
	Apply(ctx context.Context, cronJob *batchv1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CronJob, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, cronJob *batchv1.CronJobApplyConfiguration, opts metav1.ApplyOptions) (result *v1.CronJob, err error)
	CronJobExpansion
}

// cronJobs implements CronJobInterface
type cronJobs struct {
	*gentype.ClientWithListAndApply[*v1.CronJob, *v1.CronJobList, *batchv1.CronJobApplyConfiguration]
}

// newCronJobs returns a CronJobs
func newCronJobs(c *BatchV1Client, namespace string) *cronJobs {
	return &cronJobs{
		gentype.NewClientWithListAndApply[*v1.CronJob, *v1.CronJobList, *batchv1.CronJobApplyConfiguration](
			"cronjobs",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.CronJob { return &v1.CronJob{} },
			func() *v1.CronJobList { return &v1.CronJobList{} },
			true,
		),
	}
}
