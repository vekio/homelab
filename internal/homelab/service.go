package homelab

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/pkg/services"
)

func serviceCmds() []*cli.Command {
	var cmds []*cli.Command

	srvs := services.Available()

	for _, srv := range srvs {
		composeCmds := []*cli.Command{
			configCmd,
			downCmd,
			logsCmd,
			pullCmd,
			restartCmd,
			stopCmd,
			upCmd,
			upgradeCmd,
		}

		srvCmd := &cli.Command{
			Name:        srv.Name,
			Usage:       fmt.Sprintf("Manage %s service", srv.Name),
			Subcommands: composeCmds,
		}

		cmds = append(cmds, srvCmd)
	}

	return cmds
}

// import (
// 	"fmt"

// 	"github.com/urfave/cli/v2"
// 	services "github.com/vekio/homelab/cli/internal/pkg/services_old"
// )

// func serviceCmdFactory(service string) *cli.Command {
// 	defaultCmds := []*cli.Command{
// 		configCmd,
// 		pullCmd,
// 		upCmd,
// 		logsCmd,
// 		stopCmd,
// 		downCmd,
// 		upgradeCmd,
// 	}

// 	if service == services.PROTONMAIL_BRIDGE {
// 		defaultCmds = append(defaultCmds, initSmtp)
// 	}

// 	return &cli.Command{
// 		Name:        service,
// 		Usage:       fmt.Sprintf("Manage %s service", service),
// 		Subcommands: defaultCmds,
// 	}
// }

// var initSmtp = &cli.Command{
// 	Name:    "init",
// 	Aliases: []string{"i"},
// 	Usage:   "Initialize required folders and config files",
// 	Action: func(ctx *cli.Context) error {
// 		if err := services.InitProtonmailBridge(); err != nil {
// 			return err
// 		}
// 		if err := execDockerCompose(services.PROTONMAIL_BRIDGE, "run", "--rm", "protonmail-bridge"); err != nil {
// 			return err
// 		}
// 		return nil
// 	},
// }
