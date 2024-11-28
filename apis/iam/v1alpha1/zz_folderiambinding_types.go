/*
Copyright 2022 YANDEX LLC
This is modified version of the software, made by the Crossplane Authors
and available at: https://github.com/crossplane-contrib/provider-jet-template

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

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FolderIAMBindingInitParameters struct {

	// ID of the folder to attach a policy to.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// An array of identities that will be granted the privilege that is specified in the role field.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/crossplane-provider-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountsRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountsSelector
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_folder_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// References to ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountsRef []v1.Reference `json:"serviceAccountsRef,omitempty" tf:"-"`

	// Selector for a list of ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountsSelector *v1.Selector `json:"serviceAccountsSelector,omitempty" tf:"-"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type FolderIAMBindingObservation struct {

	// ID of the folder to attach a policy to.
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An array of identities that will be granted the privilege that is specified in the role field.
	// Each entry can have one of the following values:
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_folder_iam_binding can be used per role.
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

type FolderIAMBindingParameters struct {

	// ID of the folder to attach a policy to.
	// +crossplane:generate:reference:type=github.com/yandex-cloud/crossplane-provider-yc/apis/resourcemanager/v1alpha1.Folder
	// +kubebuilder:validation:Optional
	FolderID *string `json:"folderId,omitempty" tf:"folder_id,omitempty"`

	// Reference to a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDRef *v1.Reference `json:"folderIdRef,omitempty" tf:"-"`

	// Selector for a Folder in resourcemanager to populate folderId.
	// +kubebuilder:validation:Optional
	FolderIDSelector *v1.Selector `json:"folderIdSelector,omitempty" tf:"-"`

	// An array of identities that will be granted the privilege that is specified in the role field.
	// Each entry can have one of the following values:
	// +crossplane:generate:reference:type=ServiceAccount
	// +crossplane:generate:reference:extractor=github.com/yandex-cloud/crossplane-provider-yc/config/iam.ServiceAccountRefValue()
	// +crossplane:generate:reference:refFieldName=ServiceAccountsRef
	// +crossplane:generate:reference:selectorFieldName=ServiceAccountsSelector
	// +kubebuilder:validation:Optional
	// +listType=set
	Members []*string `json:"members,omitempty" tf:"members,omitempty"`

	// The role that should be assigned. Only one
	// yandex_resourcemanager_folder_iam_binding can be used per role.
	// +kubebuilder:validation:Optional
	Role *string `json:"role,omitempty" tf:"role,omitempty"`

	// References to ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountsRef []v1.Reference `json:"serviceAccountsRef,omitempty" tf:"-"`

	// Selector for a list of ServiceAccount to populate members.
	// +kubebuilder:validation:Optional
	ServiceAccountsSelector *v1.Selector `json:"serviceAccountsSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	SleepAfter *float64 `json:"sleepAfter,omitempty" tf:"sleep_after,omitempty"`
}

// FolderIAMBindingSpec defines the desired state of FolderIAMBinding
type FolderIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FolderIAMBindingParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider FolderIAMBindingInitParameters `json:"initProvider,omitempty"`
}

// FolderIAMBindingStatus defines the observed state of FolderIAMBinding.
type FolderIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FolderIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// FolderIAMBinding is the Schema for the FolderIAMBindings API. Allows management of a single IAM binding for a Yandex Resource Manager folder.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type FolderIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.role) || (has(self.initProvider) && has(self.initProvider.role))",message="spec.forProvider.role is a required parameter"
	Spec   FolderIAMBindingSpec   `json:"spec"`
	Status FolderIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FolderIAMBindingList contains a list of FolderIAMBindings
type FolderIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FolderIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	FolderIAMBinding_Kind             = "FolderIAMBinding"
	FolderIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FolderIAMBinding_Kind}.String()
	FolderIAMBinding_KindAPIVersion   = FolderIAMBinding_Kind + "." + CRDGroupVersion.String()
	FolderIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(FolderIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&FolderIAMBinding{}, &FolderIAMBindingList{})
}
