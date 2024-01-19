package secrets

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	_file "github.com/vekio/fs/file"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/yaml.v3"
)

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

	giteaSecrets, err := giteaSecrets()
	if err != nil {
		return fmt.Errorf("initSecrets: failed to generate Gitea secrets: %w", err)
	}

	immichSecrets, err := immichSecrets()
	if err != nil {
		return fmt.Errorf("initSecrets: failed to generate Immich secrets: %w", err)
	}

	lldapSecrets, err := lldapSecrets()
	if err != nil {
		return fmt.Errorf("initSecrets: failed to generate Lldap secrets: %w", err)
	}

	Secrets = s{
		Authelia: autheliaSecrets,
		Gitea:    giteaSecrets,
		Immich:   immichSecrets,
		Lldap:    lldapSecrets,
		Traefik:  Traefik{},
	}

	if err := saveSecrets(filename); err != nil {
		return fmt.Errorf("initSecrets: failed to save secrets: %w", err)
	}

	return nil
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

func WriteSecret(filename, secret string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(secret)
	if err != nil {
		return err
	}

	return nil
}

func Bcrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedString := string(hashedBytes)
	return hashedString, nil
}
