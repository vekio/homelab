package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vekio/homelab/internal/cmd/homelab"
	"github.com/vekio/homelab/internal/config"
)

func main() {
	// Read, load and validate configuration.
	conf, err := readConfig()
	if err != nil {
		log.Fatalf("error config homelab: %v", err)
	}

	// Homelab root command.
	rootCmd := homelab.NewCmdHomelab(conf)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error homelab: %v", err)
	}
}

func readConfig() (*config.ConfigManager[config.Config], error) {
	configFile := "config.yml" // Default configuration file
	env := os.Getenv("HOMELAB_ENV")
	switch env {
	case "develop":
		configFile = "config.dev.yml"
	case "production":
		// configFile stays default "config.yml"
	case "":
		// Handle empty env as default or error
	default:
		return nil, fmt.Errorf("unknown environment setting %s", env)
	}

	// Create a new configManager instance for homelab with the specific configuration file.
	conf, err := config.NewConfigManager[config.Config]("homelab", configFile)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
