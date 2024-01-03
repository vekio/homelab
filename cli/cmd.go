package homelab

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/cli/services"
)

const command = "homelab"
const version = "v0.0.1"
const authorEmail = "alberto@casta.me"
const authorName = "Alberto Castañeiras"
const summary = "CLI for manage homelab"
const description = "Manage homelab services and their config files."

var Cmd = &cli.App{
	Name:        command,
	Usage:       summary,
	Description: description,
	UsageText:   fmt.Sprintf("%s COMMAND SUBCOMMAND", command),
	// HideHelpCommand: true,
	Authors: []*cli.Author{
		{
			Name:  authorName,
			Email: authorEmail,
		},
	},
	Version:   version,
	Compiled:  time.Now(),
	Copyright: fmt.Sprintf("%s (%s) Copyright %s\nLicense Apache-2.0", command, version, authorName),
	Commands:  commands(),
}

func commands() []*cli.Command {
	var commands []*cli.Command

	for _, srv := range services.AvailableServices() {
		commands = append(commands, serviceCmdFactory(srv))
	}

	commands = append(commands, initCmd)
	commands = append(commands, testCmd)

	return commands
}

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize required folders and config files",
	Action: func(cCtx *cli.Context) error {

		services.InitAuthelia()

		return nil
	},
}
