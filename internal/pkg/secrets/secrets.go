package secrets

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	_file "github.com/vekio/fs/file"
	"gopkg.in/yaml.v3"
)

var Secrets secretConfig

// SoftInitSecrets checks if the secrets file exists. If not, it calls InitSecrets,
// otherwise, it loads the existing secrets file.
func SoftInitSecrets(filename string) error {
	// Check if the file exists
	exists, err := _file.Exists(filename)
	if err != nil {
		return fmt.Errorf("SoftInitSecrets: error checking file existence: %w", err)
	}

	// If the file does not exist, initialize secrets using InitSecrets
	if !exists {
		return InitSecrets(filename)
	}

	// If the file exists, load existing secrets
	if err = loadSecrets(filename); err != nil {
		return fmt.Errorf("SoftInitSecrets: error loading secrets file %s: %w", filename, err)
	}

	return nil
}

// InitSecrets initializes service secrets and saves them in the specified file.
func InitSecrets(filename string) error {
	// Generate Authelia secrets
	autheliaSecrets, err := autheliaSecrets()
	if err != nil {
		return fmt.Errorf("InitSecrets: failed generating Authelia secrets: %w", err)
	}

	// Generate Gitea secrets
	giteaSecrets, err := giteaSecrets()
	if err != nil {
		return fmt.Errorf("InitSecrets: failed generating Gitea secrets: %w", err)
	}

	// Generate Immich secrets
	immichSecrets, err := immichSecrets()
	if err != nil {
		return fmt.Errorf("InitSecrets: failed generating Immich secrets: %w", err)
	}

	// Generate Lldap secrets
	lldapSecrets, err := lldapSecrets()
	if err != nil {
		return fmt.Errorf("InitSecrets: failed generating Lldap secrets: %w", err)
	}

	// Populate the global Secrets variable with the generated secrets
	Secrets = secretConfig{
		Authelia: autheliaSecrets,
		Gitea:    giteaSecrets,
		Immich:   immichSecrets,
		Lldap:    lldapSecrets,
		Traefik:  traefik{},
	}

	// Save the generated secrets to the specified file
	if err := saveSecrets(filename); err != nil {
		return fmt.Errorf("InitSecrets: failed saving secrets: %w", err)
	}

	return nil
}

// loadSecrets loads secret data from a YAML file.
// It reads the specified file, unmarshals the YAML data, and populates
// the Secrets variable.
func loadSecrets(filename string) error {
	// Read the content of the YAML file.
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("loadSecrets: failed reading %s: %w", filename, err)
	}

	// Unmarshal YAML data into the Secrets variable.
	err = yaml.Unmarshal(yamlFile, &Secrets)
	if err != nil {
		return fmt.Errorf("loadSecrets: failed to unmarshal YAML: %w", err)
	}

	return nil
}

// saveSecrets saves the secret data to a YAML file.
// It marshals the Secrets variable into YAML format and writes it
// to the specified file.
func saveSecrets(filename string) error {
	// Marshal the Secrets variable into YAML format.
	yamlData, err := yaml.Marshal(&Secrets)
	if err != nil {
		return fmt.Errorf("saveSecrets: failed to marshal YAML: %w", err)
	}

	// Write the YAML data to the specified file with restricted file permissions.
	err = os.WriteFile(filename, yamlData, _fs.RestrictedFilePerms)
	if err != nil {
		return fmt.Errorf("saveSecrets: failed to write file: %w", err)
	}

	return nil
}
