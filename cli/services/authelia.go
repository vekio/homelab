package services

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/cli/utils"
)

func InitAuthelia(repoConfig, envConfig string) error {

	if err := initAutheliaConfig(envConfig); err != nil {
		return err
	}

	if err := initAutheliaSecrets(envConfig); err != nil {
		return err
	}

	return nil
}

func initAutheliaConfig(envConfig string) error {
	// Create config folder
	configDir := fmt.Sprintf("%s/config/", envConfig)
	err := _fs.Create(configDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	// Read configuration.yml and parse with env variables
	configurationYMLFile := fmt.Sprintf("%s/configuration.yml", envConfig)
	data := map[string]string{
		"DOMAIN": os.Getenv("DOMAIN"),
		"SLD":    os.Getenv("SLD"),
		"TLD":    os.Getenv("TLD"),
	}
	if err := parseConfigFile(configurationYMLFile, data); err != nil {
		return err
	}

	return nil
}

func initAutheliaSecrets(envConfig string) error {
	// Create secrets folder
	secretsDir := fmt.Sprintf("%s/secrets/", envConfig)

	err := _fs.Create(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	// Generate alphanumeric secrets
	secrets := []utils.Secret{
		{Name: "AUTHELIA_JWT_SECRET_FILE", Length: 64},
		{Name: "AUTHELIA_IDENTITY_PROVIDERS_OIDC_HMAC_SECRET_FILE", Length: 64},
		{Name: "AUTHELIA_SESSION_SECRET_FILE", Length: 64},
		{Name: "AUTHELIA_STORAGE_ENCRYPTION_KEY_FILE", Length: 64},
	}

	for _, secret := range secrets {
		secretFile := fmt.Sprintf("%s/%s", secretsDir, secret.Name)
		err = utils.CreateAlphanumericSecret(secretFile, secret.Length)
		if err != nil {
			return err
		}
	}

	// Generate private key secret
	keyFile := fmt.Sprintf("%s/%s", secretsDir, "AUTHELIA_IDENTITY_PROVIDERS_OIDC_ISSUER_PRIVATE_KEY_FILE")
	if err := utils.GenerateRSAPrivateKey(keyFile); err != nil {
		return err
	}

	return nil
}
