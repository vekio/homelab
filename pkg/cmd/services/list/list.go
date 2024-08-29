package list

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdList(conf config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "Listing homelab services",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := conf.Data()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading configuration file: %v\n", err)
				os.Exit(1)
			}
			for srvName, srv := range config.Services {
				fmt.Printf("[%s]\n", srvName)
				fmt.Printf("\t* %s\n", srv.Server)
			}
		},
	}
	return cmd
}
