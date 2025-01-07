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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha2

import (
	kubeeyev1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
)

// SysRuleApplyConfiguration represents a declarative configuration of the SysRule type for use
// with apply.
type SysRuleApplyConfiguration struct {
	RuleItemBasesApplyConfiguration `json:",inline"`
	NodeApplyConfiguration          `json:",inline"`
}

// SysRuleApplyConfiguration constructs a declarative configuration of the SysRule type for use with
// apply.
func SysRule() *SysRuleApplyConfiguration {
	return &SysRuleApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SysRuleApplyConfiguration) WithName(value string) *SysRuleApplyConfiguration {
	b.RuleItemBasesApplyConfiguration.Name = &value
	return b
}

// WithRule sets the Rule field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Rule field is set to the value of the last call.
func (b *SysRuleApplyConfiguration) WithRule(value string) *SysRuleApplyConfiguration {
	b.RuleItemBasesApplyConfiguration.Rule = &value
	return b
}

// WithDesc sets the Desc field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Desc field is set to the value of the last call.
func (b *SysRuleApplyConfiguration) WithDesc(value string) *SysRuleApplyConfiguration {
	b.RuleItemBasesApplyConfiguration.Desc = &value
	return b
}

// WithLevel sets the Level field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Level field is set to the value of the last call.
func (b *SysRuleApplyConfiguration) WithLevel(value kubeeyev1alpha2.Level) *SysRuleApplyConfiguration {
	b.RuleItemBasesApplyConfiguration.Level = &value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *SysRuleApplyConfiguration) WithNodeName(value string) *SysRuleApplyConfiguration {
	b.NodeApplyConfiguration.NodeName = &value
	return b
}

// WithNodeSelector puts the entries into the NodeSelector field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the NodeSelector field,
// overwriting an existing map entries in NodeSelector field with the same key.
func (b *SysRuleApplyConfiguration) WithNodeSelector(entries map[string]string) *SysRuleApplyConfiguration {
	if b.NodeApplyConfiguration.NodeSelector == nil && len(entries) > 0 {
		b.NodeApplyConfiguration.NodeSelector = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.NodeApplyConfiguration.NodeSelector[k] = v
	}
	return b
}
