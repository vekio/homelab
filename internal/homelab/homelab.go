package homelab

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/pkg/secrets"
	"github.com/vekio/homelab/pkg/conf"
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
	// Commands:  commands(),
}

func init() {

	// Config logger
	// TODO custom logger https://betterstack.com/community/guides/logging/logging-in-go/#creating-custom-handlers
	opts := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	filename := conf.Config.DirPath() + "/secrets.yml"
	if err := secrets.SoftInitSecrets(filename); err != nil {
		log.Fatalf("failed to init secrets %s", err)
	}

}
