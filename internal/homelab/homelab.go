package homelab

import (
	"fmt"
	"os"
	"sync"

	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/utils"
)

type Homelab struct {
	Services Services
}

type Services map[string]Service

type Service struct {
	context string
}

func NewHomelabApp(conf *config.ConfigManager[config.Config]) error {
	services := make(Services)
	for srvName, srv := range conf.Data.Services {
		service := Service{
			context: srv.Context,
		}
		services[srvName] = service
	}

	if err := downloadServiceComposeFiles(conf, services); err != nil {
		return err
	}

	return nil
}

// downloadServiceComposeFiles downloads Docker Compose files for the specified services.
// It ensures that all compose files are fetched and stored in a specific directory based on the repository branch.
func downloadServiceComposeFiles(conf *config.ConfigManager[config.Config], services Services) error {
	// Construct the directory path where the compose files will be stored.
	dirPath := conf.DirPath() + "/dockercompose-" + conf.Data.Repository.Branch
	if err := _dir.EnsureDir(dirPath, _dir.DefaultDirPerms); err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(services)) // Error channel for concurrent error handling.

	// Iterate over all service names and download their respective compose files.
	for srvName := range services {
		wg.Add(1)
		go func(srvName, dirPath string) { // Start the goroutine
			defer wg.Done()

			// Build URLs for compose files.
			urls, err := utils.BuildComposeFileURLs(conf.Data.Repository.URL, conf.Data.Repository.Branch, srvName)
			if err != nil {
				errCh <- err
				return
			}

			// Fetch compose files.
			results, err := utils.FetchComposeFiles(urls)
			if err != nil {
				errCh <- err
				return
			}

			// Ensure directory for the specific service.
			servicePath := dirPath + "/" + srvName
			if err := _dir.EnsureDir(servicePath, _dir.DefaultDirPerms); err != nil {
				errCh <- err
				return
			}

			// Save each fetched file.
			for _, result := range results {
				if result.Err != nil {
					errCh <- result.Err
					return
				}
				filePath := fmt.Sprintf("%s/%s", servicePath, result.FileName)
				fmt.Println(filePath)
				if err := os.WriteFile(filePath, []byte(result.Content), _file.DefaultFilePerms); err != nil {
					errCh <- fmt.Errorf("failed to write file %s: %w", filePath, err)
					return
				}
			}

		}(srvName, dirPath)
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
