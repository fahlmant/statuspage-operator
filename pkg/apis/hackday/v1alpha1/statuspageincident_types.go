package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StatusPageIncidentSpec defines the desired state of StatusPageIncident
// +k8s:openapi-gen=true
type StatusPageIncidentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// StatusPageIncidentStatus defines the observed state of StatusPageIncident
// +k8s:openapi-gen=true
type StatusPageIncidentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StatusPageIncident is the Schema for the statuspageincidents API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type StatusPageIncident struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StatusPageIncidentSpec   `json:"spec,omitempty"`
	Status StatusPageIncidentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StatusPageIncidentList contains a list of StatusPageIncident
type StatusPageIncidentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StatusPageIncident `json:"items"`
}

func init() {
	SchemeBuilder.Register(&StatusPageIncident{}, &StatusPageIncidentList{})
}
