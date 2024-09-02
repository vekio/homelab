package services

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdServices(conf *config.ConfigManager[config.Config]) *cobra.Command {
	servicesCmd := &cobra.Command{
		Use:     "services",
		Aliases: []string{"srv"},
		Short:   "Manage homelab services",
	}

	// Subcommands
	servicesCmd.AddCommand(newCmdList(conf))
	servicesCmd.AddCommand(newCmdPull(conf))
	return servicesCmd
}
