package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeyVaultMapSpec defines the desired state of KeyVaultMap
// +k8s:openapi-gen=true
type KeyVaultMapSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// VaultName is the name of the vault to fetch secrets from
	VaultName string `json:"vaultName"`

	TenantID string `json:"tenantID"`

	Credentials KeyVaultCredentialsSpec `json:"credentials"`

	// Mappings are the mapping of Key Vault secrets to environment variables
	// +listType=set
	Mappings []KeyVaultSecretBindingSpec `json:"mappings"`
}

// KeyVaultCredentialsSpec defines the credentials to connect to a key
// vault instance. This is used if not using managed identity.
type KeyVaultCredentialsSpec struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

// KeyVaultSecretBindingSpec defines the desired state of KeyVaultMap
// +k8s:openapi-gen=true
type KeyVaultSecretBindingSpec struct {
	SecretName                 string `json:"secretName"`
	EnvironmentVariableMapping string `json:"environmentVariableMapping"`
}

// KeyVaultMapStatus defines the observed state of KeyVaultMap
// +k8s:openapi-gen=true
type KeyVaultMapStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeyVaultMap is the Schema for the keyvaultmaps API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=keyvaultmaps,scope=Namespaced
type KeyVaultMap struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyVaultMapSpec   `json:"spec,omitempty"`
	Status KeyVaultMapStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeyVaultMapList contains a list of KeyVaultMap
type KeyVaultMapList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyVaultMap `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeyVaultMap{}, &KeyVaultMapList{})
}
