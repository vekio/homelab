package homelab

import (
	"github.com/spf13/cobra"
	cmdServices "github.com/vekio/homelab/internal/cmd/services"
	"github.com/vekio/homelab/pkg/config"

	"github.com/vekio/homelab/internal/homelab"
)

func NewCmdHomelab(homelab homelab.Homelab) *cobra.Command {
	homelabCmd := &cobra.Command{
		Use:   "homelab",
		Short: "Manage homelab services",
	}

	// Subcommands
	homelabCmd.AddCommand(config.Cmd)
	homelabCmd.AddCommand(cmdServices.NewCmdServices(homelab))

	return homelabCmd
}
