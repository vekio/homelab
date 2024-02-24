package homelab

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/services"
)

func (h Homelab) initGitea() (err error) {
	err = initGiteaConfig()
	if err != nil {
		return err
	}

	return nil
}

func initGiteaConfig() (err error) {
	volumeDir := config.Manager.DirPath() + "/" + services.GITEA

	err = _fs.CreateDir(volumeDir, os.FileMode(_fs.DefaultDirPerms))
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
