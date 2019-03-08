// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/aws/aws-app-mesh-controller-for-k8s/pkg/apis/appmesh/v1alpha1"
	scheme "github.com/aws/aws-app-mesh-controller-for-k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MeshesGetter has a method to return a MeshInterface.
// A group's client should implement this interface.
type MeshesGetter interface {
	Meshes(namespace string) MeshInterface
}

// MeshInterface has methods to work with Mesh resources.
type MeshInterface interface {
	Create(*v1alpha1.Mesh) (*v1alpha1.Mesh, error)
	Update(*v1alpha1.Mesh) (*v1alpha1.Mesh, error)
	UpdateStatus(*v1alpha1.Mesh) (*v1alpha1.Mesh, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Mesh, error)
	List(opts v1.ListOptions) (*v1alpha1.MeshList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Mesh, err error)
	MeshExpansion
}

// meshes implements MeshInterface
type meshes struct {
	client rest.Interface
	ns     string
}

// newMeshes returns a Meshes
func newMeshes(c *AppmeshV1alpha1Client, namespace string) *meshes {
	return &meshes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mesh, and returns the corresponding mesh object, and an error if there is any.
func (c *meshes) Get(name string, options v1.GetOptions) (result *v1alpha1.Mesh, err error) {
	result = &v1alpha1.Mesh{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("meshes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Meshes that match those selectors.
func (c *meshes) List(opts v1.ListOptions) (result *v1alpha1.MeshList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MeshList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("meshes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested meshes.
func (c *meshes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("meshes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a mesh and creates it.  Returns the server's representation of the mesh, and an error, if there is any.
func (c *meshes) Create(mesh *v1alpha1.Mesh) (result *v1alpha1.Mesh, err error) {
	result = &v1alpha1.Mesh{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("meshes").
		Body(mesh).
		Do().
		Into(result)
	return
}

// Update takes the representation of a mesh and updates it. Returns the server's representation of the mesh, and an error, if there is any.
func (c *meshes) Update(mesh *v1alpha1.Mesh) (result *v1alpha1.Mesh, err error) {
	result = &v1alpha1.Mesh{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("meshes").
		Name(mesh.Name).
		Body(mesh).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *meshes) UpdateStatus(mesh *v1alpha1.Mesh) (result *v1alpha1.Mesh, err error) {
	result = &v1alpha1.Mesh{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("meshes").
		Name(mesh.Name).
		SubResource("status").
		Body(mesh).
		Do().
		Into(result)
	return
}

// Delete takes name of the mesh and deletes it. Returns an error if one occurs.
func (c *meshes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("meshes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *meshes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("meshes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched mesh.
func (c *meshes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Mesh, err error) {
	result = &v1alpha1.Mesh{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("meshes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
