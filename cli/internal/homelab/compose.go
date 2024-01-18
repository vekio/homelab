package homelab

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/cli/internal/pkg/utils"
)

func execDockerCompose(service string, command ...string) error {

	return nil
}

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Parse, resolve and render compose file in canonical format",
	Action: func(cCtx *cli.Context) error {
		service := utils.ParentCommandName(cCtx)
		err := execDockerCompose(service, "config")
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
		service := utils.ParentCommandName(cCtx)
		err := execDockerCompose(service, "pull")
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
		service := utils.ParentCommandName(cCtx)
		err := execDockerCompose(service, "up", "-d")
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
		service := utils.ParentCommandName(cCtx)
		err := execDockerCompose(service, "logs", "-f")
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
		service := utils.ParentCommandName(cCtx)
		err := execDockerCompose(service, "stop")
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
		service := utils.ParentCommandName(cCtx)
		err := execDockerCompose(service, "down", "-v")
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
		service := utils.ParentCommandName(cCtx)
		err := execDockerCompose(service, "pull")
		if err != nil {
			return err
		}

		err = execDockerCompose(service, "up", "-d")
		if err != nil {
			return err
		}
		return nil
	},
}
