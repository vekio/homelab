package conf

import (
	"fmt"
)

type Environment string

type contextConfig struct {
	name        string // use hostname as context name
	environment Environment
	envFile     string
}

type settings struct {
	currentContext string
	contexts       []contextConfig
}

const (
	DEV Environment = "dev"
	PRO Environment = "pro"
)

var Settings settings

func init() {

	var contexts []contextConfig

	storm := contextConfig{
		name:        "storm",
		environment: DEV,
		envFile:     ".env-dev",
	}

	spring := contextConfig{
		name:        "spring",
		environment: PRO,
		envFile:     ".env-pro",
	}

	contexts = append(contexts, storm, spring)

	Settings = settings{
		currentContext: "storm", // use storm, dev environment, as default
		contexts:       contexts,
	}
}

func GetCurrentEnvFile() (string, error) {
	for _, context := range Settings.contexts {
		if context.name == Settings.currentContext {
			return context.envFile, nil
		}
	}
	return "", fmt.Errorf("Error finding current env file.")
}

func GetCurrentEnv() (Environment, error) {
	for _, context := range Settings.contexts {
		if context.name == Settings.currentContext {
			return context.environment, nil
		}
	}
	return "", fmt.Errorf("Error finding current environment.")
}
