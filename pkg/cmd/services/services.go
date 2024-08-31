package services

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
	cmdList "github.com/vekio/homelab/pkg/cmd/services/list"
)

func NewCmdServices(conf *config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "services <command>",
		Short: "Manage homelab services",
	}

	// Add subcommands to services command.
	cmd.AddCommand(cmdList.NewCmdList(conf))
	return cmd
}
