/*
Copyright 2022.

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

package v1alpha2

import (
	"context"
	"time"

	v1alpha2 "github.com/kubesphere/kubeeye-v1alpha2/apis/kubeeye/v1alpha2"
	scheme "github.com/kubesphere/kubeeye-v1alpha2/clients/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InspectResultsGetter has a method to return a InspectResultInterface.
// A group's client should implement this interface.
type InspectResultsGetter interface {
	InspectResults() InspectResultInterface
}

// InspectResultInterface has methods to work with InspectResult resources.
type InspectResultInterface interface {
	Create(ctx context.Context, inspectResult *v1alpha2.InspectResult, opts v1.CreateOptions) (*v1alpha2.InspectResult, error)
	Update(ctx context.Context, inspectResult *v1alpha2.InspectResult, opts v1.UpdateOptions) (*v1alpha2.InspectResult, error)
	UpdateStatus(ctx context.Context, inspectResult *v1alpha2.InspectResult, opts v1.UpdateOptions) (*v1alpha2.InspectResult, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.InspectResult, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.InspectResultList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.InspectResult, err error)
	InspectResultExpansion
}

// inspectResults implements InspectResultInterface
type inspectResults struct {
	client rest.Interface
}

// newInspectResults returns a InspectResults
func newInspectResults(c *KubeeyeV1alpha2Client) *inspectResults {
	return &inspectResults{
		client: c.RESTClient(),
	}
}

// Get takes name of the inspectResult, and returns the corresponding inspectResult object, and an error if there is any.
func (c *inspectResults) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.InspectResult, err error) {
	result = &v1alpha2.InspectResult{}
	err = c.client.Get().
		Resource("inspectresults").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InspectResults that match those selectors.
func (c *inspectResults) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.InspectResultList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.InspectResultList{}
	err = c.client.Get().
		Resource("inspectresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested inspectResults.
func (c *inspectResults) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("inspectresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a inspectResult and creates it.  Returns the server's representation of the inspectResult, and an error, if there is any.
func (c *inspectResults) Create(ctx context.Context, inspectResult *v1alpha2.InspectResult, opts v1.CreateOptions) (result *v1alpha2.InspectResult, err error) {
	result = &v1alpha2.InspectResult{}
	err = c.client.Post().
		Resource("inspectresults").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inspectResult).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a inspectResult and updates it. Returns the server's representation of the inspectResult, and an error, if there is any.
func (c *inspectResults) Update(ctx context.Context, inspectResult *v1alpha2.InspectResult, opts v1.UpdateOptions) (result *v1alpha2.InspectResult, err error) {
	result = &v1alpha2.InspectResult{}
	err = c.client.Put().
		Resource("inspectresults").
		Name(inspectResult.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inspectResult).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *inspectResults) UpdateStatus(ctx context.Context, inspectResult *v1alpha2.InspectResult, opts v1.UpdateOptions) (result *v1alpha2.InspectResult, err error) {
	result = &v1alpha2.InspectResult{}
	err = c.client.Put().
		Resource("inspectresults").
		Name(inspectResult.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(inspectResult).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the inspectResult and deletes it. Returns an error if one occurs.
func (c *inspectResults) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("inspectresults").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *inspectResults) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("inspectresults").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched inspectResult.
func (c *inspectResults) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.InspectResult, err error) {
	result = &v1alpha2.InspectResult{}
	err = c.client.Patch(pt).
		Resource("inspectresults").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
