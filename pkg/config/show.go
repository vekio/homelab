package config

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newCmdShow() *cobra.Command {
	return &cobra.Command{
		Use:   "show",
		Short: "Display the content of the current configuration file",
		RunE: func(cmd *cobra.Command, args []string) error {
			buf, err := config.content()
			if err != nil {
				return err
			}
			fmt.Print(string(buf))
			return nil
		},
	}
}
