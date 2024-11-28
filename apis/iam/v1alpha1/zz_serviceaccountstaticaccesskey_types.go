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

type ServiceAccountStaticAccessKeyInitParameters struct {

	// The description of the service account static key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// ID of the service account which is used to get a static key.
	// +crossplane:generate:reference:type=ServiceAccount
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}

type ServiceAccountStaticAccessKeyObservation struct {

	// ID of the static access key.
	AccessKey *string `json:"accessKey,omitempty" tf:"access_key,omitempty"`

	// Creation timestamp of the static access key.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// The description of the service account static key.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The encrypted secret, base64 encoded. This is only populated when pgp_key is supplied.
	EncryptedSecretKey *string `json:"encryptedSecretKey,omitempty" tf:"encrypted_secret_key,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The fingerprint of the PGP key used to encrypt the secret key. This is only populated when pgp_key is supplied.
	KeyFingerprint *string `json:"keyFingerprint,omitempty" tf:"key_fingerprint,omitempty"`

	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// ID of the service account which is used to get a static key.
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`
}

type ServiceAccountStaticAccessKeyParameters struct {

	// The description of the service account static key.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An optional PGP key to encrypt the resulting secret key material. May either be a base64-encoded public key or a keybase username in the form keybase:keybaseusername.
	// +kubebuilder:validation:Optional
	PgpKey *string `json:"pgpKey,omitempty" tf:"pgp_key,omitempty"`

	// ID of the service account which is used to get a static key.
	// +crossplane:generate:reference:type=ServiceAccount
	// +kubebuilder:validation:Optional
	ServiceAccountID *string `json:"serviceAccountId,omitempty" tf:"service_account_id,omitempty"`

	// Reference to a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDRef *v1.Reference `json:"serviceAccountIdRef,omitempty" tf:"-"`

	// Selector for a ServiceAccount to populate serviceAccountId.
	// +kubebuilder:validation:Optional
	ServiceAccountIDSelector *v1.Selector `json:"serviceAccountIdSelector,omitempty" tf:"-"`
}

// ServiceAccountStaticAccessKeySpec defines the desired state of ServiceAccountStaticAccessKey
type ServiceAccountStaticAccessKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceAccountStaticAccessKeyParameters `json:"forProvider"`
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
	InitProvider ServiceAccountStaticAccessKeyInitParameters `json:"initProvider,omitempty"`
}

// ServiceAccountStaticAccessKeyStatus defines the observed state of ServiceAccountStaticAccessKey.
type ServiceAccountStaticAccessKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceAccountStaticAccessKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ServiceAccountStaticAccessKey is the Schema for the ServiceAccountStaticAccessKeys API. Allows management of a Yandex.Cloud IAM service account static access key.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,yandex-cloud}
type ServiceAccountStaticAccessKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountStaticAccessKeySpec   `json:"spec"`
	Status            ServiceAccountStaticAccessKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceAccountStaticAccessKeyList contains a list of ServiceAccountStaticAccessKeys
type ServiceAccountStaticAccessKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceAccountStaticAccessKey `json:"items"`
}

// Repository type metadata.
var (
	ServiceAccountStaticAccessKey_Kind             = "ServiceAccountStaticAccessKey"
	ServiceAccountStaticAccessKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceAccountStaticAccessKey_Kind}.String()
	ServiceAccountStaticAccessKey_KindAPIVersion   = ServiceAccountStaticAccessKey_Kind + "." + CRDGroupVersion.String()
	ServiceAccountStaticAccessKey_GroupVersionKind = CRDGroupVersion.WithKind(ServiceAccountStaticAccessKey_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceAccountStaticAccessKey{}, &ServiceAccountStaticAccessKeyList{})
}
