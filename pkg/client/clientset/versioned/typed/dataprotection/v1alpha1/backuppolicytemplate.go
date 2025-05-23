/*
Copyright (C) 2022-2025 ApeCloud Co., Ltd

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
	"context"
	"time"

	v1alpha1 "github.com/apecloud/kubeblocks/apis/dataprotection/v1alpha1"
	scheme "github.com/apecloud/kubeblocks/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupPolicyTemplatesGetter has a method to return a BackupPolicyTemplateInterface.
// A group's client should implement this interface.
type BackupPolicyTemplatesGetter interface {
	BackupPolicyTemplates() BackupPolicyTemplateInterface
}

// BackupPolicyTemplateInterface has methods to work with BackupPolicyTemplate resources.
type BackupPolicyTemplateInterface interface {
	Create(ctx context.Context, backupPolicyTemplate *v1alpha1.BackupPolicyTemplate, opts v1.CreateOptions) (*v1alpha1.BackupPolicyTemplate, error)
	Update(ctx context.Context, backupPolicyTemplate *v1alpha1.BackupPolicyTemplate, opts v1.UpdateOptions) (*v1alpha1.BackupPolicyTemplate, error)
	UpdateStatus(ctx context.Context, backupPolicyTemplate *v1alpha1.BackupPolicyTemplate, opts v1.UpdateOptions) (*v1alpha1.BackupPolicyTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BackupPolicyTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BackupPolicyTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupPolicyTemplate, err error)
	BackupPolicyTemplateExpansion
}

// backupPolicyTemplates implements BackupPolicyTemplateInterface
type backupPolicyTemplates struct {
	client rest.Interface
}

// newBackupPolicyTemplates returns a BackupPolicyTemplates
func newBackupPolicyTemplates(c *DataprotectionV1alpha1Client) *backupPolicyTemplates {
	return &backupPolicyTemplates{
		client: c.RESTClient(),
	}
}

// Get takes name of the backupPolicyTemplate, and returns the corresponding backupPolicyTemplate object, and an error if there is any.
func (c *backupPolicyTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BackupPolicyTemplate, err error) {
	result = &v1alpha1.BackupPolicyTemplate{}
	err = c.client.Get().
		Resource("backuppolicytemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupPolicyTemplates that match those selectors.
func (c *backupPolicyTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackupPolicyTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BackupPolicyTemplateList{}
	err = c.client.Get().
		Resource("backuppolicytemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupPolicyTemplates.
func (c *backupPolicyTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("backuppolicytemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backupPolicyTemplate and creates it.  Returns the server's representation of the backupPolicyTemplate, and an error, if there is any.
func (c *backupPolicyTemplates) Create(ctx context.Context, backupPolicyTemplate *v1alpha1.BackupPolicyTemplate, opts v1.CreateOptions) (result *v1alpha1.BackupPolicyTemplate, err error) {
	result = &v1alpha1.BackupPolicyTemplate{}
	err = c.client.Post().
		Resource("backuppolicytemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupPolicyTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backupPolicyTemplate and updates it. Returns the server's representation of the backupPolicyTemplate, and an error, if there is any.
func (c *backupPolicyTemplates) Update(ctx context.Context, backupPolicyTemplate *v1alpha1.BackupPolicyTemplate, opts v1.UpdateOptions) (result *v1alpha1.BackupPolicyTemplate, err error) {
	result = &v1alpha1.BackupPolicyTemplate{}
	err = c.client.Put().
		Resource("backuppolicytemplates").
		Name(backupPolicyTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupPolicyTemplate).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *backupPolicyTemplates) UpdateStatus(ctx context.Context, backupPolicyTemplate *v1alpha1.BackupPolicyTemplate, opts v1.UpdateOptions) (result *v1alpha1.BackupPolicyTemplate, err error) {
	result = &v1alpha1.BackupPolicyTemplate{}
	err = c.client.Put().
		Resource("backuppolicytemplates").
		Name(backupPolicyTemplate.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backupPolicyTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backupPolicyTemplate and deletes it. Returns an error if one occurs.
func (c *backupPolicyTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("backuppolicytemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupPolicyTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("backuppolicytemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backupPolicyTemplate.
func (c *backupPolicyTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BackupPolicyTemplate, err error) {
	result = &v1alpha1.BackupPolicyTemplate{}
	err = c.client.Patch(pt).
		Resource("backuppolicytemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
