package config

import (
	"fmt"
	"os/exec"
	"strings"
)

type Config struct {
	Services map[string]struct {
		Server string `yaml:"context"`
	} `yaml:",flow"`
	Repository struct {
		URL    string `yaml:"url"`
		Branch string `yaml:"branch"`
	} `yaml:"repo"`
	Contexts []string `yaml:"contexts"`
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
	return nil
}
