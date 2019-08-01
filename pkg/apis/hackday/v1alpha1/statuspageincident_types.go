package v1alpha1

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// StatusPageIncidentSpec defines the desired state of StatusPageIncident
// +k8s:openapi-gen=true
type StatusPageIncidentSpec struct {
	StartTime  time.Time `json:"startTime",omit:empty` // UTC
	EndTime    time.Time `json:"endTime",omit:empty`   // UTC
	Message    string    `json:"message"`
	Type       SPType    `json:"type"`
	Components []string  `json:"components"`
	Name       string    `json:"name"`
}

// SPType special type limiting the types supported in StatusPageIncident
type SPType string

// Scheduled type supported by StatusPageIncidentSpec
var Scheduled SPType = "scheduled"

// StatusPageIncidentStatus defines the observed state of StatusPageIncident
// +k8s:openapi-gen=true
type StatusPageIncidentStatus struct {
	Scheduled bool `json:"scheduled"`
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
