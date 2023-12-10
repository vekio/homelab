package homelab

import (
	"fmt"
	"log/slog"
	"os/user"

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

func composeFile(service string) (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Error getting current user: %s\n", err)
	}
	if err != nil {
		return "", err
	}
	slog.Debug("load compose file", "service", service)

	composeFilePath := fmt.Sprintf("%s/src/homelab/services/%s/compose.yml", currentUser.HomeDir, service)

	return composeFilePath, nil
}

// func composeEnvFile(service string) (string, error) {
// 	currentUser, err := user.Current()
// 	if err != nil {
// 		return "", fmt.Errorf("Error getting current user: %s\n", err)
// 	}
// 	if err != nil {
// 		return "", err
// 	}

// 	envFile, err := config.GetCurrentEnvFile()
// 	if err != nil {
// 		return "", err
// 	}
// 	slog.Debug("load env file", "env", envFile)

// 	envFilePath := fmt.Sprintf("%s/src/homelab/%s", currentUser.HomeDir, envFile)

// 	return envFilePath, nil
// }
