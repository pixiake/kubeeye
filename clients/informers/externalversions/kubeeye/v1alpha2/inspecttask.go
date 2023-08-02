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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	"context"
	time "time"

	kubeeyev1alpha2 "github.com/kubesphere/kubeeye-v1alpha2/apis/kubeeye/v1alpha2"
	versioned "github.com/kubesphere/kubeeye-v1alpha2/clients/clientset/versioned"
	internalinterfaces "github.com/kubesphere/kubeeye-v1alpha2/clients/informers/externalversions/internalinterfaces"
	v1alpha2 "github.com/kubesphere/kubeeye-v1alpha2/clients/listers/kubeeye/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InspectTaskInformer provides access to a shared informer and lister for
// InspectTasks.
type InspectTaskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.InspectTaskLister
}

type inspectTaskInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewInspectTaskInformer constructs a new informer for InspectTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInspectTaskInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInspectTaskInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredInspectTaskInformer constructs a new informer for InspectTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInspectTaskInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeeyeV1alpha2().InspectTasks().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeeyeV1alpha2().InspectTasks().Watch(context.TODO(), options)
			},
		},
		&kubeeyev1alpha2.InspectTask{},
		resyncPeriod,
		indexers,
	)
}

func (f *inspectTaskInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInspectTaskInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *inspectTaskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubeeyev1alpha2.InspectTask{}, f.defaultInformer)
}

func (f *inspectTaskInformer) Lister() v1alpha2.InspectTaskLister {
	return v1alpha2.NewInspectTaskLister(f.Informer().GetIndexer())
}
