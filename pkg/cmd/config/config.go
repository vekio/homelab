package config

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
	cmdEdit "github.com/vekio/homelab/pkg/cmd/config/edit"
	cmdShow "github.com/vekio/homelab/pkg/cmd/config/show"
)

func NewCmdConfig(conf *config.ConfigManager) *cli.Command {
	cmd := &cli.Command{
		Name:    "config",
		Aliases: []string{"conf"},
		Usage:   "Manage configuration for homelab",
		Subcommands: []*cli.Command{
			cmdShow.NewCmdShow(conf),
			cmdEdit.NewCmdEdit(conf),
		},
	}
	return cmd
}
