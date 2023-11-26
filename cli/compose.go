package homelab

import (
	"fmt"
	"log/slog"
	"os/user"

	config "github.com/vekio/homelab/cli/conf"
)

func getHomeDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", fmt.Errorf("Error getting current user: %s\n", err)
	}
	return currentUser.HomeDir, nil
}

func ComposeFile(service string) (string, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}
	slog.Debug("load compose file", "service", service)

	composeFilePath := fmt.Sprintf("%s/src/homelab/services/%s/compose.yml", homeDir, service)

	return composeFilePath, nil
}

func ComposeEnvFile(service string) (string, error) {
	homeDir, err := getHomeDir()
	if err != nil {
		return "", err
	}

	envFile, err := config.GetCurrentEnvFile()
	if err != nil {
		return "", err
	}
	slog.Debug("load env file", "env", envFile)

	envFilePath := fmt.Sprintf("%s/src/homelab/%s", homeDir, envFile)

	return envFilePath, nil
}
