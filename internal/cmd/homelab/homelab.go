package homelab

import (
	"github.com/spf13/cobra"
	cmdConfig "github.com/vekio/homelab/internal/cmd/config"
	cmdServices "github.com/vekio/homelab/internal/cmd/services"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdHomelab(conf *config.ConfigManager[config.Config]) *cobra.Command {
	homelabCmd := &cobra.Command{
		Use:   "homelab",
		Short: "Manage my homelab docker services",
	}

	// Subcommands
	homelabCmd.AddCommand(cmdConfig.NewCmdConfig(conf))
	homelabCmd.AddCommand(cmdServices.NewCmdServices(conf))

	return homelabCmd
}
