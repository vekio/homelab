package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/pkg/config"
)

func newCmdShow() *cobra.Command {
	showCmd := &cobra.Command{
		Use:   "show",
		Short: "Display current configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			buf, err := config.Content()
			if err != nil {
				return err
			}
			fmt.Print(string(buf))
			return nil
		},
	}
	return showCmd
}
