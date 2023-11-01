//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrontabConfig) DeepCopyInto(out *CrontabConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrontabConfig.
func (in *CrontabConfig) DeepCopy() *CrontabConfig {
	if in == nil {
		return nil
	}
	out := new(CrontabConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CrontabConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrontabConfigList) DeepCopyInto(out *CrontabConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CrontabConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrontabConfigList.
func (in *CrontabConfigList) DeepCopy() *CrontabConfigList {
	if in == nil {
		return nil
	}
	out := new(CrontabConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CrontabConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrontabSpec) DeepCopyInto(out *CrontabSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Crontab != nil {
		in, out := &in.Crontab, &out.Crontab
		*out = new(CrontabTaskConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrontabSpec.
func (in *CrontabSpec) DeepCopy() *CrontabSpec {
	if in == nil {
		return nil
	}
	out := new(CrontabSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrontabTaskConfig) DeepCopyInto(out *CrontabTaskConfig) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Fields != nil {
		in, out := &in.Fields, &out.Fields
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrontabTaskConfig.
func (in *CrontabTaskConfig) DeepCopy() *CrontabTaskConfig {
	if in == nil {
		return nil
	}
	out := new(CrontabTaskConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Golang) DeepCopyInto(out *Golang) {
	*out = *in
	if in.EnabledTypes != nil {
		in, out := &in.EnabledTypes, &out.EnabledTypes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = make(map[string]int32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Golang.
func (in *Golang) DeepCopy() *Golang {
	if in == nil {
		return nil
	}
	out := new(Golang)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Java) DeepCopyInto(out *Java) {
	*out = *in
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Java.
func (in *Java) DeepCopy() *Java {
	if in == nil {
		return nil
	}
	out := new(Java)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Message) DeepCopyInto(out *Message) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Message.
func (in *Message) DeepCopy() *Message {
	if in == nil {
		return nil
	}
	out := new(Message)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Native) DeepCopyInto(out *Native) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(string)
		**out = **in
	}
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(uint32)
		**out = **in
	}
	if in.OnCpu != nil {
		in, out := &in.OnCpu, &out.OnCpu
		*out = new(OnCPUConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OffCpu != nil {
		in, out := &in.OffCpu, &out.OffCpu
		*out = new(OffCPUConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Native.
func (in *Native) DeepCopy() *Native {
	if in == nil {
		return nil
	}
	out := new(Native)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OffCPUConfig) DeepCopyInto(out *OffCPUConfig) {
	*out = *in
	if in.MaxBlock != nil {
		in, out := &in.MaxBlock, &out.MaxBlock
		*out = new(uint64)
		**out = **in
	}
	if in.MinBlock != nil {
		in, out := &in.MinBlock, &out.MinBlock
		*out = new(uint64)
		**out = **in
	}
	if in.KernelOnly != nil {
		in, out := &in.KernelOnly, &out.KernelOnly
		*out = new(bool)
		**out = **in
	}
	if in.UserOnly != nil {
		in, out := &in.UserOnly, &out.UserOnly
		*out = new(bool)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OffCPUConfig.
func (in *OffCPUConfig) DeepCopy() *OffCPUConfig {
	if in == nil {
		return nil
	}
	out := new(OffCPUConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OnCPUConfig) DeepCopyInto(out *OnCPUConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Frequency != nil {
		in, out := &in.Frequency, &out.Frequency
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OnCPUConfig.
func (in *OnCPUConfig) DeepCopy() *OnCPUConfig {
	if in == nil {
		return nil
	}
	out := new(OnCPUConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileConfig) DeepCopyInto(out *ProfileConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileConfig.
func (in *ProfileConfig) DeepCopy() *ProfileConfig {
	if in == nil {
		return nil
	}
	out := new(ProfileConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileConfigList) DeepCopyInto(out *ProfileConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ProfileConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileConfigList.
func (in *ProfileConfigList) DeepCopy() *ProfileConfigList {
	if in == nil {
		return nil
	}
	out := new(ProfileConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProfileConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SkywalkingProfileConfig) DeepCopyInto(out *SkywalkingProfileConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SkywalkingProfileConfig.
func (in *SkywalkingProfileConfig) DeepCopy() *SkywalkingProfileConfig {
	if in == nil {
		return nil
	}
	out := new(SkywalkingProfileConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Spec) DeepCopyInto(out *Spec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Java != nil {
		in, out := &in.Java, &out.Java
		*out = new(Java)
		(*in).DeepCopyInto(*out)
	}
	if in.Golang != nil {
		in, out := &in.Golang, &out.Golang
		*out = new(Golang)
		(*in).DeepCopyInto(*out)
	}
	if in.Native != nil {
		in, out := &in.Native, &out.Native
		*out = new(Native)
		(*in).DeepCopyInto(*out)
	}
	if in.SwProfileConfig != nil {
		in, out := &in.SwProfileConfig, &out.SwProfileConfig
		*out = new(SkywalkingProfileConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Spec.
func (in *Spec) DeepCopy() *Spec {
	if in == nil {
		return nil
	}
	out := new(Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Status) DeepCopyInto(out *Status) {
	*out = *in
	out.Message = in.Message
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Status.
func (in *Status) DeepCopy() *Status {
	if in == nil {
		return nil
	}
	out := new(Status)
	in.DeepCopyInto(out)
	return out
}
