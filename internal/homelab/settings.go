package homelab

import (
	"fmt"
	"log"
	"os/exec"
	"slices"
	"strings"

	"github.com/vekio/homelab/pkg/config"
)

type Settings struct {
	Services map[string]struct {
		Context      string   `yaml:"context"`
		ComposeFiles []string `yaml:"compose_files"`
		ExtraFiles   []string `yaml:"extras"`
	} `yaml:",flow"`
	Repository struct {
		URL    string `yaml:"url"`
		Branch string `yaml:"branch"`
	} `yaml:"repo"`
	Contexts []string `yaml:"contexts"`
}

var settings Settings

func init() {
	if err := config.SoftInit(); err != nil {
		log.Fatalf("error creating configuration file: %v", err)
	}

	if err := config.Load(&settings); err != nil {
		log.Fatalf("error loading configuration: %v", err)
	}
}

func (c Settings) Validate() error {
	// Validate contexts.
	for _, context := range c.Contexts {
		if context == "local" {
			context = "default"
		}
		cmd := exec.Command("docker", "context", "inspect", context)
		output, err := cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "not found") {
				fmt.Printf("warning: context '%s' not found ⚠️\n", context)
			} else {
				return fmt.Errorf("failed checking docker context '%s': %w", context, err)
			}
		}
	}

	// Validate services context.
	for service, srv := range c.Services {
		if !slices.Contains(c.Contexts, srv.Context) {
			fmt.Printf("warning: unkown context '%s' in '%s' ⚠️\n", srv.Context, service)
		}
	}
	return nil
}

// composeFilesBasePath returns the base directory path where Docker Compose files are stored.
// This path is based on the repository's branch name.
func composeFilesBasePath() string {
	return fmt.Sprintf("%s/dockercomposefiles-%s", config.DirPath(), settings.Repository.Branch)
}

// envFilePath
func envFilePath() string {
	return fmt.Sprintf("%s/.env-%s", config.DirPath(), settings.Repository.Branch)
}
