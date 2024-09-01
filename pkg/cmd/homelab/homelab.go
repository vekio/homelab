package homelab

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
	cmdConfig "github.com/vekio/homelab/pkg/cmd/config"
)

func NewCmdHomelab(conf *config.ConfigManager[config.Config]) []*cli.Command {
	cmd := []*cli.Command{
		cmdConfig.NewCmdConfig(conf),
	}
	return cmd
}
