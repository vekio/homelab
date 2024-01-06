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

	commands = append(commands, initCmd, allUpCmd, allDownCmd)
	commands = append(commands, testCmd)

	return commands
}

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize required folders and config files",
	Action: func(cCtx *cli.Context) error {

		if err := services.InitAuthelia(); err != nil {
			return err
		}

		if err := services.InitGitea(); err != nil {
			return err
		}

		if err := services.InitJellyfin(); err != nil {
			return err
		}

		if err := services.InitLldap(); err != nil {
			return err
		}

		if err := services.InitTraefik(); err != nil {
			return err
		}

		return nil
	},
}

var allUpCmd = &cli.Command{
	Name:    "allup",
	Aliases: []string{"au"},
	Usage:   "Create and start all service containers",
	Action: func(cCtx *cli.Context) (err error) {
		// Order by priority
		err = execDockerCompose(services.TRAEFIK, "up", "-d")
		if err != nil {
			return err
		}

		err = execDockerCompose(services.LLDAP, "up", "-d")
		if err != nil {
			return err
		}

		err = execDockerCompose(services.AUTHELIA, "up", "-d")
		if err != nil {
			return err
		}

		err = execDockerCompose(services.GITEA, "up", "-d")
		if err != nil {
			return err
		}

		return
	},
}

var allDownCmd = &cli.Command{
	Name:    "alldown",
	Aliases: []string{"ad"},
	Usage:   "Stop and remove services containers, networks and volumes",
	Action: func(cCtx *cli.Context) (err error) {
		// Order by less priority
		err = execDockerCompose(services.GITEA, "down", "-v")
		if err != nil {
			return err
		}

		err = execDockerCompose(services.AUTHELIA, "down", "-v")
		if err != nil {
			return err
		}

		err = execDockerCompose(services.LLDAP, "down", "-v")
		if err != nil {
			return err
		}

		err = execDockerCompose(services.TRAEFIK, "down", "-v")
		if err != nil {
			return err
		}

		return
	},
}
