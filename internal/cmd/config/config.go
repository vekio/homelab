package config

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdConfig(conf *config.ConfigManager[config.Config]) *cobra.Command {
	configCmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"conf"},
		Short:   "Manage configuration for homelab",
	}

	// Subcommands
	configCmd.AddCommand(newCmdShow(conf))
	configCmd.AddCommand(newCmdEdit(conf))
	return configCmd
}
