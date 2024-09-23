package homelab

import (
	"github.com/spf13/cobra"
	cmdConfig "github.com/vekio/homelab/internal/cmd/config"
	cmdServices "github.com/vekio/homelab/internal/cmd/services"
	"github.com/vekio/homelab/internal/homelab"
)

func NewCmdHomelab(homelab homelab.Homelab) *cobra.Command {
	homelabCmd := &cobra.Command{
		Use:   "homelab",
		Short: "Manage my homelab docker services",
	}

	// Subcommands
	homelabCmd.AddCommand(cmdConfig.NewCmdConfig())
	homelabCmd.AddCommand(cmdServices.NewCmdServices(homelab))

	return homelabCmd
}
