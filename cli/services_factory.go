package homelab

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func serviceCmdFactory(service string) *cli.Command {

	defaultCmds := []*cli.Command{
		configCmd,
		pullCmd,
		upCmd,
		logsCmd,
		stopCmd,
		downCmd,
		upgradeCmd,
		// initCmd,
	}

	return &cli.Command{
		Name:        service,
		Usage:       fmt.Sprintf("Manage %s service", service),
		Subcommands: defaultCmds,
	}
}
