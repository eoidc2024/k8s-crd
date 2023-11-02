package v1beta1

const (
	PathStderr = "stderr"
)

type NamespaceSelector struct {
	NamespaceSelector        []string `json:"namespaceSelector,omitempty"`
	ExcludeNamespaceSelector []string `json:"excludeNamespaceSelector,omitempty"`
}

type WorkloadSelector struct {
	Type                     []string `json:"type,omitempty"`
	NameSelector             []string `json:"nameSelector,omitempty"`
	NamespaceSelector        []string `json:"namespaceSelector,omitempty"`
	ExcludeNamespaceSelector []string `json:"excludeNamespaceSelector,omitempty"`
}

type LabelExpr struct {
	Key   string   `json:"key,omitempty"`
	Value []string `json:"value,omitempty"`
	Expr  string   `json:"expr,omitempty"`
}

type EoiPodSelector struct {
	EoiPodSelector []LabelExpr `json:"eoiPodSelector,omitempty"`
}
