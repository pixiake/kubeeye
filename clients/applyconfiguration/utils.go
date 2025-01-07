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

package applyconfiguration

import (
	v1alpha2 "github.com/kubesphere/kubeeye/apis/kubeeye/v1alpha2"
	internal "github.com/kubesphere/kubeeye/clients/applyconfiguration/internal"
	kubeeyev1alpha2 "github.com/kubesphere/kubeeye/clients/applyconfiguration/kubeeye/v1alpha2"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	testing "k8s.io/client-go/testing"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=kubeeye, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithKind("BaseResult"):
		return &kubeeyev1alpha2.BaseResultApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("Cluster"):
		return &kubeeyev1alpha2.ClusterApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ClusterInfo"):
		return &kubeeyev1alpha2.ClusterInfoApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("CommandResultItem"):
		return &kubeeyev1alpha2.CommandResultItemApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ComponentResultItem"):
		return &kubeeyev1alpha2.ComponentResultItemApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("CustomCommandRule"):
		return &kubeeyev1alpha2.CustomCommandRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ExtraInfo"):
		return &kubeeyev1alpha2.ExtraInfoApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("FileChangeResultItem"):
		return &kubeeyev1alpha2.FileChangeResultItemApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("FileChangeRule"):
		return &kubeeyev1alpha2.FileChangeRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("FileFilterRule"):
		return &kubeeyev1alpha2.FileFilterRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectPlan"):
		return &kubeeyev1alpha2.InspectPlanApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectPlanSpec"):
		return &kubeeyev1alpha2.InspectPlanSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectPlanStatus"):
		return &kubeeyev1alpha2.InspectPlanStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectResult"):
		return &kubeeyev1alpha2.InspectResultApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectResultSpec"):
		return &kubeeyev1alpha2.InspectResultSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectResultStatus"):
		return &kubeeyev1alpha2.InspectResultStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectRule"):
		return &kubeeyev1alpha2.InspectRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectRuleNames"):
		return &kubeeyev1alpha2.InspectRuleNamesApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectRuleSpec"):
		return &kubeeyev1alpha2.InspectRuleSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectRuleStatus"):
		return &kubeeyev1alpha2.InspectRuleStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectTask"):
		return &kubeeyev1alpha2.InspectTaskApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectTaskSpec"):
		return &kubeeyev1alpha2.InspectTaskSpecApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("InspectTaskStatus"):
		return &kubeeyev1alpha2.InspectTaskStatusApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("JobPhase"):
		return &kubeeyev1alpha2.JobPhaseApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("KubeeyeOpaResult"):
		return &kubeeyev1alpha2.KubeeyeOpaResultApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("Node"):
		return &kubeeyev1alpha2.NodeApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("NodeInfoResultItem"):
		return &kubeeyev1alpha2.NodeInfoResultItemApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("NodeInfoRule"):
		return &kubeeyev1alpha2.NodeInfoRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("NodeMetricsResultItem"):
		return &kubeeyev1alpha2.NodeMetricsResultItemApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("OpaRule"):
		return &kubeeyev1alpha2.OpaRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("PrometheusConfig"):
		return &kubeeyev1alpha2.PrometheusConfigApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("PrometheusResult"):
		return &kubeeyev1alpha2.PrometheusResultApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("PrometheusRule"):
		return &kubeeyev1alpha2.PrometheusRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ResourceResult"):
		return &kubeeyev1alpha2.ResourceResultApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ResourcesType"):
		return &kubeeyev1alpha2.ResourcesTypeApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ResultItem"):
		return &kubeeyev1alpha2.ResultItemApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("RuleItemBases"):
		return &kubeeyev1alpha2.RuleItemBasesApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ScoreInfo"):
		return &kubeeyev1alpha2.ScoreInfoApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ServiceConnectResultItem"):
		return &kubeeyev1alpha2.ServiceConnectResultItemApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("ServiceConnectRule"):
		return &kubeeyev1alpha2.ServiceConnectRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("SysRule"):
		return &kubeeyev1alpha2.SysRuleApplyConfiguration{}
	case v1alpha2.SchemeGroupVersion.WithKind("TaskNames"):
		return &kubeeyev1alpha2.TaskNamesApplyConfiguration{}

	}
	return nil
}

func NewTypeConverter(scheme *runtime.Scheme) *testing.TypeConverter {
	return &testing.TypeConverter{Scheme: scheme, TypeResolver: internal.Parser()}
}
