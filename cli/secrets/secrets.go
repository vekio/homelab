package secrets

import (
	"fmt"
	"os"

	_rsa "github.com/vekio/crypto/rsa"
	_fs "github.com/vekio/fs"
	_file "github.com/vekio/fs/file"
	_sgen "github.com/vekio/rand/secretgen"
	"gopkg.in/yaml.v3"
)

type s struct {
	Authelia Authelia `yaml:"authelia"`
	Gitea    Gitea    `yaml:"gitea"`
	Lldap    Lldap    `yaml:"lldap"`
	Traefik  Traefik  `yaml:"traefik"`
}

type Authelia struct {
	JWTSecret                        string `yaml:"jwt_secret"`
	IdentityProviderOIDCHMACSecret   string `yaml:"identity_providers_oidc_hmac_secret"`
	IdentityProviderIssuerPrivateKey string `yaml:"identity_providers_issuer_private_key"`
	SessionSecret                    string `yaml:"session_secret"`
	StorageEncryptionKey             string `yaml:"storage_encryption_key"`
}

type Gitea struct {
}

type Lldap struct {
	JWTSecret    string `yaml:"jwt_secret"`
	LDAPUserPass string `yaml:"ldap_user_pass"`
}

type Traefik struct {
}

var Secrets s

// SoftInitSecrets initializes secrets if the file does not exist.
// Checks if the specified file exists, and if not, calls initSecrets
// to generate and save secrets.
func SoftInitSecrets(filename string) error {
	exists, err := _file.Exists(filename)
	if err != nil {
		return fmt.Errorf("softInitSecrets: failed to check file existence: %w", err)
	}

	if !exists {
		return InitSecrets(filename)
	}

	if err = loadSecrets(filename); err != nil {
		return err
	}

	return nil
}

// InitSecrets initializes and saves secrets to a file and saves them
// to the specified file.
func InitSecrets(filename string) error {
	autheliaSecrets, err := autheliaSecrets()
	if err != nil {
		return fmt.Errorf("initSecrets: failed to generate Authelia secrets: %w", err)
	}

	lldapSecrets, err := lldapSecrets()
	if err != nil {
		return fmt.Errorf("initSecrets: failed to generate Lldap secrets: %w", err)
	}

	Secrets = s{
		Authelia: autheliaSecrets,
		Gitea:    Gitea{},
		Lldap:    lldapSecrets,
		Traefik:  Traefik{},
	}

	if err := saveSecrets(filename); err != nil {
		return fmt.Errorf("initSecrets: failed to save secrets: %w", err)
	}

	return nil
}

// autheliaSecrets generates various secrets required for Authelia.
// It returns an Authelia struct containing the secrets and an error if any.
func autheliaSecrets() (Authelia, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate JWT secret: %w", err)
	}

	identityProviderOIDCHMACSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate OIDC HMAC secret: %w", err)
	}

	keyPair, err := _rsa.GenerateKeyPair(4096)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate RSA key pair: %w", err)
	}
	identityProviderIssuerPrivateKey, err := _rsa.PrivateKeyData(keyPair)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to extract private key data: %w", err)
	}

	sessionSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate session secret: %w", err)
	}

	storageEncryptionKey, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Authelia{}, fmt.Errorf("autheliaSecrets: failed to generate storage encryption key: %w", err)
	}

	autheliaSecrets := Authelia{
		JWTSecret:                        jwtSecret,
		IdentityProviderOIDCHMACSecret:   identityProviderOIDCHMACSecret,
		IdentityProviderIssuerPrivateKey: identityProviderIssuerPrivateKey,
		SessionSecret:                    sessionSecret,
		StorageEncryptionKey:             storageEncryptionKey,
	}

	return autheliaSecrets, nil
}

// lldapSecrets generates various secrets required for LLDAP.
func lldapSecrets() (Lldap, error) {
	jwtSecret, err := _sgen.GenerateRandomAlphaNumeric(64)
	if err != nil {
		return Lldap{}, fmt.Errorf("lldapSecrets: failed to generate JWT secret: %w", err)
	}

	ldapUserPass, err := _sgen.GenerateRandomAlphaNumeric(16)
	if err != nil {
		return Lldap{}, fmt.Errorf("lldapSecrets: failed to generate LDAP user password: %w", err)
	}

	lldapSecrets := Lldap{
		JWTSecret:    jwtSecret,
		LDAPUserPass: ldapUserPass,
	}

	return lldapSecrets, nil
}

// loadSecrets loads secrets data from a YAML file.
// It reads the file, unmarshals the YAML data, and populates the 'secrets' variable.
func loadSecrets(filename string) error {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("loadSecrets: failed to read file %s: %w", filename, err)
	}

	err = yaml.Unmarshal(yamlFile, &Secrets)
	if err != nil {
		return fmt.Errorf("loadSecrets: failed to unmarshal YAML: %w", err)
	}

	return nil
}

// saveSecrets saves the secrets data to a YAML file.
// It marshals the secrets into YAML format and writes it to the specified file.
func saveSecrets(filename string) error {
	yamlData, err := yaml.Marshal(&Secrets)
	if err != nil {
		return fmt.Errorf("saveSecrets: failed to marshal YAML: %w", err)
	}

	err = os.WriteFile(filename, yamlData, _fs.RestrictedFilePerms)
	if err != nil {
		return fmt.Errorf("saveSecrets: failed to write file: %w", err)
	}

	return nil
}
