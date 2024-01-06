package services

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/cli/conf"
)

func InitJellyfin() error {
	jellyfinConf := conf.Config.DirPath() + "/" + JELLYFIN

	err := _fs.CreateDir(jellyfinConf, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	configDir := jellyfinConf + "/config"
	err = _fs.CreateDir(configDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	cacheDir := jellyfinConf + "/cache"
	err = _fs.CreateDir(cacheDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}
