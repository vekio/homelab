package root

import (
	"github.com/spf13/cobra"
	configCmd "github.com/vekio/homelab/pkg/cmd/config"
	servicesCmd "github.com/vekio/homelab/pkg/cmd/services"
)

func NewCmdRoot() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "homelab <command> <subcommand> [flags]",
		Short: "Manage docker homelab services",
	}

	// Commands
	cmd.AddCommand(configCmd.NewCmdConfig())
	cmd.AddCommand(servicesCmd.NewCmdServices())

	return cmd, nil
}
