package homelab

import (
	"log"
	"log/slog"
	"os"

	"github.com/vekio/homelab/cli/internal/pkg/secrets"
	"github.com/vekio/homelab/cli/pkg/conf"
)

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
