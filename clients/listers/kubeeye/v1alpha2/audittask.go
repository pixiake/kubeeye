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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuditTaskLister helps list AuditTasks.
// All objects returned here must be treated as read-only.
type AuditTaskLister interface {
	// List lists all AuditTasks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.AuditTask, err error)
	// AuditTasks returns an object that can list and get AuditTasks.
	AuditTasks(namespace string) AuditTaskNamespaceLister
	AuditTaskListerExpansion
}

// auditTaskLister implements the AuditTaskLister interface.
type auditTaskLister struct {
	indexer cache.Indexer
}

// NewAuditTaskLister returns a new AuditTaskLister.
func NewAuditTaskLister(indexer cache.Indexer) AuditTaskLister {
	return &auditTaskLister{indexer: indexer}
}

// List lists all AuditTasks in the indexer.
func (s *auditTaskLister) List(selector labels.Selector) (ret []*v1alpha2.AuditTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.AuditTask))
	})
	return ret, err
}

// AuditTasks returns an object that can list and get AuditTasks.
func (s *auditTaskLister) AuditTasks(namespace string) AuditTaskNamespaceLister {
	return auditTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AuditTaskNamespaceLister helps list and get AuditTasks.
// All objects returned here must be treated as read-only.
type AuditTaskNamespaceLister interface {
	// List lists all AuditTasks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha2.AuditTask, err error)
	// Get retrieves the AuditTask from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha2.AuditTask, error)
	AuditTaskNamespaceListerExpansion
}

// auditTaskNamespaceLister implements the AuditTaskNamespaceLister
// interface.
type auditTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AuditTasks in the indexer for a given namespace.
func (s auditTaskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.AuditTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.AuditTask))
	})
	return ret, err
}

// Get retrieves the AuditTask from the indexer for a given namespace and name.
func (s auditTaskNamespaceLister) Get(name string) (*v1alpha2.AuditTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("audittask"), name)
	}
	return obj.(*v1alpha2.AuditTask), nil
}
