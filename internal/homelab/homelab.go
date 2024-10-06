package homelab

import (
	"fmt"
	"sync"

	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/internal/utils"
	"github.com/vekio/homelab/pkg/config"
)

type Homelab struct {
	Services Services
}

// NewHomelab initializes a new Homelab instance.
func NewHomelab() (Homelab, error) {
	var services Services = make(Services, len(settings.Services))

	// Construct the directory path where the service Docker Compose files will be stored.
	if err := _dir.EnsureDir(composeFilesBasePath(), _dir.DefaultDirPerms); err != nil {
		return Homelab{}, err
	}

	// Download (or updates) the .env-example file from the repository.
	urls, err := utils.GenerateGithubURL(settings.Repository.URL, settings.Repository.Branch, ".env-example")
	if err != nil {
		return Homelab{}, err
	}
	if err := utils.DownloadFile(urls, config.DirPath()); err != nil {
		return Homelab{}, err
	}

	// Check if the .env file exists.
	exits, err := _file.FileExists(envFilePath())
	if err != nil {
		return Homelab{}, err
	}

	// If the .env file does not exist, generate .evn from .env-example.
	if !exits {
		envExamplePath := fmt.Sprintf("%s/.env-example", config.DirPath())
		if err := _file.MoveFile(envExamplePath, envFilePath()); err != nil {
			return Homelab{}, err
		}
	}

	// Return the initialized Homelab with all the services set up.
	homelab := Homelab{
		Services: services,
	}

	// Create all services as defined in the settings.
	if err := homelab.createServices(); err != nil {
		return Homelab{}, nil
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

// createServices initializes the services defined in the settings and adds them to the Homelab Services map.
func (h *Homelab) createServices() error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(settings.Services)) // Channel to capture any errors.

	// Iterate over all services defined in the settings.
	for serviceName, serviceConfig := range settings.Services {
		wg.Add(1)
		go func(name, context string, composeFiles, extraFiles []string) {
			defer wg.Done()

			// Create a new service using the provided name, context, and compose file list.
			service, err := NewService(name, context, composeFiles, extraFiles)
			if err != nil {
				// If an error occurs, send it to the error channel.
				errCh <- err
				return
			}

			// Add the service to the Homelab's Services map.
			h.Services[name] = service
		}(serviceName, serviceConfig.Context, serviceConfig.ComposeFiles, serviceConfig.ExtraFiles)
	}

	// Wait for all goroutines to complete.
	wg.Wait()
	close(errCh) // Close the error channel once all services are processed.

	// Check if there were any errors during the service creation process.
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
