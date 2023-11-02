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
// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	internalinterfaces "code.eoitek.net/monitor/dc/k8s-crd/loggie/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ClusterLogConfigs returns a ClusterLogConfigInformer.
	ClusterLogConfigs() ClusterLogConfigInformer
	// Interceptors returns a InterceptorInformer.
	Interceptors() InterceptorInformer
	// LogConfigs returns a LogConfigInformer.
	LogConfigs() LogConfigInformer
	// Sinks returns a SinkInformer.
	Sinks() SinkInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ClusterLogConfigs returns a ClusterLogConfigInformer.
func (v *version) ClusterLogConfigs() ClusterLogConfigInformer {
	return &clusterLogConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Interceptors returns a InterceptorInformer.
func (v *version) Interceptors() InterceptorInformer {
	return &interceptorInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// LogConfigs returns a LogConfigInformer.
func (v *version) LogConfigs() LogConfigInformer {
	return &logConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Sinks returns a SinkInformer.
func (v *version) Sinks() SinkInformer {
	return &sinkInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
