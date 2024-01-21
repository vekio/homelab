package homelab

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/pkg/services"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize required folders and config files",
	Action: func(cCtx *cli.Context) error {
		srvs := services.Available()
		for _, srv := range srvs {
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

		fmt.Println(srvs)
		// for _, srv := range srvs {
		// 	srv.Down()
		// }
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

		fmt.Println(srvs)
		// for _, srv := range srvs {
		// 	srv.Up()
		// }
		return nil
	},
}
