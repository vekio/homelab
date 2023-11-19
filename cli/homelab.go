package homelab

import (
	"log"
	"log/slog"
	"os"

	config "github.com/vekio/homelab/cli/conf"
)

func init() {

	env, err := config.GetCurrentEnv()
	if err != nil {
		log.Fatal(err)
	}

	switch env {
	case config.DEV:
		opts := &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}
		handler := slog.NewTextHandler(os.Stdout, opts)
		slog.SetDefault(slog.New(handler))
	case config.PRO:
		opts := &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}
		handler := slog.NewTextHandler(os.Stdout, opts)
		slog.SetDefault(slog.New(handler))
	}
}
