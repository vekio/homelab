package services

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/homelab"
)

func newCmdPull(homelab homelab.Homelab) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "Parse, resolve and render compose file in canonical format",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			// Check if the first arg is a valid service.
			service := strings.TrimSpace(args[0])
			_, exists := homelab.Services[service]
			if !exists {
				return fmt.Errorf("invalid service: %s not exists", service)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
}
