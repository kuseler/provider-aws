/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ThingParameters defines the desired state of Thing
type ThingParameters struct {
	// Region is which region the Thing will be created.
	// +kubebuilder:validation:Required
	Region string `json:"region"`
	// The attribute payload, which consists of up to three name/value pairs in
	// a JSON document. For example:
	//
	// {\"attributes\":{\"string1\":\"string2\"}}
	AttributePayload *AttributePayload `json:"attributePayload,omitempty"`
	// The name of the billing group the thing will be added to.
	BillingGroupName *string `json:"billingGroupName,omitempty"`
	// The name of the thing type associated with the new thing.
	ThingTypeName         *string `json:"thingTypeName,omitempty"`
	CustomThingParameters `json:",inline"`
}

// ThingSpec defines the desired state of Thing
type ThingSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       ThingParameters `json:"forProvider"`
}

// ThingObservation defines the observed state of Thing
type ThingObservation struct {
	// The ARN of the new thing.
	ThingARN *string `json:"thingARN,omitempty"`
	// The thing ID.
	ThingID *string `json:"thingID,omitempty"`

	CustomThingObservation `json:",inline"`
}

// ThingStatus defines the observed state of Thing.
type ThingStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          ThingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Thing is the Schema for the Things API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Thing struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ThingSpec   `json:"spec"`
	Status            ThingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ThingList contains a list of Things
type ThingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Thing `json:"items"`
}

// Repository type metadata.
var (
	ThingKind             = "Thing"
	ThingGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ThingKind}.String()
	ThingKindAPIVersion   = ThingKind + "." + GroupVersion.String()
	ThingGroupVersionKind = GroupVersion.WithKind(ThingKind)
)

func init() {
	SchemeBuilder.Register(&Thing{}, &ThingList{})
}
