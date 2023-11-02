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
	"context"
	time "time"

	agenteoitekiov1beta1 "code.eoitek.net/monitor/dc/k8s-crd/apis/agent.eoitek.io/v1beta1"
	versioned "code.eoitek.net/monitor/dc/k8s-crd/client/clientset/versioned"
	internalinterfaces "code.eoitek.net/monitor/dc/k8s-crd/client/informers/externalversions/internalinterfaces"
	v1beta1 "code.eoitek.net/monitor/dc/k8s-crd/client/listers/agent.eoitek.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CrontabConfigInformer provides access to a shared informer and lister for
// CrontabConfigs.
type CrontabConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.CrontabConfigLister
}

type crontabConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCrontabConfigInformer constructs a new informer for CrontabConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCrontabConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCrontabConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCrontabConfigInformer constructs a new informer for CrontabConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCrontabConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AgentV1beta1().CrontabConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AgentV1beta1().CrontabConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&agenteoitekiov1beta1.CrontabConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *crontabConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCrontabConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *crontabConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&agenteoitekiov1beta1.CrontabConfig{}, f.defaultInformer)
}

func (f *crontabConfigInformer) Lister() v1beta1.CrontabConfigLister {
	return v1beta1.NewCrontabConfigLister(f.Informer().GetIndexer())
}