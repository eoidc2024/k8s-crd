//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogConfig) DeepCopyInto(out *ClusterLogConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogConfig.
func (in *ClusterLogConfig) DeepCopy() *ClusterLogConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterLogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterLogConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterLogConfigList) DeepCopyInto(out *ClusterLogConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterLogConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterLogConfigList.
func (in *ClusterLogConfigList) DeepCopy() *ClusterLogConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterLogConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EoiPodSelector) DeepCopyInto(out *EoiPodSelector) {
	*out = *in
	if in.EoiPodSelector != nil {
		in, out := &in.EoiPodSelector, &out.EoiPodSelector
		*out = make([]LabelExpr, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EoiPodSelector.
func (in *EoiPodSelector) DeepCopy() *EoiPodSelector {
	if in == nil {
		return nil
	}
	out := new(EoiPodSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interceptor) DeepCopyInto(out *Interceptor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interceptor.
func (in *Interceptor) DeepCopy() *Interceptor {
	if in == nil {
		return nil
	}
	out := new(Interceptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorList) DeepCopyInto(out *InterceptorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Interceptor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorList.
func (in *InterceptorList) DeepCopy() *InterceptorList {
	if in == nil {
		return nil
	}
	out := new(InterceptorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterceptorSpec) DeepCopyInto(out *InterceptorSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterceptorSpec.
func (in *InterceptorSpec) DeepCopy() *InterceptorSpec {
	if in == nil {
		return nil
	}
	out := new(InterceptorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelExpr) DeepCopyInto(out *LabelExpr) {
	*out = *in
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelExpr.
func (in *LabelExpr) DeepCopy() *LabelExpr {
	if in == nil {
		return nil
	}
	out := new(LabelExpr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfig) DeepCopyInto(out *LogConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfig.
func (in *LogConfig) DeepCopy() *LogConfig {
	if in == nil {
		return nil
	}
	out := new(LogConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogConfigList) DeepCopyInto(out *LogConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LogConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogConfigList.
func (in *LogConfigList) DeepCopy() *LogConfigList {
	if in == nil {
		return nil
	}
	out := new(LogConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Message) DeepCopyInto(out *Message) {
	*out = *in
	return
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
func (in *NamespaceSelector) DeepCopyInto(out *NamespaceSelector) {
	*out = *in
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeNamespaceSelector != nil {
		in, out := &in.ExcludeNamespaceSelector, &out.ExcludeNamespaceSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceSelector.
func (in *NamespaceSelector) DeepCopy() *NamespaceSelector {
	if in == nil {
		return nil
	}
	out := new(NamespaceSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSelector) DeepCopyInto(out *NodeSelector) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSelector.
func (in *NodeSelector) DeepCopy() *NodeSelector {
	if in == nil {
		return nil
	}
	out := new(NodeSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodSelector) DeepCopyInto(out *PodSelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodSelector.
func (in *PodSelector) DeepCopy() *PodSelector {
	if in == nil {
		return nil
	}
	out := new(PodSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
	if in.WorkloadSelector != nil {
		in, out := &in.WorkloadSelector, &out.WorkloadSelector
		*out = make([]WorkloadSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PodSelector.DeepCopyInto(&out.PodSelector)
	in.NodeSelector.DeepCopyInto(&out.NodeSelector)
	in.NamespaceSelector.DeepCopyInto(&out.NamespaceSelector)
	in.EoiPodSelector.DeepCopyInto(&out.EoiPodSelector)
	return
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
func (in *Sink) DeepCopyInto(out *Sink) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sink.
func (in *Sink) DeepCopy() *Sink {
	if in == nil {
		return nil
	}
	out := new(Sink)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkList) DeepCopyInto(out *SinkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sink, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkList.
func (in *SinkList) DeepCopy() *SinkList {
	if in == nil {
		return nil
	}
	out := new(SinkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SinkSpec) DeepCopyInto(out *SinkSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SinkSpec.
func (in *SinkSpec) DeepCopy() *SinkSpec {
	if in == nil {
		return nil
	}
	out := new(SinkSpec)
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
	if in.Pipeline != nil {
		in, out := &in.Pipeline, &out.Pipeline
		*out = new(Pipeline)
		**out = **in
	}
	return
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
	return
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

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSelector) DeepCopyInto(out *WorkloadSelector) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NameSelector != nil {
		in, out := &in.NameSelector, &out.NameSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.NamespaceSelector != nil {
		in, out := &in.NamespaceSelector, &out.NamespaceSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeNamespaceSelector != nil {
		in, out := &in.ExcludeNamespaceSelector, &out.ExcludeNamespaceSelector
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSelector.
func (in *WorkloadSelector) DeepCopy() *WorkloadSelector {
	if in == nil {
		return nil
	}
	out := new(WorkloadSelector)
	in.DeepCopyInto(out)
	return out
}
