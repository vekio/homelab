package services

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/homelab"
)

func newCmdList(homelab homelab.Homelab) *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "Listing available homelab services",
		RunE: func(cmd *cobra.Command, args []string) error {
			for srvName, srv := range homelab.Services {
				fmt.Printf("[%s]\n", srvName)
				fmt.Printf("  * %s\n", srv.Context)
			}
			return nil
		},
	}
}
