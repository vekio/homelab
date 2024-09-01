package list

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdList(conf *config.ConfigManager[config.Config]) *cli.Command {
	cmd := &cli.Command{
		Name:    "list",
		Aliases: []string{"l"},
		Usage:   "Listing available homelab services",
		Action: func(cCtx *cli.Context) error {
			for srvName, srv := range conf.Data.Services {
				fmt.Printf("[%s]\n", srvName)
				fmt.Printf("\t* %s\n", srv.Context)
			}
			return nil
		},
	}
	return cmd
}
