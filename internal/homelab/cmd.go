package homelab

import (
	"cmp"
	"slices"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/pkg/services"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize required folders and config files",
	Action: func(cCtx *cli.Context) error {
		var srvs []services.Service

		// Order by priority
		for _, srv := range services.Available() {
			srvs = append(srvs, srv)
		}

		slices.SortFunc(srvs, func(a, b services.Service) int {
			return cmp.Compare(a.Priority, b.Priority)
		})

		for _, srv := range srvs {
			if srv.Name == services.PROTONMAIL_BRIDGE {
				continue
				// err := os.Setenv("PROTONMAIL_BRIDGE_COMMAND", "init")
				// if err != nil {
				// 	return err
				// }
				// err = srv.ExComposeCmd("run", "--rm", "protonmail-bridge")
				// if err != nil {
				// 	return err
				// }
			}
			srv.Init()
		}

		return nil
	},
}

var allDownCmd = &cli.Command{
	Name:    "alldown",
	Aliases: []string{"ad"},
	Usage:   "Stop and remove services containers, networks and volumes",
	Action: func(cCtx *cli.Context) error {
		var srvs []services.Service

		// Order by priority and reverse
		for _, srv := range services.Available() {
			srvs = append(srvs, srv)
		}

		slices.SortFunc(srvs, func(a, b services.Service) int {
			return cmp.Compare(a.Priority, b.Priority)
		})
		slices.Reverse(srvs)

		for _, srv := range srvs {
			srv.Down()
		}
		return nil
	},
}

var allUpCmd = &cli.Command{
	Name:    "allup",
	Aliases: []string{"au"},
	Usage:   "Create and start all service containers",
	Action: func(cCtx *cli.Context) error {
		var srvs []services.Service

		// Order by priority
		for _, srv := range services.Available() {
			srvs = append(srvs, srv)
		}
		slices.SortFunc(srvs, func(a, b services.Service) int {
			return cmp.Compare(a.Priority, b.Priority)
		})

		for _, srv := range srvs {
			srv.Up()
		}
		return nil
	},
}
