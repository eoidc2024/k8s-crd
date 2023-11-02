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
	PathStdout = "stdout"
)

// LogConfig Deployment with Datadog Operator.
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=logconfigs,shortName=lgc,scope=Namespaced
// +genclient
type LogConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Spec   `json:"spec"`
	Status Status `json:"status,omitempty"`
}

func (in *LogConfig) Validate() error {
	if in.Spec.Pipeline == nil {
		return errors.New("spec.pipelines is required")
	}
	if in.Spec.Selector == nil {
		return errors.New("spec.selector is required")
	}

	tp := in.Spec.Selector.Type
	if tp != SelectorTypePod {
		return errors.New("only selector.type:pod is supported in LogConfig")
	}

	if in.Spec.Pipeline.Sources == "" {
		return errors.New("pipeline sources is empty")
	}

	return nil
}

// +kubebuilder:object:root=true
type LogConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LogConfig `json:"items"`
}

func (in *LogConfig) ToClusterLogConfig() *ClusterLogConfig {
	lgc := in.DeepCopy()

	out := &ClusterLogConfig{}
	out.Name = lgc.Name
	out.Labels = lgc.Labels
	out.Annotations = lgc.Annotations

	out.Spec = lgc.Spec

	return out
}
