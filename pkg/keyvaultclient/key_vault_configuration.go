package keyvaultclient

// KeyVaultConfiguration defines a configuration for connection to
// Azure key vault.
type KeyVaultConfiguration struct {
	VaultName string

	CloudName string

	TenantID string

	UsePodIdentity bool

	UseVMManagedIdentity bool

	VMManagedIdentityClientID string

	AADClientSecret string

	AADClientID string

	PodName string

	PodNamespace string
}
