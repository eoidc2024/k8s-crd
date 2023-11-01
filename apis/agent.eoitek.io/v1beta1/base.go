package v1beta1

const (
	SelectorTypePod = "pod"
)

type Selector struct {
	Cluster         string            `json:"cluster,omitempty"`
	Type            string            `json:"type,omitempty"`
	LabelSelector   map[string]string `json:"labelSelector,omitempty"`
	ContainerString string            `json:"container,omitempty"`
}

type Message struct {
	Reason             string `json:"reason,omitempty"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
}

type Status struct {
	Message Message `json:"message,omitempty"`
}
