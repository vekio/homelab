package show

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdShow(conf *config.ConfigManager) *cli.Command {
	cmd := &cli.Command{
		Name:  "show",
		Usage: "Display current configuration file",
		Action: func(cCtx *cli.Context) error {
			buf, err := conf.Content()
			if err != nil {
				return err
			}
			fmt.Print(string(buf))
			return nil
		},
	}
	return cmd
}
