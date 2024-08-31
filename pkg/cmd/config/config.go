package config

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
	cmdEdit "github.com/vekio/homelab/pkg/cmd/config/edit"
	cmdShow "github.com/vekio/homelab/pkg/cmd/config/show"
)

func NewCmdConfig(conf *config.ConfigManager) *cobra.Command {
	// longDoc := strings.Builder{}
	// longDoc.WriteString("Display or change configuration settings for homelab.\n\n")
	// longDoc.WriteString("Current respected settings:\n") // TODO: check optiones in gh

	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "Manage configuration for homelab",
		// Long:  longDoc.String(),
	}

	// Add subcommands to config command.
	cmd.AddCommand(cmdShow.NewCmdShow(conf))
	cmd.AddCommand(cmdEdit.NewCmdEdit(conf))

	return cmd
}
