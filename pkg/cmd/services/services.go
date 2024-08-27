package services

import (
	"github.com/spf13/cobra"
)

func NewCmdServices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "services command",
		Short: "Manage homelab available services",
	}

	return cmd
}
