package homelab

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/utils"
)

type Services = map[string]*Service

type Service struct {
	name         string
	context      string
	composeFiles []string
}

func NewService(name, context string, composeFiles []string, conf *config.ConfigManager[config.Config]) (*Service, error) {
	service := &Service{
		name:         name,
		context:      context,
		composeFiles: composeFiles,
	}

	if err := downloadComposeFiles(service, conf); err != nil {
		return nil, err
	}
	return nil, nil
}

// downloadComposeFiles downloads docker compose files for the service.
// It ensures that all compose files are fetched and stored in a specific directory based on the repository branch.
func downloadComposeFiles(service *Service, conf *config.ConfigManager[config.Config]) error {
	// Construct the directory path where the compose files will be stored.
	dirPath := conf.DirPath() + "/dockercompose-" + conf.Data.Repository.Branch
	if err := _dir.EnsureDir(dirPath, _dir.DefaultDirPerms); err != nil {
		return err
	}

	// Ensure directory for the service.
	servicePath := dirPath + "/" + service.name
	if err := _dir.EnsureDir(servicePath, _dir.DefaultDirPerms); err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(service.composeFiles)) // Error channel for concurrent error handling.

	// Iterate over all service names and download their respective compose files.
	for _, composeFile := range service.composeFiles {
		wg.Add(1)
		go func(composeFile string) { // Start the goroutine
			defer wg.Done()

			// Build URL for the compose file.
			url, err := utils.BuildComposeFileURL(conf.Data.Repository.URL, conf.Data.Repository.Branch, service.name, composeFile)
			if err != nil {
				errCh <- err
			}

			resp, err := http.Get(url)
			if err != nil {
				errCh <- fmt.Errorf("error fetching %s: %w", url, err)
				return
			}
			defer resp.Body.Close()

			// Read the body of the response.
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errCh <- fmt.Errorf("error reading response from %s: %w", url, err)
				return
			}

			// Check the HTTP status code.
			if resp.StatusCode != 200 {
				errCh <- fmt.Errorf("HTTP error %d for %s: %s", resp.StatusCode, url, body)
				return
			}

			// Save file content.
			filePath := fmt.Sprintf("%s/%s", servicePath, composeFile)
			if err := os.WriteFile(filePath, body, _file.DefaultFilePerms); err != nil {
				errCh <- fmt.Errorf("failed to write file %s: %w", filePath, err)
				return
			}
		}(composeFile)
	}

	wg.Wait()    // Wait for all goroutines to finish.
	close(errCh) // Close the channel after all goroutines report they are done.

	// Check for errors from the error channel.
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
