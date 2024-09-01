package services

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
	cmdList "github.com/vekio/homelab/pkg/cmd/services/list"
)

func NewCmdServices(conf *config.ConfigManager) *cli.Command {
	cmd := &cli.Command{
		Name:    "services",
		Aliases: []string{"srv"},
		Usage:   "Manage homelab services",
		Subcommands: []*cli.Command{
			cmdList.NewCmdList(conf),
		},
	}
	return cmd
}
