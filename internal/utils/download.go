package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"sync"
)

type Result struct {
	FileName string
	Content  string
	Err      error
}

// BuildComposeFileURLs constructs URLs for docker compose files based on the repository URL, branch, and service name.
// It returns a slice of URLs pointing to the raw content of compose files on GitHub.
// Check https://github.com/vekio/dockercomposefiles struct.
func BuildComposeFileURLs(repoURL, branch, service string) ([]string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return nil, err
	}
	segments := strings.Split(parsedURL.String(), "/")
	var urlRawContent = "https://raw.githubusercontent.com"
	var username = segments[3]
	var repository = segments[4]

	return []string{
		fmt.Sprintf("%s/%s/%s/%s/%s/compose.yml", urlRawContent, username, repository, branch, service),
		fmt.Sprintf("%s/%s/%s/%s/%s/compose.traefik.yml", urlRawContent, username, repository, branch, service),
	}, nil
}

// FetchComposeFiles concurrently downloads all files from the given URLs and returns a slice of Results.
// Each result either contains the content of the file or an error if the download failed.
func FetchComposeFiles(urls []string) ([]Result, error) {
	var wg sync.WaitGroup
	ch := make(chan Result, len(urls)) // Buffered channel to hold the results.

	for _, url := range urls {
		wg.Add(1)
		go FetchURLContent(url, &wg, ch) // Start the goroutine to fetch each URL.
	}

	wg.Wait() // Wait for all goroutines to finish.
	close(ch) // Close the channel after all goroutines report they are done.

	var results []Result
	for content := range ch { // Collect all results from the channel.
		results = append(results, content)
	}

	return results, nil
}

// FetchURLContent fetches content from a specified URL and sends a Result to the provided channel.
// It handles HTTP errors by sending an error result if the status code is not 200.
func FetchURLContent(contentURL string, wg *sync.WaitGroup, ch chan<- Result) {
	defer wg.Done() // Ensure that the wait group counter decreases on function exit.

	resp, err := http.Get(contentURL)
	if err != nil {
		ch <- Result{Err: fmt.Errorf("error fetching %s: %w", contentURL, err)}
		return
	}
	defer resp.Body.Close() // Ensure that the response body is closed on function exit.

	// Read the body of the response.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{Err: fmt.Errorf("error reading response from %s: %w", contentURL, err)}
		return
	}

	// Check the HTTP status code.
	if resp.StatusCode != 200 {
		ch <- Result{Err: fmt.Errorf("HTTP error %d for %s: %s", resp.StatusCode, contentURL, body)}
		return
	}

	// Use the path package to extract the last element of the path
	parsedURL, _ := url.Parse(contentURL)
	filename := path.Base(parsedURL.Path)

	// If everything is fine, send the content back over the channel.
	ch <- Result{Content: string(body), FileName: filename}
}
