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

var immichSrv = Service{
	Name:        IMMICH,
	ComposeFile: fmt.Sprintf("%s/%s/compose.yml", repoConfig, IMMICH),
	Context:     context.DEFAULT,
	Priority:    7,
	Init: func() error {
		return nil
	},
}

func initImmich() error {
	immichConf := cmd.Config.DirPath() + "/" + IMMICH

	err := _fs.CreateDir(immichConf, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	dataDir := immichConf + "/data"
	if err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms)); err != nil {
		return err
	}

	cacheDir := immichConf + "/cache"
	if err = _fs.CreateDir(cacheDir, os.FileMode(_fs.DefaultDirPerms)); err != nil {
		return err
	}

	err = initImmichSecrets(immichConf)
	if err != nil {
		return err
	}

	return nil
}

func initImmichSecrets(immichConf string) error {
	secretsDir := immichConf + "/secrets"
	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	if err = utils.WriteSecret(secretsDir+"/IMMICH_DB_PASS",
		secrets.Secrets.Immich.DBPass); err != nil {
		return err
	}

	return nil
}
