package homelab

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/secrets"
	"github.com/vekio/homelab/internal/services"
)

func (h Homelab) initImmich() (err error) {
	err = initImmichConfig(h.Config, h.Secrets)
	if err != nil {
		return err
	}

	err = initImmichSecrets(h.Config, h.Secrets)
	if err != nil {
		return err
	}

	return nil
}

func initImmichConfig(c config.HomelabConfig, s secrets.HomelabSecrets) (err error) {
	volumeDir := config.Manager.DirPath() + "/" + services.IMMICH

	err = _fs.CreateDir(volumeDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	dataDir := volumeDir + "/data"
	if err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms)); err != nil {
		return err
	}

	cacheDir := volumeDir + "/cache"
	if err = _fs.CreateDir(cacheDir, os.FileMode(_fs.DefaultDirPerms)); err != nil {
		return err
	}

	return nil
}

func initImmichSecrets(c config.HomelabConfig, s secrets.HomelabSecrets) (err error) {
	secretsDir := secrets.Manager.DirPath() + "/secrets/" + services.IMMICH
	err = _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	if err = WriteSecret(secretsDir+"/IMMICH_DB_PASS",
		s.Immich.DBPass); err != nil {
		return err
	}

	return nil
}
