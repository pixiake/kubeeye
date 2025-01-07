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
	context "context"
	time "time"

	apiskubeeyev1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	versioned "github.com/kubesphere/kubeeye/clients/clientset/versioned"
	internalinterfaces "github.com/kubesphere/kubeeye/clients/informers/externalversions/internalinterfaces"
	kubeeyev1alpha2 "github.com/kubesphere/kubeeye/clients/listers/kubeeye/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// InspectPlanInformer provides access to a shared informer and lister for
// InspectPlans.
type InspectPlanInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() kubeeyev1alpha2.InspectPlanLister
}

type inspectPlanInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewInspectPlanInformer constructs a new informer for InspectPlan type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewInspectPlanInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredInspectPlanInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredInspectPlanInformer constructs a new informer for InspectPlan type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredInspectPlanInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeeyeV1alpha2().InspectPlans().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubeeyeV1alpha2().InspectPlans().Watch(context.TODO(), options)
			},
		},
		&apiskubeeyev1alpha2.InspectPlan{},
		resyncPeriod,
		indexers,
	)
}

func (f *inspectPlanInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredInspectPlanInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *inspectPlanInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiskubeeyev1alpha2.InspectPlan{}, f.defaultInformer)
}

func (f *inspectPlanInformer) Lister() kubeeyev1alpha2.InspectPlanLister {
	return kubeeyev1alpha2.NewInspectPlanLister(f.Informer().GetIndexer())
}
