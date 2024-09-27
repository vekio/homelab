package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// GenerateGithubURLs generates raw content URLs for the given repository and branch on GitHub.
// It takes the base repository URL, branch, and a list of file paths to create the appropriate raw content URLs.
func GenerateGithubURLs(repoURL, branch string, filePaths []string) ([]string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return nil, err
	}

	// Extract the username and repository name from the URL.
	segments := strings.Split(parsedURL.String(), "/")
	if len(segments) < 5 {
		return nil, fmt.Errorf("invalid GitHub URL format")
	}
	username := segments[3]
	repository := segments[4]

	// Base URL for raw content from GitHub.
	githubRawContentURL := "https://raw.githubusercontent.com"

	// Generate URLs for each file.
	var urls []string
	for _, file := range filePaths {
		url := fmt.Sprintf("%s/%s/%s/%s/%s", githubRawContentURL, username, repository, branch, file)
		urls = append(urls, url)
	}

	return urls, nil
}
