package secrets

import (
	"os"
	"path/filepath"

	_fs "github.com/vekio/fs"
	_file "github.com/vekio/fs/file"
	"gopkg.in/yaml.v3"
)

var Manager secretsManager

func init() {
	dir, _ := os.UserConfigDir()

	Manager = secretsManager{
		dir:  dir,
		file: "secrets.yml",
		name: "homelab",
	}
}

func (sm secretsManager) DirPath() string {
	return filepath.Join(Manager.dir, Manager.name)
}

func (sm secretsManager) Path() string {
	return filepath.Join(Manager.dir, Manager.name, Manager.file)
}

func (sm secretsManager) Data() (string, error) {
	buf, err := os.ReadFile(sm.Path())
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// SoftInit checks if the secrets file exists.
// If not, it calls InitSecrets, otherwise, it loads the existing secrets file.
func (sm secretsManager) SoftInit() error {
	// Check if the file exists
	exists, err := _file.Exists(sm.Path())
	if err != nil {
		return err
	}

	// If the file does not exist, initialize secrets using InitSecrets
	if !exists {
		return sm.Init()
	}

	return nil
}

// Init initializes service secrets and saves them in the specified file.
func (sm secretsManager) Init() (err error) {

	s := &HomelabSecrets{}

	// Generate Authelia secrets
	s.Authelia, err = generateAutheliaSecrets()
	if err != nil {
		return err
	}

	// Generate Gitea secrets
	s.Gitea, err = generateGiteaSecrets()
	if err != nil {
		return err
	}

	// Generate Immich secrets
	s.Immich, err = generateImmichSecrets()
	if err != nil {
		return err
	}

	// Generate Lldap secrets
	s.Lldap, err = generateLldapSecrets()
	if err != nil {
		return err
	}

	// Marshal the Secrets variable into YAML format.
	yamlData, err := yaml.Marshal(&s)
	if err != nil {
		return err
	}

	// Write the YAML data to the specified file with restricted file permissions.
	err = os.WriteFile(sm.Path(), yamlData, _fs.RestrictedFilePerms)
	if err != nil {
		return err
	}

	return nil
}
