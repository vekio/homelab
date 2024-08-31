package list

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdList(conf *config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "Listing homelab services",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := conf.Data()
			if err != nil {
				return fmt.Errorf("error reading configuration file: %w", err)
			}
			for srvName, srv := range config.Services {
				fmt.Printf("[%s]\n", srvName)
				fmt.Printf("\t* %s\n", srv.Server)
			}
			return nil
		},
	}
	return cmd
}
