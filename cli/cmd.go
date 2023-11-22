package homelab

import (
	"fmt"
	"time"

	"github.com/urfave/cli/v2"
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
	Commands:  initializeCommands(),
}

func initializeCommands() []*cli.Command {
	var commands []*cli.Command

	services := []string{
		AUTHELIA,
		GITEA,
		IMMICH,
		LLDAP,
		TRAEFIK,
	}

	for _, srv := range services {
		commands = append(commands, ServiceCmdFactory(srv))
	}

	return commands
}
