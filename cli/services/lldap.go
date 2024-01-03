package services

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/cli/conf"
	"github.com/vekio/homelab/cli/secrets"
	"github.com/vekio/homelab/cli/utils"
)

func InitLldap() error {
	lldapConf := conf.Config.DirPath() + "/" + LLDAP

	err := _fs.CreateDir(lldapConf, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	dataDir := lldapConf + "/data"
	err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	err = initLldapSecrets(lldapConf)
	if err != nil {
		return err
	}

	return nil
}

func initLldapSecrets(lldapConf string) error {
	secretsDir := lldapConf + "/secrets"
	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/LLDAP_JWT_SECRET_FILE",
		secrets.Secrets.Lldap.JWTSecret); err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/LLDAP_LDAP_USER_PASS_FILE",
		secrets.Secrets.Lldap.LDAPUserPass); err != nil {
		return err
	}

	return nil
}

// func InitLldap(envConfig string) error {
// 	// Create secrets folder
// 	secretsDir := fmt.Sprintf("%s/secrets/", envConfig)
// 	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
// 	if err != nil {
// 		return err
// 	}

// 	// Generate alphanumeric secrets
// 	secrets := []utils.Secret{
// 		{Name: "LLDAP_JWT_SECRET_FILE", Length: 64},
// 		{Name: "LLDAP_LDAP_USER_PASS_FILE", Length: 64},
// 	}

// 	for _, secret := range secrets {
// 		secretFile := fmt.Sprintf("%s/%s", secretsDir, secret.Name)
// 		err = utils.CreateAlphanumericSecret(secretFile, secret.Length)
// 		if err != nil {
// 			return err
// 		}
// 	}

// 	// Create data folder
// 	dataDir := fmt.Sprintf("%s/data/", envConfig)
// 	err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
