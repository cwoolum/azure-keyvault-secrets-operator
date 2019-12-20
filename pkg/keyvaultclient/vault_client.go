package keyvaultclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"
)

func getVaultsClient(credentials KeyVaultCredentials) keyvault.VaultsClient {
	vaultsClient := keyvault.NewVaultsClient(credentials.SubscriptionID)
	a, _ := GetResourceManagementAuthorizer()
	vaultsClient.Authorizer = a
	vaultsClient.AddToUserAgent(UserAgent())
	return vaultsClient
}

// GetVault returns an existing vault
func GetVault(ctx context.Context, vaultName string) (keyvault.Vault, error) {
	vaultsClient := getVaultsClient()
	return vaultsClient.Get(ctx, GroupName(), vaultName)
}
