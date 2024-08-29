package edit

import (
	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"

	_file "github.com/vekio/fs/file"
)

func NewCmdEdit(conf config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Open configuration file on your favorite editor",
		Run: func(cmd *cobra.Command, args []string) {
			_file.Edit(conf.Path())
		},
	}
	return cmd
}
