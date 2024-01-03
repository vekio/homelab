package services

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	_dir "github.com/vekio/fs/dir"
	"github.com/vekio/homelab/cli/utils"
)

func InitAuthelia(repoConfig, envConfig string) error {

	if err := initAutheliaConfig(repoConfig, envConfig); err != nil {
		return err
	}

	if err := initAutheliaSecrets(envConfig); err != nil {
		return err
	}

	return nil
}

func initAutheliaConfig(repoConfig, envConfig string) error {
	// Create config folder
	autheliaConfig := fmt.Sprintf("%s/config/", envConfig)
	err := _dir.Copy(repoConfig, autheliaConfig)
	if err != nil {
		return err
	}

	// Read LLDAP_LDAP_USER_PASS_FILE from lldap config folder
	// TODO try again secrets
	// ldapPassFile := fmt.Sprintf("%s/lldap/secrets/LLDAP_LDAP_USER_PASS_FILE", filepath.Dir(envConfig))
	// ldapPass, err := os.ReadFile(ldapPassFile)
	// if err != nil {
	// 	return err
	// }

	// Read configuration.yml and parse with env variables
	configurationYMLFile := fmt.Sprintf("%s/configuration.yml", autheliaConfig)
	data := map[string]string{
		"DOMAIN": os.Getenv("DOMAIN"),
		"SLD":    os.Getenv("SLD"),
		"TLD":    os.Getenv("TLD"),
		// "AUTHELIA_AUTHENTICATION_BACKEND_LDAP_PASSWORD": string(ldapPass),
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
