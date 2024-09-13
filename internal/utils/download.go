package utils

import (
	"fmt"
	"net/url"
	"strings"
)

// BuildComposeFileURL constructs URL for docker compose file based on the repository URL, branch, and service name.
// It returns a URL pointing to the raw content of compose files on GitHub.
// Check https://github.com/vekio/dockercomposefiles struct.
func BuildComposeFileURL(repoURL, branch, service string, composeFile string) (string, error) {
	parsedURL, err := url.Parse(repoURL)
	if err != nil {
		return "", err
	}
	segments := strings.Split(parsedURL.String(), "/")
	var urlRawContent = "https://raw.githubusercontent.com"
	var username = segments[3]
	var repository = segments[4]

	return fmt.Sprintf("%s/%s/%s/%s/%s/%s", urlRawContent, username, repository, branch, service, composeFile), nil
}
