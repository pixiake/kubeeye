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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	context "context"

	kubeeyev1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	applyconfigurationkubeeyev1alpha2 "github.com/kubesphere/kubeeye/clients/applyconfiguration/kubeeye/v1alpha2"
	scheme "github.com/kubesphere/kubeeye/clients/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// InspectRulesGetter has a method to return a InspectRuleInterface.
// A group's client should implement this interface.
type InspectRulesGetter interface {
	InspectRules() InspectRuleInterface
}

// InspectRuleInterface has methods to work with InspectRule resources.
type InspectRuleInterface interface {
	Create(ctx context.Context, inspectRule *kubeeyev1alpha2.InspectRule, opts v1.CreateOptions) (*kubeeyev1alpha2.InspectRule, error)
	Update(ctx context.Context, inspectRule *kubeeyev1alpha2.InspectRule, opts v1.UpdateOptions) (*kubeeyev1alpha2.InspectRule, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, inspectRule *kubeeyev1alpha2.InspectRule, opts v1.UpdateOptions) (*kubeeyev1alpha2.InspectRule, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*kubeeyev1alpha2.InspectRule, error)
	List(ctx context.Context, opts v1.ListOptions) (*kubeeyev1alpha2.InspectRuleList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubeeyev1alpha2.InspectRule, err error)
	Apply(ctx context.Context, inspectRule *applyconfigurationkubeeyev1alpha2.InspectRuleApplyConfiguration, opts v1.ApplyOptions) (result *kubeeyev1alpha2.InspectRule, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, inspectRule *applyconfigurationkubeeyev1alpha2.InspectRuleApplyConfiguration, opts v1.ApplyOptions) (result *kubeeyev1alpha2.InspectRule, err error)
	InspectRuleExpansion
}

// inspectRules implements InspectRuleInterface
type inspectRules struct {
	*gentype.ClientWithListAndApply[*kubeeyev1alpha2.InspectRule, *kubeeyev1alpha2.InspectRuleList, *applyconfigurationkubeeyev1alpha2.InspectRuleApplyConfiguration]
}

// newInspectRules returns a InspectRules
func newInspectRules(c *KubeeyeV1alpha2Client) *inspectRules {
	return &inspectRules{
		gentype.NewClientWithListAndApply[*kubeeyev1alpha2.InspectRule, *kubeeyev1alpha2.InspectRuleList, *applyconfigurationkubeeyev1alpha2.InspectRuleApplyConfiguration](
			"inspectrules",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *kubeeyev1alpha2.InspectRule { return &kubeeyev1alpha2.InspectRule{} },
			func() *kubeeyev1alpha2.InspectRuleList { return &kubeeyev1alpha2.InspectRuleList{} },
		),
	}
}
