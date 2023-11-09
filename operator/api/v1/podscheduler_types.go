/*
Copyright 2023.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type Container struct {
	Image string `json:"image"`
	Port  int    `json:"port"`
}

type Service struct {
	Port int `json:"port"`
}

type Scheduling struct {
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=23
	StartTime int `json:"startTime"`
	// +kubebuilder:validation:Minimum=0
	// +kubebuilder:validation:Maximum=23
	EndTime int `json:"endTime"`
	// +kubebuilder:validation:Minimum=0
	Replica int `json:"replica"`
}

// PodSchedulerSpec defines the desired state of PodScheduler
type PodSchedulerSpec struct {
	// +kubebuilder:validation:Required
	Container Container `json:"container"`
	// +kubebuilder:validation:Optional
	Service Service `json:"service,omitempty"`
	// +kubebuilder:validation:Required
	SchedulingConfig []*Scheduling `json:"schedulingConfig"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Minimum=1
	DefaultReplica int32 `json:"defaultReplica"`
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=1440
	IntervalMint int32 `json:"intervalMint"`
}

// PodSchedulerStatus defines the observed state of PodScheduler
type PodSchedulerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Conditions []metav1.Condition `json:"conditions,omitempty" patchStrategy:"merge" patchMergeKey:"type" protobuf:"bytes,1,rep,name=conditions"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// PodScheduler is the Schema for the podschedulers API
type PodScheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodSchedulerSpec   `json:"spec,omitempty"`
	Status PodSchedulerStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// PodSchedulerList contains a list of PodScheduler
type PodSchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodScheduler `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PodScheduler{}, &PodSchedulerList{})
}
