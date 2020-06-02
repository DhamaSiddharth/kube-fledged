/*
Copyright The kube-fledged authors.

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
	v1alpha1 "github.com/senthilrch/kube-fledged/pkg/apis/kubefledged/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ImageCacheLister helps list ImageCaches.
type ImageCacheLister interface {
	// List lists all ImageCaches in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ImageCache, err error)
	// ImageCaches returns an object that can list and get ImageCaches.
	ImageCaches(namespace string) ImageCacheNamespaceLister
	ImageCacheListerExpansion
}

// imageCacheLister implements the ImageCacheLister interface.
type imageCacheLister struct {
	indexer cache.Indexer
}

// NewImageCacheLister returns a new ImageCacheLister.
func NewImageCacheLister(indexer cache.Indexer) ImageCacheLister {
	return &imageCacheLister{indexer: indexer}
}

// List lists all ImageCaches in the indexer.
func (s *imageCacheLister) List(selector labels.Selector) (ret []*v1alpha1.ImageCache, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageCache))
	})
	return ret, err
}

// ImageCaches returns an object that can list and get ImageCaches.
func (s *imageCacheLister) ImageCaches(namespace string) ImageCacheNamespaceLister {
	return imageCacheNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ImageCacheNamespaceLister helps list and get ImageCaches.
type ImageCacheNamespaceLister interface {
	// List lists all ImageCaches in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ImageCache, err error)
	// Get retrieves the ImageCache from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ImageCache, error)
	ImageCacheNamespaceListerExpansion
}

// imageCacheNamespaceLister implements the ImageCacheNamespaceLister
// interface.
type imageCacheNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ImageCaches in the indexer for a given namespace.
func (s imageCacheNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ImageCache, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ImageCache))
	})
	return ret, err
}

// Get retrieves the ImageCache from the indexer for a given namespace and name.
func (s imageCacheNamespaceLister) Get(name string) (*v1alpha1.ImageCache, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("imagecache"), name)
	}
	return obj.(*v1alpha1.ImageCache), nil
}