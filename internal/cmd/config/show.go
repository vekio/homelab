package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func newCmdShow(conf *config.ConfigManager[config.Config]) *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Display current configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			buf, err := conf.Content()
			if err != nil {
				return err
			}
			fmt.Print(string(buf))
			return nil
		},
	}
	return showCmd
}
