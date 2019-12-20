package keyvaultclient

import (
	"regexp"

	kv "github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"github.com/pkg/errors"
)

// GetKeyVaultClient returns a new key vault client
func (options *KeyVaultConfiguration) GetKeyVaultClient() (*kv.BaseClient, error) {
	kvClient := kv.New()

	token, err := GetKeyvaultToken(AuthGrantType(), options.CloudName, options.TenantID, options.UsePodIdentity, options.UseVMManagedIdentity, options.VMManagedIdentityClientID, options.AADClientSecret, options.AADClientID, options.PodName, options.PodNamespace)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get key vault token")
	}

	kvClient.Authorizer = token
	return &kvClient, nil
}

// GetVaultURL returns the vault url for the specific vault name.
func (adapter *KeyVaultConfiguration) GetVaultURL() (vaultURL *string, err error) {
	// See docs for validation spec: https://docs.microsoft.com/en-us/azure/key-vault/about-keys-secrets-and-certificates#objects-identifiers-and-versioning
	if match, _ := regexp.MatchString("[-a-zA-Z0-9]{3,24}", adapter.VaultName); !match {
		return nil, errors.Errorf("Invalid vault name: %q, must match [-a-zA-Z0-9]{3,24}")
	}
	vaultDNSSuffix, err := GetVaultDNSSuffix(adapter.CloudName)
	if err != nil {
		return nil, err
	}

	vaultDNSSuffixValue := *vaultDNSSuffix

	vaultURI := "https://" + adapter.VaultName + "." + vaultDNSSuffixValue + "/"
	return &vaultURI, nil
}

// GetVaultDNSSuffix gets the suffix based on the cloud used.
func GetVaultDNSSuffix(cloudName string) (vaultTld *string, err error) {
	environment, err := ParseAzureEnvironment(cloudName)

	if err != nil {
		return nil, err
	}

	return &environment.KeyVaultDNSSuffix, nil
}
