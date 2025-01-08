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

// CommandResultItemApplyConfiguration represents a declarative configuration of the CommandResultItem type for use
// with apply.
type CommandResultItemApplyConfiguration struct {
	BaseResultApplyConfiguration `json:",inline"`
	Command                      *string `json:"command,omitempty"`
	Value                        *string `json:"value,omitempty"`
	NodeName                     *string `json:"nodeName,omitempty"`
}

// CommandResultItemApplyConfiguration constructs a declarative configuration of the CommandResultItem type for use with
// apply.
func CommandResultItem() *CommandResultItemApplyConfiguration {
	return &CommandResultItemApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *CommandResultItemApplyConfiguration) WithName(value string) *CommandResultItemApplyConfiguration {
	b.BaseResultApplyConfiguration.Name = &value
	return b
}

// WithAssert sets the Assert field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Assert field is set to the value of the last call.
func (b *CommandResultItemApplyConfiguration) WithAssert(value bool) *CommandResultItemApplyConfiguration {
	b.BaseResultApplyConfiguration.Assert = &value
	return b
}

// WithLevel sets the Level field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Level field is set to the value of the last call.
func (b *CommandResultItemApplyConfiguration) WithLevel(value kubeeyev1alpha2.Level) *CommandResultItemApplyConfiguration {
	b.BaseResultApplyConfiguration.Level = &value
	return b
}

// WithCommand sets the Command field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Command field is set to the value of the last call.
func (b *CommandResultItemApplyConfiguration) WithCommand(value string) *CommandResultItemApplyConfiguration {
	b.Command = &value
	return b
}

// WithValue sets the Value field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Value field is set to the value of the last call.
func (b *CommandResultItemApplyConfiguration) WithValue(value string) *CommandResultItemApplyConfiguration {
	b.Value = &value
	return b
}

// WithNodeName sets the NodeName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeName field is set to the value of the last call.
func (b *CommandResultItemApplyConfiguration) WithNodeName(value string) *CommandResultItemApplyConfiguration {
	b.NodeName = &value
	return b
}