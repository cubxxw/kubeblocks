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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/apecloud/kubeblocks/apis/operations/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// OpsRequestLister helps list OpsRequests.
// All objects returned here must be treated as read-only.
type OpsRequestLister interface {
	// List lists all OpsRequests in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OpsRequest, err error)
	// OpsRequests returns an object that can list and get OpsRequests.
	OpsRequests(namespace string) OpsRequestNamespaceLister
	OpsRequestListerExpansion
}

// opsRequestLister implements the OpsRequestLister interface.
type opsRequestLister struct {
	indexer cache.Indexer
}

// NewOpsRequestLister returns a new OpsRequestLister.
func NewOpsRequestLister(indexer cache.Indexer) OpsRequestLister {
	return &opsRequestLister{indexer: indexer}
}

// List lists all OpsRequests in the indexer.
func (s *opsRequestLister) List(selector labels.Selector) (ret []*v1alpha1.OpsRequest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsRequest))
	})
	return ret, err
}

// OpsRequests returns an object that can list and get OpsRequests.
func (s *opsRequestLister) OpsRequests(namespace string) OpsRequestNamespaceLister {
	return opsRequestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// OpsRequestNamespaceLister helps list and get OpsRequests.
// All objects returned here must be treated as read-only.
type OpsRequestNamespaceLister interface {
	// List lists all OpsRequests in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.OpsRequest, err error)
	// Get retrieves the OpsRequest from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.OpsRequest, error)
	OpsRequestNamespaceListerExpansion
}

// opsRequestNamespaceLister implements the OpsRequestNamespaceLister
// interface.
type opsRequestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all OpsRequests in the indexer for a given namespace.
func (s opsRequestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.OpsRequest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.OpsRequest))
	})
	return ret, err
}

// Get retrieves the OpsRequest from the indexer for a given namespace and name.
func (s opsRequestNamespaceLister) Get(name string) (*v1alpha1.OpsRequest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("opsrequest"), name)
	}
	return obj.(*v1alpha1.OpsRequest), nil
}
