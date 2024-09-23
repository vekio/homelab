package config

import (
	"github.com/spf13/cobra"
)

func NewCmdConfig() *cobra.Command {
	configCmd := &cobra.Command{
		Use:     "config",
		Aliases: []string{"conf"},
		Short:   "Manage configuration for homelab",
	}

	// Subcommands
	configCmd.AddCommand(newCmdShow())
	configCmd.AddCommand(newCmdEdit())
	return configCmd
}
