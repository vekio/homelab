package show

import "github.com/spf13/cobra"

func NewCmdShow() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show",
		Short: "Display present config.yml",
	}

	return cmd
}
