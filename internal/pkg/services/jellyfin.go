package services

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/pkg/context"
	cmd "github.com/vekio/homelab/pkg/conf"
)

var jellyfinSrv = Service{
	Name:        JELLYFIN,
	ComposeFile: fmt.Sprintf("%s/%s/compose.yml", repoConfig, JELLYFIN),
	Context:     context.DEFAULT,
	Priority:    6,
	Init:        initJellyfin,
}

func initJellyfin() error {
	jellyfinConf := cmd.Config.DirPath() + "/" + JELLYFIN

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
