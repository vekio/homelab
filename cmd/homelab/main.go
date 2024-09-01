package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/pkg/cmd/homelab"
)

func main() {
	// Read, load and validate configuration.
	conf, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Read and load configuration.
	commands := homelab.NewCmdHomelab(conf)

	// Homelab app.
	app := &cli.App{
		Name:     "homelab",
		Usage:    "Manage homelab services",
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func readConfig() (*config.ConfigManager, error) {
	configFile := "config.yml" // Default configuration file
	if env := os.Getenv("HOMELAB_ENV"); env == "develop" {
		configFile = "config.dev.yml"
	} else if env != "" && env != "production" {
		return nil, fmt.Errorf("unknown environment setting %s", env)
	}

	// Create a new configManager instance for homelab with the specific configuration file.
	conf := config.NewConfigManager("homelab", configFile)

	// Validate configuration.
	if err := validate(conf); err != nil {
		return nil, fmt.Errorf("configuration scheman invalid: %w", err)
	}

	// Initialize the configuration file.
	if err := conf.SoftInit(); err != nil {
		return nil, fmt.Errorf("failed to initialize configuration: %w", err)
	}

	return conf, nil
}

func validate(conf *config.ConfigManager) error {
	config, err := conf.Data()
	if err != nil {
		return err
	}

	// Validate contexts.
	for _, context := range config.Contexts {
		if context == "local" {
			context = "default"
		}
		cmd := exec.Command("docker", "context", "inspect", context)
		output, err := cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "not found") {
				fmt.Fprintf(os.Stdout, "Warning: context '%s' not found ⚠️\n", context)
			} else {
				return fmt.Errorf("error inspecting docker context '%s': %w", context, err)
			}
		}
	}
	return nil
}
