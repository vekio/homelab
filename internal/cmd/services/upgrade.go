package services

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/homelab"
)

func newCmdUpgrade(homelab homelab.Homelab) *cobra.Command {
	return &cobra.Command{
		Use:     "upgrade",
		Short:   "Upgrade containers version",
		Args:    cobra.ExactArgs(1),
		GroupID: "compose",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Check if the first arg is a valid service.
			service := strings.TrimSpace(args[0])
			_, exists := homelab.Services[service]
			if !exists {
				return fmt.Errorf("service %s not exists", service)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			service := homelab.Services[strings.TrimSpace(args[0])]

			if err := service.Upgrade(); err != nil {
				return err
			}
			return nil
		},
	}
}
