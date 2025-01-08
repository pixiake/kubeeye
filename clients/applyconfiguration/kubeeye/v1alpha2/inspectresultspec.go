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

// InspectResultSpecApplyConfiguration represents a declarative configuration of the InspectResultSpec type for use
// with apply.
type InspectResultSpecApplyConfiguration struct {
	InspectCluster       *ClusterApplyConfiguration                   `json:"inspectCluster,omitempty"`
	InspectRuleTotal     map[string]int                               `json:"inspectRuleTotal,omitempty"`
	PrometheusResult     []PrometheusResultApplyConfiguration         `json:"prometheusResult,omitempty"`
	OpaResult            *KubeeyeOpaResultApplyConfiguration          `json:"opaResult,omitempty"`
	NodeInfo             []NodeInfoResultItemApplyConfiguration       `json:"nodeInfo,omitempty"`
	FileChangeResult     []FileChangeResultItemApplyConfiguration     `json:"fileChangeResult,omitempty"`
	FileFilterResult     []FileChangeResultItemApplyConfiguration     `json:"fileFilterResult,omitempty"`
	SysctlResult         []NodeMetricsResultItemApplyConfiguration    `json:"sysctlResult,omitempty"`
	SystemdResult        []NodeMetricsResultItemApplyConfiguration    `json:"systemdResult,omitempty"`
	CommandResult        []CommandResultItemApplyConfiguration        `json:"commandResult,omitempty"`
	ComponentResult      []ComponentResultItemApplyConfiguration      `json:"componentResult,omitempty"`
	ServiceConnectResult []ServiceConnectResultItemApplyConfiguration `json:"serviceConnectResult,omitempty"`
}

// InspectResultSpecApplyConfiguration constructs a declarative configuration of the InspectResultSpec type for use with
// apply.
func InspectResultSpec() *InspectResultSpecApplyConfiguration {
	return &InspectResultSpecApplyConfiguration{}
}

// WithInspectCluster sets the InspectCluster field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InspectCluster field is set to the value of the last call.
func (b *InspectResultSpecApplyConfiguration) WithInspectCluster(value *ClusterApplyConfiguration) *InspectResultSpecApplyConfiguration {
	b.InspectCluster = value
	return b
}

// WithInspectRuleTotal puts the entries into the InspectRuleTotal field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the InspectRuleTotal field,
// overwriting an existing map entries in InspectRuleTotal field with the same key.
func (b *InspectResultSpecApplyConfiguration) WithInspectRuleTotal(entries map[string]int) *InspectResultSpecApplyConfiguration {
	if b.InspectRuleTotal == nil && len(entries) > 0 {
		b.InspectRuleTotal = make(map[string]int, len(entries))
	}
	for k, v := range entries {
		b.InspectRuleTotal[k] = v
	}
	return b
}

// WithPrometheusResult adds the given value to the PrometheusResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PrometheusResult field.
func (b *InspectResultSpecApplyConfiguration) WithPrometheusResult(values ...*PrometheusResultApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPrometheusResult")
		}
		b.PrometheusResult = append(b.PrometheusResult, *values[i])
	}
	return b
}

// WithOpaResult sets the OpaResult field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the OpaResult field is set to the value of the last call.
func (b *InspectResultSpecApplyConfiguration) WithOpaResult(value *KubeeyeOpaResultApplyConfiguration) *InspectResultSpecApplyConfiguration {
	b.OpaResult = value
	return b
}

// WithNodeInfo adds the given value to the NodeInfo field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NodeInfo field.
func (b *InspectResultSpecApplyConfiguration) WithNodeInfo(values ...*NodeInfoResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNodeInfo")
		}
		b.NodeInfo = append(b.NodeInfo, *values[i])
	}
	return b
}

// WithFileChangeResult adds the given value to the FileChangeResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the FileChangeResult field.
func (b *InspectResultSpecApplyConfiguration) WithFileChangeResult(values ...*FileChangeResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFileChangeResult")
		}
		b.FileChangeResult = append(b.FileChangeResult, *values[i])
	}
	return b
}

// WithFileFilterResult adds the given value to the FileFilterResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the FileFilterResult field.
func (b *InspectResultSpecApplyConfiguration) WithFileFilterResult(values ...*FileChangeResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithFileFilterResult")
		}
		b.FileFilterResult = append(b.FileFilterResult, *values[i])
	}
	return b
}

// WithSysctlResult adds the given value to the SysctlResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SysctlResult field.
func (b *InspectResultSpecApplyConfiguration) WithSysctlResult(values ...*NodeMetricsResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSysctlResult")
		}
		b.SysctlResult = append(b.SysctlResult, *values[i])
	}
	return b
}

// WithSystemdResult adds the given value to the SystemdResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the SystemdResult field.
func (b *InspectResultSpecApplyConfiguration) WithSystemdResult(values ...*NodeMetricsResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSystemdResult")
		}
		b.SystemdResult = append(b.SystemdResult, *values[i])
	}
	return b
}

// WithCommandResult adds the given value to the CommandResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the CommandResult field.
func (b *InspectResultSpecApplyConfiguration) WithCommandResult(values ...*CommandResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithCommandResult")
		}
		b.CommandResult = append(b.CommandResult, *values[i])
	}
	return b
}

// WithComponentResult adds the given value to the ComponentResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ComponentResult field.
func (b *InspectResultSpecApplyConfiguration) WithComponentResult(values ...*ComponentResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithComponentResult")
		}
		b.ComponentResult = append(b.ComponentResult, *values[i])
	}
	return b
}

// WithServiceConnectResult adds the given value to the ServiceConnectResult field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ServiceConnectResult field.
func (b *InspectResultSpecApplyConfiguration) WithServiceConnectResult(values ...*ServiceConnectResultItemApplyConfiguration) *InspectResultSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithServiceConnectResult")
		}
		b.ServiceConnectResult = append(b.ServiceConnectResult, *values[i])
	}
	return b
}