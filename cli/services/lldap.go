package services

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/cli/utils"
)

func InitLldap(envConfig string) error {
	// Create secrets folder
	secretsDir := fmt.Sprintf("%s/secrets/", envConfig)
	err := _fs.Create(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	// Generate alphanumeric secrets
	secrets := []utils.Secret{
		{Name: "LLDAP_JWT_SECRET_FILE", Length: 64},
		{Name: "LLDAP_LDAP_USER_PASS_FILE", Length: 64},
	}

	for _, secret := range secrets {
		secretFile := fmt.Sprintf("%s/%s", secretsDir, secret.Name)
		err = utils.CreateAlphanumericSecret(secretFile, secret.Length)
		if err != nil {
			return err
		}
	}

	// Create data folder
	dataDir := fmt.Sprintf("%s/data/", envConfig)
	err = _fs.Create(dataDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}
