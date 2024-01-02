package services

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

const (
	AUTHELIA string = "authelia"
	GITEA    string = "gitea"
	IMMICH   string = "immich"
	LLDAP    string = "lldap"
	TRAEFIK  string = "traefik"
)

func AvailableServices() []string {
	return []string{AUTHELIA, GITEA, IMMICH, LLDAP, TRAEFIK}
}

func parseConfigFile(filename string, data interface{}) error {
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(string(fileContent))
	if err != nil {
		return err
	}

	outputFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("executing template for %s: %w", filepath.Base(filename), err)
	}

	return nil
}
