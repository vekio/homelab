package homelab

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
	_dir "github.com/vekio/fs/dir"

	config "github.com/vekio/homelab/cli/conf"
)

func init() {

	envFile, err := config.GetCurrentEnvFile()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
	}

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

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize required folders and config files",
	Action: func(cCtx *cli.Context) (err error) {
		service := getService(cCtx)

		switch service {
		case TRAEFIK:
			err = initTraefik(service)
			return err
		}

		return
	},
}

func initTraefik(service string) (err error) {
	serviceRepo := config.GetServiceRepo()
	traefikConfig := fmt.Sprintf("%s/%s/config", serviceRepo, service)

	// Copy config folder
	localConfig := fmt.Sprintf("%s/%s", config.Settings.ConfigDir, service)
	err = _dir.Copy(traefikConfig, localConfig)
	if err != nil {
		return err
	}

	return
}
