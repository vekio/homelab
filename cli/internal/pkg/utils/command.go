package utils

import "github.com/urfave/cli/v2"

func ParentCommandName(cCtx *cli.Context) string {
	return cCtx.Lineage()[1].Command.Name
}
