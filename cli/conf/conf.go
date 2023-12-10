package conf

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"log"

	"gopkg.in/yaml.v3"
)

//go:embed config.yml
var defaultConfig string

type Environment string

type ContextProp struct {
	Name        string      `yaml:"name"`
	Environment Environment `yaml:"environment"`
	EnvFile     string      `yaml:"env_file"`
}

type Context struct {
	Current   string        `yaml:"current"`
	Available []ContextProp `yaml:",flow"`
}

type Config struct {
	Context Context `yaml:"context"`
}

const (
	DEV Environment = "dev"
	PRO Environment = "pro"
)

var Settings Config

func init() {

	// Get user default config folder
	dir, _ := os.UserConfigDir()

	// Check command config in user config
	exePath, _ := os.Executable()
	exePath, _ = filepath.EvalSymlinks(exePath)
	exeName := strings.TrimSuffix(
		filepath.Base(exePath), filepath.Ext(exePath),
	)

	// Check config directory
	configDir := fmt.Sprintf("%s/%s", dir, exeName)
	_, err := os.Stat(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(configDir, 0755)
			if err != nil {
				log.Fatalf("Error creating the config directory: %v", err)
			}
		}
	}

	// Check config file
	configFile := fmt.Sprintf("%s/config.yml", configDir)
	_, err = os.Stat(configFile)
	if err != nil {
		err = os.WriteFile(configFile, []byte(defaultConfig), 0644)
		if err != nil {
			log.Fatalf("Error creating the config file: %v", err)
		}
	}

	yamlFile, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	err = yaml.Unmarshal(yamlFile, &Settings)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	if !Settings.isValid() {
		log.Fatalf("Error settings in config.yml not valid")
	}
}

func (c *Config) isValid() bool {
	return c.Context.Current != "" && len(c.Context.Available) > 0
}

func GetCurrentEnvFile() (string, error) {
	for _, context := range Settings.Context.Available {
		if context.Name == Settings.Context.Current {
			return context.EnvFile, nil
		}
	}
	return "", fmt.Errorf("Error current context doesn't exists: %s", Settings.Context.Current)
}

func GetCurrentEnv() (Environment, error) {
	for _, context := range Settings.Context.Available {
		if context.Name == Settings.Context.Current {
			return context.Environment, nil
		}
	}
	return "", fmt.Errorf("Error current context doesn't exists: %s", Settings.Context.Current)
}
