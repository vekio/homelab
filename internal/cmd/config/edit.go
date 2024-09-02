package config

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"

	_file "github.com/vekio/fs/file"
)

func newCmdEdit(conf *config.ConfigManager[config.Config]) *cobra.Command {
	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit the configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := _file.Edit(conf.Path())
			if err != nil {
				return err
			}
			return nil
		},
	}
	return editCmd
}
