/*
Copyright 2015 The Kubernetes Authors.

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

package v1_test

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	apiequality "k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/diff"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/v1"
	storageutil "k8s.io/kubernetes/pkg/apis/storage/util"
)

func TestPodLogOptions(t *testing.T) {
	sinceSeconds := int64(1)
	sinceTime := metav1.NewTime(time.Date(2000, 1, 1, 12, 34, 56, 0, time.UTC).Local())
	tailLines := int64(2)
	limitBytes := int64(3)

	versionedLogOptions := &v1.PodLogOptions{
		Container:    "mycontainer",
		Follow:       true,
		Previous:     true,
		SinceSeconds: &sinceSeconds,
		SinceTime:    &sinceTime,
		Timestamps:   true,
		TailLines:    &tailLines,
		LimitBytes:   &limitBytes,
	}
	unversionedLogOptions := &api.PodLogOptions{
		Container:    "mycontainer",
		Follow:       true,
		Previous:     true,
		SinceSeconds: &sinceSeconds,
		SinceTime:    &sinceTime,
		Timestamps:   true,
		TailLines:    &tailLines,
		LimitBytes:   &limitBytes,
	}
	expectedParameters := url.Values{
		"container":    {"mycontainer"},
		"follow":       {"true"},
		"previous":     {"true"},
		"sinceSeconds": {"1"},
		"sinceTime":    {"2000-01-01T12:34:56Z"},
		"timestamps":   {"true"},
		"tailLines":    {"2"},
		"limitBytes":   {"3"},
	}

	codec := runtime.NewParameterCodec(api.Scheme)

	// unversioned -> query params
	{
		actualParameters, err := codec.EncodeParameters(unversionedLogOptions, v1.SchemeGroupVersion)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(actualParameters, expectedParameters) {
			t.Fatalf("Expected\n%#v\ngot\n%#v", expectedParameters, actualParameters)
		}
	}

	// versioned -> query params
	{
		actualParameters, err := codec.EncodeParameters(versionedLogOptions, v1.SchemeGroupVersion)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(actualParameters, expectedParameters) {
			t.Fatalf("Expected\n%#v\ngot\n%#v", expectedParameters, actualParameters)
		}
	}

	// query params -> versioned
	{
		convertedLogOptions := &v1.PodLogOptions{}
		err := codec.DecodeParameters(expectedParameters, v1.SchemeGroupVersion, convertedLogOptions)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(convertedLogOptions, versionedLogOptions) {
			t.Fatalf("Unexpected deserialization:\n%s", diff.ObjectGoPrintSideBySide(versionedLogOptions, convertedLogOptions))
		}
	}

	// query params -> unversioned
	{
		convertedLogOptions := &api.PodLogOptions{}
		err := codec.DecodeParameters(expectedParameters, v1.SchemeGroupVersion, convertedLogOptions)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(convertedLogOptions, unversionedLogOptions) {
			t.Fatalf("Unexpected deserialization:\n%s", diff.ObjectGoPrintSideBySide(unversionedLogOptions, convertedLogOptions))
		}
	}
}

// TestPodSpecConversion tests that ServiceAccount is an alias for
// ServiceAccountName.
func TestPodSpecConversion(t *testing.T) {
	name, other := "foo", "bar"

	// Test internal -> v1. Should have both alias (DeprecatedServiceAccount)
	// and new field (ServiceAccountName).
	i := &api.PodSpec{
		ServiceAccountName: name,
	}
	v := v1.PodSpec{}
	if err := api.Scheme.Convert(i, &v, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if v.ServiceAccountName != name {
		t.Fatalf("want v1.ServiceAccountName %q, got %q", name, v.ServiceAccountName)
	}
	if v.DeprecatedServiceAccount != name {
		t.Fatalf("want v1.DeprecatedServiceAccount %q, got %q", name, v.DeprecatedServiceAccount)
	}

	// Test v1 -> internal. Either DeprecatedServiceAccount, ServiceAccountName,
	// or both should translate to ServiceAccountName. ServiceAccountName wins
	// if both are set.
	testCases := []*v1.PodSpec{
		// New
		{ServiceAccountName: name},
		// Alias
		{DeprecatedServiceAccount: name},
		// Both: same
		{ServiceAccountName: name, DeprecatedServiceAccount: name},
		// Both: different
		{ServiceAccountName: name, DeprecatedServiceAccount: other},
	}
	for k, v := range testCases {
		got := api.PodSpec{}
		err := api.Scheme.Convert(v, &got, nil)
		if err != nil {
			t.Fatalf("unexpected error for case %d: %v", k, err)
		}
		if got.ServiceAccountName != name {
			t.Fatalf("want api.ServiceAccountName %q, got %q", name, got.ServiceAccountName)
		}
	}
}

func TestResourceListConversion(t *testing.T) {
	bigMilliQuantity := resource.NewQuantity(resource.MaxMilliValue, resource.DecimalSI)
	bigMilliQuantity.Add(resource.MustParse("12345m"))

	tests := []struct {
		input    v1.ResourceList
		expected api.ResourceList
	}{
		{ // No changes necessary.
			input: v1.ResourceList{
				v1.ResourceMemory:  resource.MustParse("30M"),
				v1.ResourceCPU:     resource.MustParse("100m"),
				v1.ResourceStorage: resource.MustParse("1G"),
			},
			expected: api.ResourceList{
				api.ResourceMemory:  resource.MustParse("30M"),
				api.ResourceCPU:     resource.MustParse("100m"),
				api.ResourceStorage: resource.MustParse("1G"),
			},
		},
		{ // Nano-scale values should be rounded up to milli-scale.
			input: v1.ResourceList{
				v1.ResourceCPU:    resource.MustParse("3.000023m"),
				v1.ResourceMemory: resource.MustParse("500.000050m"),
			},
			expected: api.ResourceList{
				api.ResourceCPU:    resource.MustParse("4m"),
				api.ResourceMemory: resource.MustParse("501m"),
			},
		},
		{ // Large values should still be accurate.
			input: v1.ResourceList{
				v1.ResourceCPU:     *bigMilliQuantity.Copy(),
				v1.ResourceStorage: *bigMilliQuantity.Copy(),
			},
			expected: api.ResourceList{
				api.ResourceCPU:     *bigMilliQuantity.Copy(),
				api.ResourceStorage: *bigMilliQuantity.Copy(),
			},
		},
	}

	for i, test := range tests {
		output := api.ResourceList{}
		err := api.Scheme.Convert(&test.input, &output, nil)
		if err != nil {
			t.Fatalf("unexpected error for case %d: %v", i, err)
		}
		if !apiequality.Semantic.DeepEqual(test.expected, output) {
			t.Errorf("unexpected conversion for case %d: Expected %+v; Got %+v", i, test.expected, output)
		}
	}
}

type pvTestSpec struct {
	annotations      map[string]string
	storageClassName string
}

func TestPersistentVolumeV1ToApiConversion(t *testing.T) {
	tests := []struct {
		input, expected pvTestSpec
	}{
		{ // No changes
			input:    pvTestSpec{},
			expected: pvTestSpec{},
		},
		{ // No changes
			input: pvTestSpec{
				storageClassName: "foo",
			},
			expected: pvTestSpec{
				storageClassName: "foo",
			},
		},
		{ // Annotation is copied to class
			input: pvTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
			},
			expected: pvTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: "foo",
			},
		},
		{ // Annotation does not overwrite existing class
			input: pvTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: "bar",
			},
			expected: pvTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: "bar",
			},
		},
	}

	for i, test := range tests {
		input := v1.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: test.input.annotations,
			},
			Spec: v1.PersistentVolumeSpec{
				StorageClassName:              test.input.storageClassName,
				PersistentVolumeReclaimPolicy: v1.PersistentVolumeReclaimRetain,
			},
			Status: v1.PersistentVolumeStatus{
				Phase: v1.VolumePending,
			},
		}
		expected := api.PersistentVolume{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: test.expected.annotations,
			},
			Spec: api.PersistentVolumeSpec{
				StorageClassName:              test.expected.storageClassName,
				PersistentVolumeReclaimPolicy: api.PersistentVolumeReclaimRetain,
			},
			Status: api.PersistentVolumeStatus{
				Phase: api.VolumePending,
			},
		}

		output := api.PersistentVolume{}

		err := api.Scheme.Convert(&input, &output, nil)
		if err != nil {
			t.Fatalf("unexpected error for case %d: %v", i, err)
		}
		if !apiequality.Semantic.DeepEqual(expected, output) {
			t.Errorf("unexpected conversion for case %d: Expected %+v; Got %+v", i, expected, output)
		}
	}
}

type claimTestSpec struct {
	annotations      map[string]string
	storageClassName *string
}

func TestPersistentVolumeClaimV1ToApiConversion(t *testing.T) {
	fooClass := "foo"
	barClass := "bar"
	emptyClass := ""

	tests := []struct {
		input, expected claimTestSpec
	}{
		{ // No changes
			input:    claimTestSpec{},
			expected: claimTestSpec{},
		},
		{ // No changes
			input:    claimTestSpec{storageClassName: &emptyClass},
			expected: claimTestSpec{storageClassName: &emptyClass},
		},
		{ // No changes
			input: claimTestSpec{
				storageClassName: &fooClass,
			},
			expected: claimTestSpec{
				storageClassName: &fooClass,
			},
		},
		{ // Annotation is copied to class
			input: claimTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
			},
			expected: claimTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: &fooClass,
			},
		},
		{ // Annotation does not overwrite existing class
			input: claimTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: &barClass,
			},
			expected: claimTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: &barClass,
			},
		},
		{ // Annotation does not overwrite existing empty class
			input: claimTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: &emptyClass,
			},
			expected: claimTestSpec{
				annotations: map[string]string{
					storageutil.BetaStorageClassAnnotation: "foo",
				},
				storageClassName: &emptyClass,
			},
		},
	}

	for i, test := range tests {
		input := v1.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: test.input.annotations,
			},
			Spec: v1.PersistentVolumeClaimSpec{
				StorageClassName: test.input.storageClassName,
			},
			Status: v1.PersistentVolumeClaimStatus{
				Phase: v1.ClaimPending,
			},
		}
		expected := api.PersistentVolumeClaim{
			ObjectMeta: metav1.ObjectMeta{
				Annotations: test.expected.annotations,
			},
			Spec: api.PersistentVolumeClaimSpec{
				StorageClassName: test.expected.storageClassName,
			},
			Status: api.PersistentVolumeClaimStatus{
				Phase: api.ClaimPending,
			},
		}

		output := api.PersistentVolumeClaim{}

		err := api.Scheme.Convert(&input, &output, nil)
		if err != nil {
			t.Fatalf("unexpected error for case %d: %v", i, err)
		}
		if !apiequality.Semantic.DeepEqual(expected, output) {
			t.Errorf("unexpected conversion for case %d: Expected %+v; Got %+v", i, expected, output)
		}
	}
}
