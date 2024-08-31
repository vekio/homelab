package show

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdShow(conf *config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display current configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			buf, err := conf.Content()
			if err != nil {
				// utils.ErrorMsg(err)
			}
			fmt.Print(string(buf))
			return nil
		},
	}
	return cmd
}
