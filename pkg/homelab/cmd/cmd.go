package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

var ConfigCmd = &cli.Command{
	Name:  "config",
	Usage: "Manage homelab configuration file",
	Action: func(cCtx *cli.Context) error {
		fmt.Println("config")
		return nil
	},
}
