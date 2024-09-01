package list

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdList(conf *config.ConfigManager) *cli.Command {
	cmd := &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "Listing available homelab services",
		Action: func(cCtx *cli.Context) error {
			config, err := conf.Data()
			if err != nil {
				return err
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
