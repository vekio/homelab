package homelab

import (
	"fmt"

	"github.com/urfave/cli/v2"
	s "github.com/vekio/homelab/cli/services"
)

func initServiceCommands() []*cli.Command {
	var commands []*cli.Command

	for _, srv := range s.AvailableServices() {
		commands = append(commands, serviceCmdFactory(srv))
	}

	commands = append(commands, testCmd)

	return commands
}

func serviceCmdFactory(service string) *cli.Command {

	defaultCmds := []*cli.Command{
		configCmd,
		pullCmd,
		upCmd,
		logsCmd,
		stopCmd,
		downCmd,
		upgradeCmd,
		initCmd,
	}

	return &cli.Command{
		Name:        service,
		Usage:       fmt.Sprintf("Manage %s service", service),
		Subcommands: defaultCmds,
	}
}
