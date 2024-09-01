package homelab

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
	cmdConfig "github.com/vekio/homelab/pkg/cmd/config"
	cmdServices "github.com/vekio/homelab/pkg/cmd/services"
)

func NewCmdHomelab(conf *config.ConfigManager[config.Config]) []*cli.Command {
	cmd := []*cli.Command{
		cmdConfig.NewCmdConfig(conf),
		cmdServices.NewCmdServices(conf),
	}
	return cmd
}
