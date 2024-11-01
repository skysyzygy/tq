package auth

import (
	"context"
	"fmt"

	"github.com/99designs/keyring"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

type Keyring_Azure struct {
	client *azsecrets.Client
}

func (ka *Keyring_Azure) Connect(key_vault_name string) error {

	vaultURI := fmt.Sprintf("https://%s.vault.azure.net/", key_vault_name)

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return fmt.Errorf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault
	client, err := azsecrets.NewClient(vaultURI, cred, nil)
	if err != nil {
		return fmt.Errorf("failed to establish a connection to the Azure key vault: %v", err)
	}

	ka.client = client
	return nil
}

func (ka Keyring_Azure) Get(key string) (keyring.Item, error) {
	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	resp, err := ka.client.GetSecret(context.TODO(), key, version, nil)
	if err != nil {
		return keyring.Item{}, fmt.Errorf("failed to get the secret: %v", err)
	}

	return keyring.Item{
		Key:  key,
		Data: []byte(*resp.Value),
	}, nil
}

// No-op
func (ka Keyring_Azure) GetMetadata(key string) (keyring.Metadata, error) {
	return keyring.Metadata{}, nil
}

func (ka Keyring_Azure) Set(item keyring.Item) error {
	value := string(item.Data)
	// Create a secret
	params := azsecrets.SetSecretParameters{Value: &value}
	_, err := ka.client.SetSecret(context.TODO(), item.Key, params, nil)
	if err != nil {
		return fmt.Errorf("failed to create a secret: %v", err)
	}
	return nil
}

func (ka Keyring_Azure) Remove(key string) error {
	// Delete a secret. DeleteSecret returns when Key Vault has begun deleting the secret.
	// That can take several seconds to complete, so it may be necessary to wait before
	// performing other operations on the deleted secret.
	_, err := ka.client.DeleteSecret(context.TODO(), key, nil)
	if err != nil {
		return fmt.Errorf("failed to delete secret: %v", err)
	}
	return nil
}

func (ka Keyring_Azure) Keys() ([]string, error) {
	// List secrets
	var secrets []string
	pager := ka.client.NewListSecretsPager(nil)
	for pager.More() {
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			return nil, err
		}
		for _, secret := range page.Value {
			secrets = append(secrets, string(*secret.ID))
		}
	}
	return secrets, nil
}
