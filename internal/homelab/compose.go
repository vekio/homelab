package homelab

import (
	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/pkg/services"
)

func serviceName(cCtx *cli.Context) string {
	return cCtx.Lineage()[1].Command.Name
}

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "Parse, resolve and render compose file in canonical format",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Config()
		}
		return nil
	},
}

var downCmd = &cli.Command{
	Name:  "down",
	Usage: "Stop and remove containers, networks",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Down()
		}
		return nil
	},
}

var logsCmd = &cli.Command{
	Name:  "logs",
	Usage: "View output from containers",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Logs()
		}
		return nil
	},
}

var pullCmd = &cli.Command{
	Name:  "pull",
	Usage: "Pull service images",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Pull()
		}
		return nil
	},
}

var restartCmd = &cli.Command{
	Name:  "restart",
	Usage: "Restart containers",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Restart()
		}
		return nil
	},
}

var stopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop services",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Stop()
		}
		return nil
	},
}

var upCmd = &cli.Command{
	Name:  "up",
	Usage: "Create and start containers",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Up()
		}
		return nil
	},
}

var upgradeCmd = &cli.Command{
	Name:  "upgrade",
	Usage: "Pull service images and start containers",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		if srv, ok := srvs[serviceName(cCtx)]; ok {
			return srv.Upgrade()
		}
		return nil
	},
}
