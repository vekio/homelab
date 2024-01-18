package homelab

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
