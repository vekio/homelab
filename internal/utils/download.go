package utils

import (
	"fmt"
	"io"
	"net/http"
	"path"

	_file "github.com/vekio/fs/file"
)

// DownloadFile downloads a file from the specified URL and saves it to the given destination directory.
// The file is stored in the target directory using the same base name as the one in the URL.
func DownloadFile(url, destination string) error {
	// Send a GET request to the URL.
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching %s: %w", url, err)
	}
	defer resp.Body.Close()

	// Read the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response from %s: %w", url, err)
	}

	// Check if the HTTP status code indicates success.
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error %d for %s: %s", resp.StatusCode, url, body)
	}

	// Generate the file path where the content will be saved.
	filePath := fmt.Sprintf("%s/%s", destination, path.Base(url))

	// Use the WriteFileContent function from _file to write the content to the file.
	if err := _file.WriteFileContent(filePath, body, _file.DefaultFilePerms); err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}
