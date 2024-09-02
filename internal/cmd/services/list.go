package services

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func newCmdList(conf *config.ConfigManager[config.Config]) *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "Listing available homelab services",
		RunE: func(cmd *cobra.Command, args []string) error {
			for srvName, srv := range conf.Data.Services {
				fmt.Printf("[%s]\n", srvName)
				fmt.Printf("\t* %s\n", srv.Context)
			}
			return nil
		},
	}
	return listCmd
}
