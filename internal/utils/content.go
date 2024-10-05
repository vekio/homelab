package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// GenerateGithubURL generates a raw content URL for the given repository and branch on GitHub.
// It takes the base repository URL, branch, and a file path to create the appropriate raw content URL.
func GenerateGithubURL(repoURL, branch, filePath string) (string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", err
	}

	// Extract the username and repository name from the URL.
	segments := strings.Split(parsedURL.String(), "/")
	if len(segments) < 5 {
		return "", fmt.Errorf("invalid GitHub URL format")
	}
	username := segments[3]
	repository := segments[4]

	// Base URL for raw content from GitHub.
	githubRawContentURL := "https://raw.githubusercontent.com"

	// Generate URL.
	return fmt.Sprintf("%s/%s/%s/%s/%s", githubRawContentURL, username, repository, branch, filePath), nil
}
