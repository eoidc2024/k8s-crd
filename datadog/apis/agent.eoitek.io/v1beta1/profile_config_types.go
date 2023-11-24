package v1beta1

import (
	"errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type SkywalkingProfileConfig struct {
	Service              string `json:"service,omitempty"`
	ServiceInstance      string `json:"serviceInstance,omitempty"`
	TaskId               string `json:"taskId,omitempty"`
	EndpointName         string `json:"endpointName,omitempty"`
	SerialNumber         string `json:"serialNumber,omitempty"`
	Duration             string `json:"duration,omitempty"`
	MinDurationThreshold string `json:"minDurationThreshold,omitempty"`
	DumpPeriod           string `json:"dumpPeriod,omitempty"`
	MaxSamplingCount     string `json:"maxSamplingCount,omitempty"`
	StartTime            string `json:"startTime,omitempty"`
	CreateTime           string `json:"createTime,omitempty"`
}

type Spec struct {
	Selector        *Selector                `json:"selector,omitempty"`
	Container       string                   `json:"container,omitempty"`
	Cmdline         string                   `json:"cmdline,omitempty"`
	Java            *Java                    `json:"java,omitempty"`
	Golang          *Golang                  `json:"golang,omitempty"`
	Native          *Native                  `json:"native,omitempty"`
	SwProfileConfig *SkywalkingProfileConfig `json:"swProfileConfig,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=profileconfigs,shortName=pfc

// ProfileConfig is the Schema for the Profile API
type ProfileConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   Spec   `json:"spec"`
	Status Status `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProfileConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ProfileConfig `json:"items"`
}

func (in *ProfileConfig) Validate() error {
	if in.Spec.Java == nil && in.Spec.Golang == nil && in.Spec.Native == nil && in.Spec.SwProfileConfig == nil {
		return errors.New("spec.pipelines(SwProfileConfig) is required")
	}

	if in.Spec.Selector == nil {
		return errors.New("spec.selector is required")
	}

	tp := in.Spec.Selector.Type
	if tp != SelectorTypePod {
		return errors.New("only selector.type:pod is supported in LogConfig")
	}

	return nil
}

func (in *ProfileConfig) Type() string {
	return in.Spec.Selector.Type
}

func (in *ProfileConfig) GetName() string {
	return in.ObjectMeta.Name
}

func (in *ProfileConfig) Namespace() string {
	return in.ObjectMeta.Namespace
}

func (in *ProfileConfig) PodLabelSelector() map[string]string {
	return in.Spec.Selector.LabelSelector
}

func (in *ProfileConfig) SelectContainer() string {
	return in.Spec.Selector.ContainerString
}
