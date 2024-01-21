package services

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/pkg/context"
	cmd "github.com/vekio/homelab/pkg/conf"
)

var protonmailBridgeSrv = Service{
	Name:        PROTONMAIL_BRIDGE,
	ComposeFile: composeFile,
	Context:     context.DEFAULT,
	Priority:    2,
	Init:        initProtonmailBridge,
}

func initProtonmailBridge() error {
	protonmailConf := cmd.Config.DirPath() + "/" + PROTONMAIL_BRIDGE

	err := _fs.CreateDir(protonmailConf, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}
