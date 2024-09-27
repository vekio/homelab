package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"sync"

	_file "github.com/vekio/fs/file"
)

// DownloadFiles retrieves files from the provided URLs and saves them to the specified destination directory.
// It downloads each file from the URLs list and stores them in the target folder with the same base name as in the URL.
func DownloadFiles(urls []string, destination string) error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(urls)) // Channel to handle any errors that occur concurrently.

	// Iterate over all URLs and start downloading their corresponding files.
	for _, url := range urls {
		wg.Add(1)
		go func(url string) { // Start a goroutine for each download
			defer wg.Done()

			// Send a GET request to the URL.
			resp, err := http.Get(url)
			if err != nil {
				errCh <- fmt.Errorf("error fetching %s: %w", url, err)
				return
			}
			defer resp.Body.Close()

			// Read the response body.
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				errCh <- fmt.Errorf("error reading response from %s: %w", url, err)
				return
			}

			// Check if the HTTP status code indicates success.
			if resp.StatusCode != http.StatusOK {
				errCh <- fmt.Errorf("HTTP error %d for %s: %s", resp.StatusCode, url, body)
				return
			}

			// Generate the file path where the content will be saved.
			filePath := fmt.Sprintf("%s/%s", destination, path.Base(url))

			// Write the downloaded content to a file at the specified location.
			if err := os.WriteFile(filePath, body, _file.DefaultFilePerms); err != nil {
				errCh <- fmt.Errorf("failed to write file %s: %w", filePath, err)
				return
			}
		}(url)
	}

	// Wait for all goroutines to complete.
	wg.Wait()
	close(errCh) // Close the error channel once all tasks have completed.

	// Check for any errors that may have occurred during the execution.
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}
