package homelab

import (
	"github.com/urfave/cli/v2"
	cmdConfig "github.com/vekio/homelab/internal/cmd/config"
	cmdServices "github.com/vekio/homelab/internal/cmd/services"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdHomelab(conf *config.ConfigManager[config.Config]) []*cli.Command {
	cmd := []*cli.Command{
		cmdConfig.NewCmdConfig(conf),
		cmdServices.NewCmdServices(conf),
	}
	return cmd
}
