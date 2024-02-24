package homelab

import (
	"os"

	_fs "github.com/vekio/fs"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/services"
)

func (h Homelab) initProtonmailBridge() (err error) {
	err = initProtonmailBridgeConfig(h.Services)
	if err != nil {
		return err
	}

	return nil
}

func initProtonmailBridgeConfig(srvs services.Services) (err error) {
	protonmailVolume := config.Manager.DirPath() + "/" + services.PROTONMAIL_BRIDGE

	err = _fs.CreateDir(protonmailVolume, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	err = os.Setenv("PROTONMAIL_BRIDGE_COMMAND", "init")
	if err != nil {
		return err
	}

	err = srvs[services.PROTONMAIL_BRIDGE].ExComposeCmd("run", "--rm", "protonmail-bridge")
	if err != nil {
		return err
	}

	return nil
}
