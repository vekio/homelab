package homelab

import (
	"fmt"
	"sync"

	_dir "github.com/vekio/fs/dir"
	"github.com/vekio/homelab/pkg/config"
)

type Homelab struct {
	Services Services
}

// NewHomelab initializes a new Homelab instance.
func NewHomelab() (Homelab, error) {
	var services Services = make(Services)

	// Construct the directory path where the service Docker Compose files will be stored.
	if err := _dir.EnsureDir(composeFilesBasePath(), _dir.DefaultDirPerms); err != nil {
		return Homelab{}, err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(settings.Services)) // Channel to capture any errors.

	// Iterate over all the services defined in the settings.
	for serviceName, serviceConfig := range settings.Services {
		wg.Add(1)
		go func(name, context string, composeFiles []string) {
			defer wg.Done()

			// Create a new service using the provided name, context, and compose file list.
			service, err := NewService(name, context, composeFiles)
			if err != nil {
				errCh <- err
				return
			}

			// Add the service to the map of services.
			services[name] = service
		}(serviceName, serviceConfig.Context, serviceConfig.ComposeFiles)
	}

	wg.Wait()    // Wait for all goroutines to complete.
	close(errCh) // Close the error channel once all services are processed.

	// Check if there were any errors during the service creation process.
	for err := range errCh {
		if err != nil {
			return Homelab{}, err
		}
	}

	// Return the initialized Homelab with all the services set up.
	homelab := Homelab{
		Services: services,
	}

	return homelab, nil
}

// ServicesNames returns a list of all service names available in the Homelab.
func (h *Homelab) ServicesNames() []string {
	keys := make([]string, 0, len(h.Services))
	// Collect all the keys (service names) from the services map.
	for k := range h.Services {
		keys = append(keys, k)
	}
	return keys
}

// ServiceByName retrieves a service by its name from the Homelab.
// Returns an error if the service is not found.
func (h *Homelab) ServiceByName(name string) (*Service, error) {
	// Check if the service exists in the map.
	service, ok := h.Services[name]
	if !ok {
		// If the service is not found, return an error.
		return nil, fmt.Errorf("service %s not found", name)
	}
	// Return the found service.
	return service, nil
}

// composeFilesBasePath returns the base directory path where Docker Compose files are stored.
// This path is based on the repository's branch name.
func composeFilesBasePath() string {
	return fmt.Sprintf("%s/dockercomposefiles-%s", config.DirPath(), settings.Repository.Branch)
}
