package services

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func newCmdPull(conf *config.ConfigManager[config.Config]) *cobra.Command {
	pullCmd := &cobra.Command{
		Use:   "pull",
		Short: "Download an image from a registry",
		RunE: func(cmd *cobra.Command, args []string) error {
			conf.DirPath()
			return nil
		},
	}
	return pullCmd
}
