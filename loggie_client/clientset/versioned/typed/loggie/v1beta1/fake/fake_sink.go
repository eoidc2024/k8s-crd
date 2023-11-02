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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "code.eoitek.net/monitor/dc/k8s-crd/apis/loggie/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSinks implements SinkInterface
type FakeSinks struct {
	Fake *FakeLoggieV1beta1
}

var sinksResource = schema.GroupVersionResource{Group: "loggie.io", Version: "v1beta1", Resource: "sinks"}

var sinksKind = schema.GroupVersionKind{Group: "loggie.io", Version: "v1beta1", Kind: "Sink"}

// Get takes name of the sink, and returns the corresponding sink object, and an error if there is any.
func (c *FakeSinks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Sink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(sinksResource, name), &v1beta1.Sink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Sink), err
}

// List takes label and field selectors, and returns the list of Sinks that match those selectors.
func (c *FakeSinks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.SinkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(sinksResource, sinksKind, opts), &v1beta1.SinkList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.SinkList{ListMeta: obj.(*v1beta1.SinkList).ListMeta}
	for _, item := range obj.(*v1beta1.SinkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sinks.
func (c *FakeSinks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(sinksResource, opts))
}

// Create takes the representation of a sink and creates it.  Returns the server's representation of the sink, and an error, if there is any.
func (c *FakeSinks) Create(ctx context.Context, sink *v1beta1.Sink, opts v1.CreateOptions) (result *v1beta1.Sink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(sinksResource, sink), &v1beta1.Sink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Sink), err
}

// Update takes the representation of a sink and updates it. Returns the server's representation of the sink, and an error, if there is any.
func (c *FakeSinks) Update(ctx context.Context, sink *v1beta1.Sink, opts v1.UpdateOptions) (result *v1beta1.Sink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(sinksResource, sink), &v1beta1.Sink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Sink), err
}

// Delete takes name of the sink and deletes it. Returns an error if one occurs.
func (c *FakeSinks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(sinksResource, name, opts), &v1beta1.Sink{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSinks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(sinksResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.SinkList{})
	return err
}

// Patch applies the patch and returns the patched sink.
func (c *FakeSinks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Sink, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(sinksResource, name, pt, data, subresources...), &v1beta1.Sink{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Sink), err
}
