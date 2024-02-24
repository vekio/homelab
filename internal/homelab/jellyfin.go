package homelab

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/services"
)

func (h Homelab) initJellyfin() (err error) {
	err = initJellyfinConfig()
	if err != nil {
		return err
	}

	return nil
}

func initJellyfinConfig() (err error) {
	volumeDir := config.Manager.DirPath() + "/" + services.JELLYFIN

	err = _fs.CreateDir(volumeDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	configDir := volumeDir + "/config"
	err = _fs.CreateDir(configDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	cacheDir := volumeDir + "/cache"
	err = _fs.CreateDir(cacheDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}
