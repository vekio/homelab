package config

import (
	"github.com/spf13/cobra"

	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/pkg/config"
)

func newCmdEdit() *cobra.Command {
	editCmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit the configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := _file.Edit(config.Path())
			if err != nil {
				return err
			}
			return nil
		},
	}
	return editCmd
}
