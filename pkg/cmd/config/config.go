package config

import (
	"strings"

	"github.com/spf13/cobra"
	cmdInit "github.com/vekio/homelab/pkg/cmd/config/init"
	cmdShow "github.com/vekio/homelab/pkg/cmd/config/show"
)

func NewCmdConfig() *cobra.Command {
	longDoc := strings.Builder{}
	longDoc.WriteString("Display or change configuration settings for homelab.\n\n")
	// longDoc.WriteString("Current respected settings:\n") // TODO: check optiones in gh

	cmd := &cobra.Command{
		Use:   "config <command>",
		Short: "Manage configuration for homelab",
		Long:  longDoc.String(),
	}

	// Subcommands
	cmd.AddCommand(cmdShow.NewCmdShow())
	cmd.AddCommand(cmdInit.NewCmdInit())

	return cmd
}
