package v1beta1

import (
	"errors"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type CrontabTaskConfig struct {
	Schedule        string            `json:"schedule"`
	UserOfContainer bool              `json:"userOfContainer,omitempty"`
	Type            string            `json:"type"`
	TypeValue       string            `json:"value"`
	StartCmd        string            `json:"startCmd"`
	Env             []string          `json:"env,omitempty"`
	Args            []string          `json:"args,omitempty"`
	Timeout         uint64            `json:"timeout"`
	Concurrent      string            `json:"concurrent"`
	Fields          map[string]string `json:"fields"`
}

type CrontabSpec struct {
	Selector *Selector          `json:"selector"`
	Crontab  *CrontabTaskConfig `json:"crontab"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:path=crontabconfigs,shortName=ctc

// CrontabConfig is the Schema for the Crontab API
type CrontabConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CrontabSpec `json:"spec"`
	Status Status      `json:"status,omitempty"`
}

func (in *CrontabConfig) Validate() error {
	if in.Spec.Crontab.Schedule == "" {
		return errors.New("Crontab.Schedule is empty")
	}

	if in.Spec.Crontab.TypeValue == "" {
		return errors.New("Crontab.TypeValue is empty")
	}

	switch in.Spec.Crontab.Type {
	case "download":
		if in.Spec.Crontab.StartCmd == "" {
			return fmt.Errorf("Crontab.Type %s StartCmd empty", in.Spec.Crontab.Type)
		}
		break
	case "local":
		break
	case "content":
		break
	default:
		return fmt.Errorf("Crontab.Type is not support %s", in.Spec.Crontab.Type)
	}

	switch in.Spec.Crontab.Concurrent {
	case "skip":
		break
	case "kill":
		break
	case "noop":
		break
	default:
		return fmt.Errorf("Crontab.Concurrent is not support %s", in.Spec.Crontab.Concurrent)
	}

	return nil
}

func (in *CrontabConfig) Type() string {
	return in.Spec.Selector.Type
}

func (in *CrontabConfig) Object() runtime.Object {
	return in
}

func (in *CrontabConfig) GetName() string {
	return in.ObjectMeta.Name
}

func (in *CrontabConfig) Namespace() string {
	return in.ObjectMeta.Namespace
}

func (in *CrontabConfig) PodLabelSelector() map[string]string {
	return in.Spec.Selector.LabelSelector
}

func (in *CrontabConfig) SelectContainer() string {
	return in.Spec.Selector.ContainerString
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CrontabConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CrontabConfig `json:"items"`
}
