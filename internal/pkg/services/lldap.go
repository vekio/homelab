package services

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/pkg/context"
	"github.com/vekio/homelab/internal/pkg/secrets"
	"github.com/vekio/homelab/internal/pkg/utils"
	cmd "github.com/vekio/homelab/pkg/conf"
)

var lldapSrv = Service{
	Name:        LLDAP,
	ComposeFile: fmt.Sprintf("%s/%s/compose.yml", repoConfig, LLDAP),
	Context:     context.DEFAULT,
	Priority:    3,
	Init:        initLldap,
}

func initLldap() error {
	lldapConf := cmd.Config.DirPath() + "/" + LLDAP

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

	if err = utils.WriteSecret(secretsDir+"/LLDAP_SMTP_OPTIONS__PASSWORD_FILE",
		os.Getenv("SMTP_PASSWORD")); err != nil {
		return err
	}

	return nil
}
