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

package v1

import (
	v1 "dbprovision.com/pkg/apis/dbprovision/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DBProvisioningLister helps list DBProvisionings.
// All objects returned here must be treated as read-only.
type DBProvisioningLister interface {
	// List lists all DBProvisionings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DBProvisioning, err error)
	// DBProvisionings returns an object that can list and get DBProvisionings.
	DBProvisionings(namespace string) DBProvisioningNamespaceLister
	DBProvisioningListerExpansion
}

// dBProvisioningLister implements the DBProvisioningLister interface.
type dBProvisioningLister struct {
	indexer cache.Indexer
}

// NewDBProvisioningLister returns a new DBProvisioningLister.
func NewDBProvisioningLister(indexer cache.Indexer) DBProvisioningLister {
	return &dBProvisioningLister{indexer: indexer}
}

// List lists all DBProvisionings in the indexer.
func (s *dBProvisioningLister) List(selector labels.Selector) (ret []*v1.DBProvisioning, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DBProvisioning))
	})
	return ret, err
}

// DBProvisionings returns an object that can list and get DBProvisionings.
func (s *dBProvisioningLister) DBProvisionings(namespace string) DBProvisioningNamespaceLister {
	return dBProvisioningNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DBProvisioningNamespaceLister helps list and get DBProvisionings.
// All objects returned here must be treated as read-only.
type DBProvisioningNamespaceLister interface {
	// List lists all DBProvisionings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.DBProvisioning, err error)
	// Get retrieves the DBProvisioning from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.DBProvisioning, error)
	DBProvisioningNamespaceListerExpansion
}

// dBProvisioningNamespaceLister implements the DBProvisioningNamespaceLister
// interface.
type dBProvisioningNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DBProvisionings in the indexer for a given namespace.
func (s dBProvisioningNamespaceLister) List(selector labels.Selector) (ret []*v1.DBProvisioning, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DBProvisioning))
	})
	return ret, err
}

// Get retrieves the DBProvisioning from the indexer for a given namespace and name.
func (s dBProvisioningNamespaceLister) Get(name string) (*v1.DBProvisioning, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("dbprovisioning"), name)
	}
	return obj.(*v1.DBProvisioning), nil
}