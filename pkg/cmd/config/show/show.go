package show

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/config"
)

func NewCmdShow(conf config.ConfigManager) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display current configuration file",
		Run: func(cmd *cobra.Command, args []string) {
			buf, err := conf.Content()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading configuration file: %v\n", err)
				os.Exit(1)
			}
			fmt.Print(string(buf))
		},
	}
	return cmd
}
