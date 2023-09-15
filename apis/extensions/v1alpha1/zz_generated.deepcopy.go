//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (C) 2022-2023 ApeCloud Co., Ltd

This file is part of KubeBlocks project

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Addon) DeepCopyInto(out *Addon) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Addon.
func (in *Addon) DeepCopy() *Addon {
	if in == nil {
		return nil
	}
	out := new(Addon)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Addon) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonDefaultInstallSpecItem) DeepCopyInto(out *AddonDefaultInstallSpecItem) {
	*out = *in
	in.AddonInstallSpec.DeepCopyInto(&out.AddonInstallSpec)
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]SelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonDefaultInstallSpecItem.
func (in *AddonDefaultInstallSpecItem) DeepCopy() *AddonDefaultInstallSpecItem {
	if in == nil {
		return nil
	}
	out := new(AddonDefaultInstallSpecItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallExtraItem) DeepCopyInto(out *AddonInstallExtraItem) {
	*out = *in
	in.AddonInstallSpecItem.DeepCopyInto(&out.AddonInstallSpecItem)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallExtraItem.
func (in *AddonInstallExtraItem) DeepCopy() *AddonInstallExtraItem {
	if in == nil {
		return nil
	}
	out := new(AddonInstallExtraItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallSpec) DeepCopyInto(out *AddonInstallSpec) {
	*out = *in
	in.AddonInstallSpecItem.DeepCopyInto(&out.AddonInstallSpecItem)
	if in.ExtraItems != nil {
		in, out := &in.ExtraItems, &out.ExtraItems
		*out = make([]AddonInstallExtraItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallSpec.
func (in *AddonInstallSpec) DeepCopy() *AddonInstallSpec {
	if in == nil {
		return nil
	}
	out := new(AddonInstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonInstallSpecItem) DeepCopyInto(out *AddonInstallSpecItem) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.PVEnabled != nil {
		in, out := &in.PVEnabled, &out.PVEnabled
		*out = new(bool)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonInstallSpecItem.
func (in *AddonInstallSpecItem) DeepCopy() *AddonInstallSpecItem {
	if in == nil {
		return nil
	}
	out := new(AddonInstallSpecItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonList) DeepCopyInto(out *AddonList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Addon, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonList.
func (in *AddonList) DeepCopy() *AddonList {
	if in == nil {
		return nil
	}
	out := new(AddonList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AddonList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonSpec) DeepCopyInto(out *AddonSpec) {
	*out = *in
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(HelmTypeInstallSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DefaultInstallValues != nil {
		in, out := &in.DefaultInstallValues, &out.DefaultInstallValues
		*out = make([]AddonDefaultInstallSpecItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InstallSpec != nil {
		in, out := &in.InstallSpec, &out.InstallSpec
		*out = new(AddonInstallSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Installable != nil {
		in, out := &in.Installable, &out.Installable
		*out = new(InstallableSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CliPlugins != nil {
		in, out := &in.CliPlugins, &out.CliPlugins
		*out = make([]CliPlugin, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonSpec.
func (in *AddonSpec) DeepCopy() *AddonSpec {
	if in == nil {
		return nil
	}
	out := new(AddonSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AddonStatus) DeepCopyInto(out *AddonStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AddonStatus.
func (in *AddonStatus) DeepCopy() *AddonStatus {
	if in == nil {
		return nil
	}
	out := new(AddonStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CliPlugin) DeepCopyInto(out *CliPlugin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CliPlugin.
func (in *CliPlugin) DeepCopy() *CliPlugin {
	if in == nil {
		return nil
	}
	out := new(CliPlugin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataObjectKeySelector) DeepCopyInto(out *DataObjectKeySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataObjectKeySelector.
func (in *DataObjectKeySelector) DeepCopy() *DataObjectKeySelector {
	if in == nil {
		return nil
	}
	out := new(DataObjectKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in HelmInstallOptions) DeepCopyInto(out *HelmInstallOptions) {
	{
		in := &in
		*out = make(HelmInstallOptions, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmInstallOptions.
func (in HelmInstallOptions) DeepCopy() HelmInstallOptions {
	if in == nil {
		return nil
	}
	out := new(HelmInstallOptions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmInstallValues) DeepCopyInto(out *HelmInstallValues) {
	*out = *in
	if in.URLs != nil {
		in, out := &in.URLs, &out.URLs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigMapRefs != nil {
		in, out := &in.ConfigMapRefs, &out.ConfigMapRefs
		*out = make([]DataObjectKeySelector, len(*in))
		copy(*out, *in)
	}
	if in.SecretRefs != nil {
		in, out := &in.SecretRefs, &out.SecretRefs
		*out = make([]DataObjectKeySelector, len(*in))
		copy(*out, *in)
	}
	if in.SetValues != nil {
		in, out := &in.SetValues, &out.SetValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SetJSONValues != nil {
		in, out := &in.SetJSONValues, &out.SetJSONValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmInstallValues.
func (in *HelmInstallValues) DeepCopy() *HelmInstallValues {
	if in == nil {
		return nil
	}
	out := new(HelmInstallValues)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmJSONValueMapType) DeepCopyInto(out *HelmJSONValueMapType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmJSONValueMapType.
func (in *HelmJSONValueMapType) DeepCopy() *HelmJSONValueMapType {
	if in == nil {
		return nil
	}
	out := new(HelmJSONValueMapType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmTypeInstallSpec) DeepCopyInto(out *HelmTypeInstallSpec) {
	*out = *in
	if in.InstallOptions != nil {
		in, out := &in.InstallOptions, &out.InstallOptions
		*out = make(HelmInstallOptions, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.InstallValues.DeepCopyInto(&out.InstallValues)
	in.ValuesMapping.DeepCopyInto(&out.ValuesMapping)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmTypeInstallSpec.
func (in *HelmTypeInstallSpec) DeepCopy() *HelmTypeInstallSpec {
	if in == nil {
		return nil
	}
	out := new(HelmTypeInstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmValueMapType) DeepCopyInto(out *HelmValueMapType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmValueMapType.
func (in *HelmValueMapType) DeepCopy() *HelmValueMapType {
	if in == nil {
		return nil
	}
	out := new(HelmValueMapType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmValuesMapping) DeepCopyInto(out *HelmValuesMapping) {
	*out = *in
	in.HelmValuesMappingItem.DeepCopyInto(&out.HelmValuesMappingItem)
	if in.ExtraItems != nil {
		in, out := &in.ExtraItems, &out.ExtraItems
		*out = make([]HelmValuesMappingExtraItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmValuesMapping.
func (in *HelmValuesMapping) DeepCopy() *HelmValuesMapping {
	if in == nil {
		return nil
	}
	out := new(HelmValuesMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmValuesMappingExtraItem) DeepCopyInto(out *HelmValuesMappingExtraItem) {
	*out = *in
	in.HelmValuesMappingItem.DeepCopyInto(&out.HelmValuesMappingItem)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmValuesMappingExtraItem.
func (in *HelmValuesMappingExtraItem) DeepCopy() *HelmValuesMappingExtraItem {
	if in == nil {
		return nil
	}
	out := new(HelmValuesMappingExtraItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmValuesMappingItem) DeepCopyInto(out *HelmValuesMappingItem) {
	*out = *in
	out.HelmValueMap = in.HelmValueMap
	out.HelmJSONMap = in.HelmJSONMap
	if in.ResourcesMapping != nil {
		in, out := &in.ResourcesMapping, &out.ResourcesMapping
		*out = new(ResourceMappingItem)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmValuesMappingItem.
func (in *HelmValuesMappingItem) DeepCopy() *HelmValuesMappingItem {
	if in == nil {
		return nil
	}
	out := new(HelmValuesMappingItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallableSpec) DeepCopyInto(out *InstallableSpec) {
	*out = *in
	if in.Selectors != nil {
		in, out := &in.Selectors, &out.Selectors
		*out = make([]SelectorRequirement, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallableSpec.
func (in *InstallableSpec) DeepCopy() *InstallableSpec {
	if in == nil {
		return nil
	}
	out := new(InstallableSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceMappingItem) DeepCopyInto(out *ResourceMappingItem) {
	*out = *in
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(ResourceReqLimItem)
		**out = **in
	}
	if in.Memory != nil {
		in, out := &in.Memory, &out.Memory
		*out = new(ResourceReqLimItem)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceMappingItem.
func (in *ResourceMappingItem) DeepCopy() *ResourceMappingItem {
	if in == nil {
		return nil
	}
	out := new(ResourceMappingItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceReqLimItem) DeepCopyInto(out *ResourceReqLimItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceReqLimItem.
func (in *ResourceReqLimItem) DeepCopy() *ResourceReqLimItem {
	if in == nil {
		return nil
	}
	out := new(ResourceReqLimItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequirements) DeepCopyInto(out *ResourceRequirements) {
	*out = *in
	if in.Limits != nil {
		in, out := &in.Limits, &out.Limits
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make(corev1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequirements.
func (in *ResourceRequirements) DeepCopy() *ResourceRequirements {
	if in == nil {
		return nil
	}
	out := new(ResourceRequirements)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelectorRequirement) DeepCopyInto(out *SelectorRequirement) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelectorRequirement.
func (in *SelectorRequirement) DeepCopy() *SelectorRequirement {
	if in == nil {
		return nil
	}
	out := new(SelectorRequirement)
	in.DeepCopyInto(out)
	return out
}
