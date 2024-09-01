package edit

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"

	_file "github.com/vekio/fs/file"
)

func NewCmdEdit(conf *config.ConfigManager) *cli.Command {
	cmd := &cli.Command{
		Name:  "edit",
		Usage: "Edit the configuration file",
		Action: func(cCtx *cli.Context) error {
			err := _file.Edit(conf.Path())
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}
