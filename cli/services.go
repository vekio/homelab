package homelab

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	AUTHELIA string = "authelia"
	GITEA    string = "gitea"
	IMMICH   string = "immich"
	LLDAP    string = "lldap"
	TRAEFIK  string = "traefik"
)

func initServiceCommands() []*cli.Command {
	var commands []*cli.Command

	services := []string{
		AUTHELIA,
		GITEA,
		IMMICH,
		LLDAP,
		TRAEFIK,
	}

	for _, srv := range services {
		commands = append(commands, serviceCmdFactory(srv))
	}

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
	}

	switch service {
	case TRAEFIK:
		defaultCmds = append(defaultCmds, initCmd)
	case AUTHELIA:
		defaultCmds = append(defaultCmds, initCmd)
	}

	return &cli.Command{
		Name:        service,
		Usage:       fmt.Sprintf("Manage %s service", service),
		Subcommands: defaultCmds,
	}
}

func getService(cCtx *cli.Context) string {
	return cCtx.Lineage()[1].Command.Name
}
