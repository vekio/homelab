package init

import (
	"strings"

	"github.com/spf13/cobra"
)

func NewCmdInit() *cobra.Command {
	longDoc := strings.Builder{}
	longDoc.WriteString("Initialize a new config.yml file if it does not already exist.\n\n")

	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize config file",
		Long:  longDoc.String(),
	}

	return cmd
}
