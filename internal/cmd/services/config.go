package services

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/homelab"
)

func newCmdConfig(homelab homelab.Homelab) *cobra.Command {
	return &cobra.Command{
		Use:     "config",
		Short:   "Parse, resolve and render compose file in canonical format",
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

			if err := service.Config(); err != nil {
				return err
			}
			return nil
		},
	}
}
