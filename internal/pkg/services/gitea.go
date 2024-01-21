package services

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/pkg/context"
	cmd "github.com/vekio/homelab/pkg/conf"
)

var giteaSrv = Service{
	Name:        GITEA,
	ComposeFile: fmt.Sprintf("%s/%s/compose.yml", repoConfig, GITEA),
	Context:     context.DEFAULT,
	Priority:    5,
	Init:        initGitea,
}

func initGitea() error {
	giteaConf := cmd.Config.DirPath() + "/" + GITEA

	err := _fs.CreateDir(giteaConf, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	dataDir := giteaConf + "/data"
	err = _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}
