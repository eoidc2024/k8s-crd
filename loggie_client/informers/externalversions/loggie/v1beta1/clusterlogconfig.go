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

	loggiev1beta1 "code.eoitek.net/monitor/dc/k8s-crd/apis/loggie/v1beta1"
	versioned "code.eoitek.net/monitor/dc/k8s-crd/loggie_client/clientset/versioned"
	internalinterfaces "code.eoitek.net/monitor/dc/k8s-crd/loggie_client/informers/externalversions/internalinterfaces"
	v1beta1 "code.eoitek.net/monitor/dc/k8s-crd/loggie_client/listers/loggie/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterLogConfigInformer provides access to a shared informer and lister for
// ClusterLogConfigs.
type ClusterLogConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ClusterLogConfigLister
}

type clusterLogConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterLogConfigInformer constructs a new informer for ClusterLogConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterLogConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterLogConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterLogConfigInformer constructs a new informer for ClusterLogConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterLogConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoggieV1beta1().ClusterLogConfigs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoggieV1beta1().ClusterLogConfigs().Watch(context.TODO(), options)
			},
		},
		&loggiev1beta1.ClusterLogConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterLogConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterLogConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterLogConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&loggiev1beta1.ClusterLogConfig{}, f.defaultInformer)
}

func (f *clusterLogConfigInformer) Lister() v1beta1.ClusterLogConfigLister {
	return v1beta1.NewClusterLogConfigLister(f.Informer().GetIndexer())
}
