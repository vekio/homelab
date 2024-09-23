package homelab

import (
	"fmt"
	"os/exec"
	"slices"
	"strings"

	"github.com/vekio/homelab/pkg/config"
)

type Config struct {
	Services map[string]struct {
		Context      string   `yaml:"context"`
		ComposeFiles []string `yaml:"compose_files"`
	} `yaml:",flow"`
	Repository struct {
		URL    string `yaml:"url"`
		Branch string `yaml:"branch"`
	} `yaml:"repo"`
	Contexts []string `yaml:"contexts"`
}

var conf Config

func init() {
	config.Load(&conf)
}

func (c Config) Validate() error {
	// Validate contexts.
	for _, context := range c.Contexts {
		if context == "local" {
			context = "default"
		}
		cmd := exec.Command("docker", "context", "inspect", context)
		output, err := cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "not found") {
				fmt.Printf("Warning: context '%s' not found ⚠️\n", context)
			} else {
				return fmt.Errorf("failed checking docker context '%s': %w", context, err)
			}
		}
	}

	// Validate services context.
	for service, srv := range c.Services {
		if !slices.Contains(c.Contexts, srv.Context) {
			fmt.Printf("Warning: unkown context '%s' in '%s' ⚠️\n", srv.Context, service)
		}
	}
	return nil
}
