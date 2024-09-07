package utils

import (
	"fmt"

	"github.com/vekio/homelab/internal/config"
)

// ValidateService checks if the service name is valid within the given configuration.
func ValidateService(conf config.Config, service string) error {
	if _, ok := conf.Services[service]; !ok {
		return fmt.Errorf("invalid argument: %s", service)
	}
	return nil
}
