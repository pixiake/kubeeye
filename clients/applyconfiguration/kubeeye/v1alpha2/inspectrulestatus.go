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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// InspectRuleStatusApplyConfiguration represents a declarative configuration of the InspectRuleStatus type for use
// with apply.
type InspectRuleStatusApplyConfiguration struct {
	StartImportTime *v1.Time                       `json:"startImportTime,omitempty"`
	EndImportTime   *v1.Time                       `json:"endImportTime,omitempty"`
	State           *kubeeyev1alpha2.State         `json:"state,omitempty"`
	LevelCount      map[kubeeyev1alpha2.Level]*int `json:"levelCount,omitempty"`
}

// InspectRuleStatusApplyConfiguration constructs a declarative configuration of the InspectRuleStatus type for use with
// apply.
func InspectRuleStatus() *InspectRuleStatusApplyConfiguration {
	return &InspectRuleStatusApplyConfiguration{}
}

// WithStartImportTime sets the StartImportTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartImportTime field is set to the value of the last call.
func (b *InspectRuleStatusApplyConfiguration) WithStartImportTime(value v1.Time) *InspectRuleStatusApplyConfiguration {
	b.StartImportTime = &value
	return b
}

// WithEndImportTime sets the EndImportTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EndImportTime field is set to the value of the last call.
func (b *InspectRuleStatusApplyConfiguration) WithEndImportTime(value v1.Time) *InspectRuleStatusApplyConfiguration {
	b.EndImportTime = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *InspectRuleStatusApplyConfiguration) WithState(value kubeeyev1alpha2.State) *InspectRuleStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithLevelCount puts the entries into the LevelCount field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the LevelCount field,
// overwriting an existing map entries in LevelCount field with the same key.
func (b *InspectRuleStatusApplyConfiguration) WithLevelCount(entries map[kubeeyev1alpha2.Level]*int) *InspectRuleStatusApplyConfiguration {
	if b.LevelCount == nil && len(entries) > 0 {
		b.LevelCount = make(map[kubeeyev1alpha2.Level]*int, len(entries))
	}
	for k, v := range entries {
		b.LevelCount[k] = v
	}
	return b
}
