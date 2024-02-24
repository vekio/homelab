package homelab

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/secrets"
	"github.com/vekio/homelab/internal/services"
)

func (h Homelab) initLldap() (err error) {
	err = initLldapConfig(h.Config)
	if err != nil {
		return err
	}

	err = initLldapSecrets(h.Secrets)
	if err != nil {
		return err
	}

	return nil
}

func initLldapConfig(hc config.HomelabConfig) error {
	volumeDir := config.Manager.DirPath() + "/" + services.LLDAP

	err := _fs.CreateDir(volumeDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	dataDir := volumeDir + "/data"
	err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}

func initLldapSecrets(s secrets.HomelabSecrets) error {
	secretsDir := secrets.Manager.DirPath() + "/secrets/" + services.LLDAP

	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/LLDAP_JWT_SECRET_FILE",
		s.Lldap.JWTSecret); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/LLDAP_LDAP_USER_PASS_FILE",
		s.Lldap.LDAPUserPass); err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/LLDAP_SMTP_OPTIONS__PASSWORD_FILE",
		os.Getenv("SMTP_PASSWORD")); err != nil {
		return err
	}

	return nil
}
