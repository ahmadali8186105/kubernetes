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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// ConcurrencyLimitStatusApplyConfiguration represents an declarative configuration of the ConcurrencyLimitStatus type for use
// with apply.
type ConcurrencyLimitStatusApplyConfiguration struct {
	APIServer *string `json:"apiServer,omitempty"`
	Limit     *int32  `json:"limit,omitempty"`
}

// ConcurrencyLimitStatusApplyConfiguration constructs an declarative configuration of the ConcurrencyLimitStatus type for use with
// apply.
func ConcurrencyLimitStatus() *ConcurrencyLimitStatusApplyConfiguration {
	return &ConcurrencyLimitStatusApplyConfiguration{}
}

// WithAPIServer sets the APIServer field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIServer field is set to the value of the last call.
func (b *ConcurrencyLimitStatusApplyConfiguration) WithAPIServer(value string) *ConcurrencyLimitStatusApplyConfiguration {
	b.APIServer = &value
	return b
}

// WithLimit sets the Limit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Limit field is set to the value of the last call.
func (b *ConcurrencyLimitStatusApplyConfiguration) WithLimit(value int32) *ConcurrencyLimitStatusApplyConfiguration {
	b.Limit = &value
	return b
}
