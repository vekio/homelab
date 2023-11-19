package homelab

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func ServiceCmdFactory(service string) *cli.Command {
	return &cli.Command{
		Name:  service,
		Usage: fmt.Sprintf("Manage %s service", service),
		Subcommands: []*cli.Command{
			configCmd,
			pullCmd,
			upCmd,
			logsCmd,
			stopCmd,
			downCmd,
			upgradeCmd,
		},
	}
}

// Extract service from command context
func getService(cCtx *cli.Context) string {
	return cCtx.Lineage()[1].Command.Name
}

// Executes docker compose file
func runDockerCompose(cCtx *cli.Context, command ...string) error {

	service := getService(cCtx)
	composeFile, err := ComposeFile(service)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Load compose file %s", composeFile))

	envFile, err := ComposeEnvFile(service)
	if err != nil {
		return err
	}
	slog.Debug(fmt.Sprintf("Load env file %s", envFile))

	cmd := exec.Command("docker", append([]string{"compose", "-f", composeFile, "--env-file", envFile}, command...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error execute %s docker compose: %s\n", service, err)
	}
	return nil
}

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Parse, resolve and render compose file in canonical format",
	Action: func(cCtx *cli.Context) error {

		err := runDockerCompose(cCtx, "config")
		if err != nil {
			return err
		}

		return nil
	},
}

var pullCmd = &cli.Command{
	Name:  "pull",
	Usage: "Pull service images",
	Action: func(cCtx *cli.Context) error {

		err := runDockerCompose(cCtx, "pull")
		if err != nil {
			return err
		}

		return nil
	},
}

var upCmd = &cli.Command{
	Name:  "up",
	Usage: "Create and start containers",
	Action: func(cCtx *cli.Context) error {

		err := runDockerCompose(cCtx, "up", "-d")
		if err != nil {
			return err
		}

		return nil
	},
}

var logsCmd = &cli.Command{
	Name:  "logs",
	Usage: "View output from containers",
	Action: func(cCtx *cli.Context) error {

		err := runDockerCompose(cCtx, "logs", "-f")
		if err != nil {
			return err
		}

		return nil
	},
}

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop services",
	Action: func(cCtx *cli.Context) error {

		err := runDockerCompose(cCtx, "stop")
		if err != nil {
			return err
		}

		return nil
	},
}

var downCmd = &cli.Command{
	Name:  "down",
	Usage: "Stop and remove containers, networks",
	Action: func(cCtx *cli.Context) error {

		err := runDockerCompose(cCtx, "down", "-v")
		if err != nil {
			return err
		}

		return nil
	},
}

var upgradeCmd = &cli.Command{
	Name:  "upgrade",
	Usage: "Pull service images and start containers",
	Action: func(cCtx *cli.Context) error {

		err := runDockerCompose(cCtx, "pull")
		if err != nil {
			return err
		}

		err = runDockerCompose(cCtx, "up", "-d")
		if err != nil {
			return err
		}

		return nil
	},
}
