package services

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/homelab"
)

func NewCmdServices(homelab homelab.Homelab) *cobra.Command {
	servicesCmd := &cobra.Command{
		Use:     "services",
		Aliases: []string{"srv"},
		Short:   "Manage homelab services",
	}

	// Subcommands
	servicesCmd.AddCommand(newCmdList(homelab), newCmdUpgrade(homelab))

	// Compose Subcommands
	servicesCmd.AddGroup(&cobra.Group{ID: "compose", Title: "Compose Commands"})
	servicesCmd.AddCommand(
		newCmdConfig(homelab), newCmdDown(homelab),
		newCmdLogs(homelab), newCmdRestart(homelab),
		newCmdStop(homelab), newCmdPull(homelab),
		newCmdUp(homelab))

	return servicesCmd
}
