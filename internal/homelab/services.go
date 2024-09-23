package homelab

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"sync"

	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/pkg/config"
)

type Services = map[string]*Service

type Service struct {
	Name         string
	Context      string
	ComposeFiles []string
}

func NewService(name, context string, composeFiles []string) (*Service, error) {
	service := &Service{
		Name:         name,
		Context:      context,
		ComposeFiles: composeFiles,
	}

	if err := downloadComposeFiles(service); err != nil {
		return nil, err
	}
	return service, nil
}

func (s Service) ComposeFilesGithubURLs(repoURL, branch string) ([]string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return nil, err
	}
	segments := strings.Split(parsedURL.String(), "/")
	var githubRawContentURL = "https://raw.githubusercontent.com"
	var username = segments[3]
	var repository = segments[4]

	var urls []string
	for _, composeFile := range s.ComposeFiles {
		url := fmt.Sprintf("%s/%s/%s/%s/%s/%s", githubRawContentURL, username, repository, branch, s.Name, composeFile)
		urls = append(urls, url)
	}
	return urls, nil
}

// downloadComposeFiles downloads docker compose files for the service.
// It ensures that all compose files are fetched and stored in a specific directory based on the repository branch.
func downloadComposeFiles(service *Service) error {
	// Construct the directory path where the compose files will be stored.
	dirPath := config.DirPath() + "/dockercompose-" + conf.Repository.Branch
	if err := _dir.EnsureDir(dirPath, _dir.DefaultDirPerms); err != nil {
		return err
	}

	// Ensure directory for the service.
	servicePath := dirPath + "/" + service.Name
	if err := _dir.EnsureDir(servicePath, _dir.DefaultDirPerms); err != nil {
		return err
	}

	urls, err := service.ComposeFilesGithubURLs(conf.Repository.URL, conf.Repository.Branch)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(service.ComposeFiles)) // Error channel for concurrent error handling.

	// Iterate over all service names and download their respective compose files.
	for _, url := range urls {
		wg.Add(1)
		go func(url string) { // Start the goroutine
			defer wg.Done()

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
			filePath := fmt.Sprintf("%s/%s", servicePath, path.Base(url))
			if err := os.WriteFile(filePath, body, _file.DefaultFilePerms); err != nil {
				errCh <- fmt.Errorf("failed to write file %s: %w", filePath, err)
				return
			}
		}(url)
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
