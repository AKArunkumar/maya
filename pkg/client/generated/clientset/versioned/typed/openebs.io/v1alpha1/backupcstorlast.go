/*
Copyright 2019 The OpenEBS Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	scheme "github.com/openebs/maya/pkg/client/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupCStorLastsGetter has a method to return a BackupCStorLastInterface.
// A group's client should implement this interface.
type BackupCStorLastsGetter interface {
	BackupCStorLasts(namespace string) BackupCStorLastInterface
}

// BackupCStorLastInterface has methods to work with BackupCStorLast resources.
type BackupCStorLastInterface interface {
	Create(*v1alpha1.BackupCStorLast) (*v1alpha1.BackupCStorLast, error)
	Update(*v1alpha1.BackupCStorLast) (*v1alpha1.BackupCStorLast, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.BackupCStorLast, error)
	List(opts v1.ListOptions) (*v1alpha1.BackupCStorLastList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupCStorLast, err error)
	BackupCStorLastExpansion
}

// backupCStorLasts implements BackupCStorLastInterface
type backupCStorLasts struct {
	client rest.Interface
	ns     string
}

// newBackupCStorLasts returns a BackupCStorLasts
func newBackupCStorLasts(c *OpenebsV1alpha1Client, namespace string) *backupCStorLasts {
	return &backupCStorLasts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupCStorLast, and returns the corresponding backupCStorLast object, and an error if there is any.
func (c *backupCStorLasts) Get(name string, options v1.GetOptions) (result *v1alpha1.BackupCStorLast, err error) {
	result = &v1alpha1.BackupCStorLast{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupcstorlasts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupCStorLasts that match those selectors.
func (c *backupCStorLasts) List(opts v1.ListOptions) (result *v1alpha1.BackupCStorLastList, err error) {
	result = &v1alpha1.BackupCStorLastList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupcstorlasts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupCStorLasts.
func (c *backupCStorLasts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupcstorlasts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a backupCStorLast and creates it.  Returns the server's representation of the backupCStorLast, and an error, if there is any.
func (c *backupCStorLasts) Create(backupCStorLast *v1alpha1.BackupCStorLast) (result *v1alpha1.BackupCStorLast, err error) {
	result = &v1alpha1.BackupCStorLast{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupcstorlasts").
		Body(backupCStorLast).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupCStorLast and updates it. Returns the server's representation of the backupCStorLast, and an error, if there is any.
func (c *backupCStorLasts) Update(backupCStorLast *v1alpha1.BackupCStorLast) (result *v1alpha1.BackupCStorLast, err error) {
	result = &v1alpha1.BackupCStorLast{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupcstorlasts").
		Name(backupCStorLast.Name).
		Body(backupCStorLast).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupCStorLast and deletes it. Returns an error if one occurs.
func (c *backupCStorLasts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupcstorlasts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupCStorLasts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupcstorlasts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupCStorLast.
func (c *backupCStorLasts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.BackupCStorLast, err error) {
	result = &v1alpha1.BackupCStorLast{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupcstorlasts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
