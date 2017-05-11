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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package discovery

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	reflect "reflect"
)

// GetGeneratedDeepCopyFuncs returns the generated funcs, since we aren't registering them.
func GetGeneratedDeepCopyFuncs() []conversion.GeneratedDeepCopyFunc {
	return []conversion.GeneratedDeepCopyFunc{
		{Fn: DeepCopy_discovery_Group, InType: reflect.TypeOf(&Group{})},
		{Fn: DeepCopy_discovery_GroupList, InType: reflect.TypeOf(&GroupList{})},
		{Fn: DeepCopy_discovery_GroupStatus, InType: reflect.TypeOf(&GroupStatus{})},
		{Fn: DeepCopy_discovery_GroupVersion, InType: reflect.TypeOf(&GroupVersion{})},
	}
}

func DeepCopy_discovery_Group(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Group)
		out := out.(*Group)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if newVal, err := c.DeepCopy(&in.Status); err != nil {
			return err
		} else {
			out.Status = *newVal.(*GroupStatus)
		}
		return nil
	}
}

func DeepCopy_discovery_GroupList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupList)
		out := out.(*GroupList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Group, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*Group)
				}
			}
		}
		return nil
	}
}

func DeepCopy_discovery_GroupStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupStatus)
		out := out.(*GroupStatus)
		*out = *in
		if in.Versions != nil {
			in, out := &in.Versions, &out.Versions
			*out = make([]GroupVersion, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*GroupVersion)
				}
			}
		}
		if in.ServerAddressByClientCIDRs != nil {
			in, out := &in.ServerAddressByClientCIDRs, &out.ServerAddressByClientCIDRs
			*out = make([]v1.ServerAddressByClientCIDR, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

func DeepCopy_discovery_GroupVersion(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*GroupVersion)
		out := out.(*GroupVersion)
		*out = *in
		if in.Resources != nil {
			in, out := &in.Resources, &out.Resources
			*out = make([]v1.APIResource, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*v1.APIResource)
				}
			}
		}
		return nil
	}
}
