package root

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"

	configCmd "github.com/vekio/homelab/pkg/cmd/config"
	servicesCmd "github.com/vekio/homelab/pkg/cmd/services"
)

func NewCmdRoot(conf *config.ConfigManager) (*cobra.Command, error) {
	// Create the root command for the CLI application.
	cmd := &cobra.Command{
		Use:   "homelab <command> <subcommand> [flags]",
		Short: "Manage docker homelab services",
	}

	// Add subcommands to the root command.
	cmd.AddCommand(configCmd.NewCmdConfig(conf))
	cmd.AddCommand(servicesCmd.NewCmdServices(conf))

	return cmd, nil
}
