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
// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/eoidc2024/k8s-crd/datadog/apis/agent.eoitek.io/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProfileConfigLister helps list ProfileConfigs.
// All objects returned here must be treated as read-only.
type ProfileConfigLister interface {
	// List lists all ProfileConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ProfileConfig, err error)
	// ProfileConfigs returns an object that can list and get ProfileConfigs.
	ProfileConfigs(namespace string) ProfileConfigNamespaceLister
	ProfileConfigListerExpansion
}

// profileConfigLister implements the ProfileConfigLister interface.
type profileConfigLister struct {
	indexer cache.Indexer
}

// NewProfileConfigLister returns a new ProfileConfigLister.
func NewProfileConfigLister(indexer cache.Indexer) ProfileConfigLister {
	return &profileConfigLister{indexer: indexer}
}

// List lists all ProfileConfigs in the indexer.
func (s *profileConfigLister) List(selector labels.Selector) (ret []*v1beta1.ProfileConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ProfileConfig))
	})
	return ret, err
}

// ProfileConfigs returns an object that can list and get ProfileConfigs.
func (s *profileConfigLister) ProfileConfigs(namespace string) ProfileConfigNamespaceLister {
	return profileConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProfileConfigNamespaceLister helps list and get ProfileConfigs.
// All objects returned here must be treated as read-only.
type ProfileConfigNamespaceLister interface {
	// List lists all ProfileConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.ProfileConfig, err error)
	// Get retrieves the ProfileConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.ProfileConfig, error)
	ProfileConfigNamespaceListerExpansion
}

// profileConfigNamespaceLister implements the ProfileConfigNamespaceLister
// interface.
type profileConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ProfileConfigs in the indexer for a given namespace.
func (s profileConfigNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.ProfileConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.ProfileConfig))
	})
	return ret, err
}

// Get retrieves the ProfileConfig from the indexer for a given namespace and name.
func (s profileConfigNamespaceLister) Get(name string) (*v1beta1.ProfileConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("profileconfig"), name)
	}
	return obj.(*v1beta1.ProfileConfig), nil
}
