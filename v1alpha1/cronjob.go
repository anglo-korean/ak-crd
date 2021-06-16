package v1alpha1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CronjobSpec is a resource for creating cronjobs.
//
// It wraps the usual cronjob stuff with all of the
// extra config and data and helpers we want in our
// jobs.
//
// If you just need a cronjob to run a container, this
// resource overkill :)
type CronjobSpec struct {
	Base
	Schedule string `json:"schedule"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cronjob struct {
	meta_v1.TypeMeta   `json:",inline"`
	meta_v1.ObjectMeta `json:"metadata"`
	Spec               CronjobSpec   `json:"spec"`
	Status             CronjobStatus `json:"status,omitempty"`
}

type CronjobStatus struct {
	State   string `json:"state,omitempty"`
	Message string `json:"message,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CronjobList struct {
	meta_v1.TypeMeta `json:",inline"`
	meta_v1.ListMeta `json:"metadata"`
	Items            []Cronjob `json:"items"`
}
