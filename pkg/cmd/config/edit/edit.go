package edit

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"

	_file "github.com/vekio/fs/file"
)

func NewCmdEdit(conf *config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "edit",
		Short: "Edit the configuration file in the default system editor",
		Run: func(cmd *cobra.Command, args []string) {
			err := _file.Edit(conf.Path())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error opening configuration file: %v\n", err)
				os.Exit(1)
			}
		},
	}
	return cmd
}
