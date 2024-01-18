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

const version = "v0.0.1"
const command = "homelab"
const authorEmail = "alberto@casta.me"
const authorName = "Alberto Castañeiras"
const summary = "CLI for manage homelab"
const description = "Manage homelab services and their config files."

var Cmd = &cli.App{
	Name:        command,
	Usage:       summary,
	Description: description,
	UsageText:   fmt.Sprintf("%s COMMAND SUBCOMMAND", command),
	// HideHelpCommand: true,
	Authors: []*cli.Author{
		{
			Name:  authorName,
			Email: authorEmail,
		},
	},
	Version:   version,
	Compiled:  time.Now(),
	Copyright: fmt.Sprintf("%s (%s) Copyright %s\nLicense Apache-2.0", command, version, authorName),
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
