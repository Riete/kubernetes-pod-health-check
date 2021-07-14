package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PodHealthCheckSpec defines the desired state of PodHealthCheck
// +k8s:openapi-gen=true
type DingTalk struct {
	Webhook string `json:"webhook"`
	Secret  string `json:"secret"`
}

type HealthCheck struct {
	Interval time.Duration `json:"interval"`
	Port     string        `json:"port"`
	Path     string        `json:"path"`
	Timeout  time.Duration `json:"timeout"`
}
type PodHealthCheckSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	DingTalk      DingTalk          `json:"dingTalk"`
	Namespace     string            `json:"namespace"`
	LabelSelector map[string]string `json:"labelSelector"`
	HealthCheck   HealthCheck       `json:"healthCheck"`
}

// PodHealthCheckStatus defines the observed state of PodHealthCheck
// +k8s:openapi-gen=true
type PodHealthCheckStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodHealthCheck is the Schema for the podhealthchecks API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type PodHealthCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodHealthCheckSpec   `json:"spec,omitempty"`
	Status PodHealthCheckStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodHealthCheckList contains a list of PodHealthCheck
type PodHealthCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodHealthCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodHealthCheck{}, &PodHealthCheckList{})
}
