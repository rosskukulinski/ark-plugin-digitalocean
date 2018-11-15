/*
Copyright 2018 the Heptio Ark contributors.

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
	v1 "github.com/heptio/ark/pkg/apis/ark/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BackupStorageLocationLister helps list BackupStorageLocations.
type BackupStorageLocationLister interface {
	// List lists all BackupStorageLocations in the indexer.
	List(selector labels.Selector) (ret []*v1.BackupStorageLocation, err error)
	// BackupStorageLocations returns an object that can list and get BackupStorageLocations.
	BackupStorageLocations(namespace string) BackupStorageLocationNamespaceLister
	BackupStorageLocationListerExpansion
}

// backupStorageLocationLister implements the BackupStorageLocationLister interface.
type backupStorageLocationLister struct {
	indexer cache.Indexer
}

// NewBackupStorageLocationLister returns a new BackupStorageLocationLister.
func NewBackupStorageLocationLister(indexer cache.Indexer) BackupStorageLocationLister {
	return &backupStorageLocationLister{indexer: indexer}
}

// List lists all BackupStorageLocations in the indexer.
func (s *backupStorageLocationLister) List(selector labels.Selector) (ret []*v1.BackupStorageLocation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BackupStorageLocation))
	})
	return ret, err
}

// BackupStorageLocations returns an object that can list and get BackupStorageLocations.
func (s *backupStorageLocationLister) BackupStorageLocations(namespace string) BackupStorageLocationNamespaceLister {
	return backupStorageLocationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BackupStorageLocationNamespaceLister helps list and get BackupStorageLocations.
type BackupStorageLocationNamespaceLister interface {
	// List lists all BackupStorageLocations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.BackupStorageLocation, err error)
	// Get retrieves the BackupStorageLocation from the indexer for a given namespace and name.
	Get(name string) (*v1.BackupStorageLocation, error)
	BackupStorageLocationNamespaceListerExpansion
}

// backupStorageLocationNamespaceLister implements the BackupStorageLocationNamespaceLister
// interface.
type backupStorageLocationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BackupStorageLocations in the indexer for a given namespace.
func (s backupStorageLocationNamespaceLister) List(selector labels.Selector) (ret []*v1.BackupStorageLocation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BackupStorageLocation))
	})
	return ret, err
}

// Get retrieves the BackupStorageLocation from the indexer for a given namespace and name.
func (s backupStorageLocationNamespaceLister) Get(name string) (*v1.BackupStorageLocation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("backupstoragelocation"), name)
	}
	return obj.(*v1.BackupStorageLocation), nil
}
