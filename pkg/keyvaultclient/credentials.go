package keyvaultclient

// KeyVaultCredentials defines a common way to connect to Key Vault
type KeyVaultCredentials struct {
	VaultName      string
	TenantID       string
	ClientID       string
	ClientSecret   string
	SubscriptionID string
}
