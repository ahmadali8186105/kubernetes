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

package rest

import (
	"testing"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/server/storage"
	"k8s.io/apiserver/pkg/storage/storagebackend"
	utilfeature "k8s.io/apiserver/pkg/util/feature"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
	"k8s.io/kubernetes/pkg/features"
)

func TestGetServersToValidate(t *testing.T) {
	tests := []struct {
		name     string
		gate     bool
		expected []string
	}{
		{
			name:     "DisableComponentStatusProbes enabled",
			gate:     true,
			expected: []string{},
		},
		{
			name:     "DisableComponentStatusProbes disable",
			gate:     false,
			expected: []string{"scheduler", "controller-manager", "etcd-0"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.DisableComponentStatusProbes, test.gate)()
			servers := componentStatusStorage{fakeStorageFactory{}}.serversToValidate()

			if len(servers) != len(test.expected) {
				t.Errorf("expected %v, got %v", len(servers), len(test.expected))
			}

			for _, server := range test.expected {
				if _, ok := servers[server]; !ok {
					t.Errorf("server list missing: %s", server)
				}
			}
		})
	}
}

type fakeStorageFactory struct{}

func (f fakeStorageFactory) NewConfig(groupResource schema.GroupResource) (*storagebackend.ConfigForResource, error) {
	return nil, nil
}

func (f fakeStorageFactory) ResourcePrefix(groupResource schema.GroupResource) string {
	return ""
}

func (f fakeStorageFactory) Backends() []storage.Backend {
	return []storage.Backend{{Server: "etcd-0"}}
}
