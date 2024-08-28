package show

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

// NewCmdShow creates a new Cobra command to display the current configuration file.
func NewCmdShow(conf config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display present config.yml",
		Run: func(cmd *cobra.Command, args []string) {
			buf, err := conf.Content()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error reading configuration: %v\n", err)
				return
			}
			fmt.Print(string(buf))
		},
	}
	return cmd
}
