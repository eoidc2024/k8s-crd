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

	loggiev1beta1 "github.com/eoidc2024/k8s-crd/loggie/apis/loggie/v1beta1"
	versioned "github.com/eoidc2024/k8s-crd/loggie/client/clientset/versioned"
	internalinterfaces "github.com/eoidc2024/k8s-crd/loggie/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/eoidc2024/k8s-crd/loggie/client/listers/loggie/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SinkInformer provides access to a shared informer and lister for
// Sinks.
type SinkInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.SinkLister
}

type sinkInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewSinkInformer constructs a new informer for Sink type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSinkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSinkInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredSinkInformer constructs a new informer for Sink type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSinkInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoggieV1beta1().Sinks().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoggieV1beta1().Sinks().Watch(context.TODO(), options)
			},
		},
		&loggiev1beta1.Sink{},
		resyncPeriod,
		indexers,
	)
}

func (f *sinkInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSinkInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sinkInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&loggiev1beta1.Sink{}, f.defaultInformer)
}

func (f *sinkInformer) Lister() v1beta1.SinkLister {
	return v1beta1.NewSinkLister(f.Informer().GetIndexer())
}
