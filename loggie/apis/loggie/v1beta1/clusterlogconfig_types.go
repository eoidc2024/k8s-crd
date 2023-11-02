/*
Copyright 2021 Loggie Authors

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

package v1beta1

import (
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SelectorTypePod      = "pod"
	SelectorTypeNode     = "node"
	SelectorTypeCluster  = "cluster"
	SelectorTypeVm       = "vm"
	SelectorTypeWorkload = "workload"
	SelectorTypeAll      = "all"
)

// ClusterLogConfig Deployment with Datadog Operator.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=clusterlogconfigs,shortName=clgc,scope=Cluster
// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterLogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Spec   `json:"spec"`
	Status Status `json:"status,omitempty"`
}

type Spec struct {
	Selector *Selector `json:"selector,omitempty"`
	Pipeline *Pipeline `json:"pipeline,omitempty"`
}

type Selector struct {
	Cluster           string             `json:"cluster,omitempty"`
	Type              string             `json:"type,omitempty"`
	WorkloadSelector  []WorkloadSelector `json:"workload_selector,omitempty"`
	PodSelector       `json:",inline"`
	NodeSelector      `json:",inline"`
	NamespaceSelector `json:",inline"`
	EoiPodSelector    `json:",inline"`
}

type PodSelector struct {
	LabelSelector map[string]string `json:"labelSelector,omitempty"`
}

type NodeSelector struct {
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`
}

type Pipeline struct {
	Name           string `json:"name,omitempty"`
	Sources        string `json:"sources,omitempty"`
	Sink           string `json:"sink,omitempty"`
	Interceptors   string `json:"interceptors,omitempty"`
	SinkRef        string `json:"sinkRef,omitempty"`
	InterceptorRef string `json:"interceptorRef,omitempty"`
}

type Status struct {
	Message Message `json:"message,omitempty"`
}

type Message struct {
	Reason             string `json:"reason,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
}

func (in *ClusterLogConfig) Validate() error {
	if in.Spec.Pipeline == nil {
		return errors.New("spec.pipelines is required")
	}
	if in.Spec.Selector == nil {
		return errors.New("spec.selector is required")
	}

	tp := in.Spec.Selector.Type
	if tp != SelectorTypePod && tp != SelectorTypeNode && tp != SelectorTypeCluster {
		return errors.New("spec.selector.type is invalid")
	}

	if tp == SelectorTypeCluster && in.Spec.Selector.Cluster == "" {
		return errors.New("selector.cluster is required when selector.type=cluster")
	}

	if in.Spec.Pipeline.Sources == "" {
		return errors.New("pipeline sources is empty")
	}

	return nil
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterLogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterLogConfig `json:"items"`
}

func (in *ClusterLogConfig) ToLogConfig() *LogConfig {
	clgc := in.DeepCopy()

	lgc := &LogConfig{}
	lgc.Name = clgc.Name
	lgc.Labels = clgc.Labels
	lgc.Annotations = clgc.Annotations

	lgc.Spec = clgc.Spec

	return lgc
}