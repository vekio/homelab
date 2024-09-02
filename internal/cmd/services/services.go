package services

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdServices(conf *config.ConfigManager[config.Config]) *cli.Command {
	cmd := &cli.Command{
		Name:    "services",
		Aliases: []string{"srv"},
		Usage:   "Manage homelab services",
		Subcommands: []*cli.Command{
			newCmdList(conf),
		},
	}
	return cmd
}
