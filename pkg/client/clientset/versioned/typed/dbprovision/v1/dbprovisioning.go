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

package v1

import (
	"context"
	"time"

	v1 "github.com/nuthankumar/cosmosdb/pkg/apis/dbprovision/v1"
	scheme "github.com/nuthankumar/cosmosdb/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DBProvisioningsGetter has a method to return a DBProvisioningInterface.
// A group's client should implement this interface.
type DBProvisioningsGetter interface {
	DBProvisionings(namespace string) DBProvisioningInterface
}

// DBProvisioningInterface has methods to work with DBProvisioning resources.
type DBProvisioningInterface interface {
	Create(ctx context.Context, dBProvisioning *v1.DBProvisioning, opts metav1.CreateOptions) (*v1.DBProvisioning, error)
	Update(ctx context.Context, dBProvisioning *v1.DBProvisioning, opts metav1.UpdateOptions) (*v1.DBProvisioning, error)
	UpdateStatus(ctx context.Context, dBProvisioning *v1.DBProvisioning, opts metav1.UpdateOptions) (*v1.DBProvisioning, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.DBProvisioning, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.DBProvisioningList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DBProvisioning, err error)
	DBProvisioningExpansion
}

// dBProvisionings implements DBProvisioningInterface
type dBProvisionings struct {
	client rest.Interface
	ns     string
}

// newDBProvisionings returns a DBProvisionings
func newDBProvisionings(c *DbprovisionV1Client, namespace string) *dBProvisionings {
	return &dBProvisionings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dBProvisioning, and returns the corresponding dBProvisioning object, and an error if there is any.
func (c *dBProvisionings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DBProvisioning, err error) {
	result = &v1.DBProvisioning{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbprovisionings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DBProvisionings that match those selectors.
func (c *dBProvisionings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DBProvisioningList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.DBProvisioningList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("dbprovisionings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dBProvisionings.
func (c *dBProvisionings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("dbprovisionings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dBProvisioning and creates it.  Returns the server's representation of the dBProvisioning, and an error, if there is any.
func (c *dBProvisionings) Create(ctx context.Context, dBProvisioning *v1.DBProvisioning, opts metav1.CreateOptions) (result *v1.DBProvisioning, err error) {
	result = &v1.DBProvisioning{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("dbprovisionings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dBProvisioning).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dBProvisioning and updates it. Returns the server's representation of the dBProvisioning, and an error, if there is any.
func (c *dBProvisionings) Update(ctx context.Context, dBProvisioning *v1.DBProvisioning, opts metav1.UpdateOptions) (result *v1.DBProvisioning, err error) {
	result = &v1.DBProvisioning{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbprovisionings").
		Name(dBProvisioning.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dBProvisioning).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dBProvisionings) UpdateStatus(ctx context.Context, dBProvisioning *v1.DBProvisioning, opts metav1.UpdateOptions) (result *v1.DBProvisioning, err error) {
	result = &v1.DBProvisioning{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("dbprovisionings").
		Name(dBProvisioning.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dBProvisioning).
		Do().
		Into(result)
	return
}

// Delete takes name of the dBProvisioning and deletes it. Returns an error if one occurs.
func (c *dBProvisionings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbprovisionings").
		Name(name).
		Body(&opts).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dBProvisionings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("dbprovisionings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dBProvisioning.
func (c *dBProvisionings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DBProvisioning, err error) {
	result = &v1.DBProvisioning{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("dbprovisionings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do().
		Into(result)
	return
}
