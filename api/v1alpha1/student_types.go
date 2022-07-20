/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
)

// StudentSpec defines the desired state of Student
type StudentSpec struct {
	//+optional
	Phone  string  `json:"phone,omitempty"`
	//+optional
	Email  string  `json:"email,omitempty"`
	//+optional
	LastName  string  `json:"lastName,omitempty"`
    //+optional
	ClassRef corev1.ObjectReference `json:"classRef,omitempty"`
}


//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=st
//+kubebuilder:printcolumn:JSONPath=".spec.email",name="EMAIL",type="string"
//+kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// Student is the Schema for the students API
type Student struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   StudentSpec   `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// StudentList contains a list of Student
type StudentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Student `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Student{}, &StudentList{})
}
