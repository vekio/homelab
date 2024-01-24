package homelab

import (
	"fmt"
	"log"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/pkg/config"
	"github.com/vekio/homelab/internal/pkg/secrets"
	acmd "github.com/vekio/homelab/pkg/conf"
)

var Version string

var Cmd = &cli.App{
	Name:        command,
	Usage:       `manage homelab`,
	Description: `Manage homelab services and their config files.`,
	UsageText:   fmt.Sprintf("%s COMMAND SUBCOMMAND", command),
	Authors: []*cli.Author{
		{
			Name:  authorName,
			Email: authorEmail,
		},
	},
	Version:   Version,
	Compiled:  time.Now(),
	Copyright: fmt.Sprintf("%s (%s) Copyright %s\nLicense Apache-2.0", command, Version, authorName),
	Commands:  commands(),
}

func commands() []*cli.Command {
	var cmds []*cli.Command

	// Services commands
	cmds = append(cmds, serviceCmds()...)

	// Init command
	cmds = append(cmds, initCmd)

	// AllUp command
	cmds = append(cmds, allUpCmd)

	// AllDown command
	cmds = append(cmds, allDownCmd)

	return cmds
}

func init() {

	// Init cmd config
	err := acmd.Config.SoftInit()
	if err != nil {
		log.Fatal(err)
	}

	err = acmd.Config.ReadConfig(&config.Settings)
	if err != nil {
		log.Fatal(err)
	}

	if err = config.Settings.LoadEnvVars(); err != nil {
		log.Fatal(err)
	}

	// Config logger
	// TODO custom logger https://betterstack.com/community/guides/logging/logging-in-go/#creating-custom-handlers
	// opts := &slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// }
	// logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	// slog.SetDefault(logger)

	// Init secrets
	secretsFile := acmd.Config.DirPath() + "/secrets.yml"
	if err := secrets.SoftInitSecrets(secretsFile); err != nil {
		log.Fatalf("failed to init secrets %s", err)
	}

}
