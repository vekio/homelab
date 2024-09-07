package services

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/utils"
)

func newCmdConfig(conf *config.ConfigManager[config.Config]) *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "config",
		Short: "Parse, resolve and render compose file in canonical format",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Check if the first arg is a valid service.
			return utils.ValidateService(conf.Data, strings.TrimSpace(args[0]))
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	return pullCmd
}
