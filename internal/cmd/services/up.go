package services

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/homelab"
)

func newCmdUp(homelab homelab.Homelab) *cobra.Command {
	return &cobra.Command{
		Use:     "up",
		Short:   "Create and start containers",
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

			if err := service.Up(); err != nil {
				return err
			}
			return nil
		},
	}
}
