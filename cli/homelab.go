package homelab

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli/v2"
	_fs "github.com/vekio/fs"
	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/cli/conf"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize required folders and config files",
	Action: func(cCtx *cli.Context) error {
		service := getService(cCtx)
		repository := settings.getRepository()
		localConfig := fmt.Sprintf("%s/%s", conf.Config.DirPath(), service)
		serviceConfig := fmt.Sprintf("%s/%s/config", repository, service)

		switch service {
		case TRAEFIK:
			if err := initTraefik(serviceConfig, localConfig); err != nil {
				return err
			}
		case AUTHELIA:
			if err := initAuthelia(serviceConfig, localConfig); err != nil {
				return err
			}
		case LLDAP:
			if err := initLldap(localConfig); err != nil {
				return err
			}
		}

		return nil
	},
}

func initAuthelia(serviceConfig, localConfig string) error {
	// Copy config folder
	autheliaConfig := fmt.Sprintf("%s/config", localConfig)
	if err := _dir.Copy(serviceConfig, autheliaConfig); err != nil {
		return err
	}

	// Parse configuration.yml
	configurationYMLFile := fmt.Sprintf("%s/configuration.yml", autheliaConfig)
	data := map[string]string{
		"DOMAIN":                          os.Getenv("DOMAIN"),
		"SLD":                             os.Getenv("SLD"),
		"TLD":                             os.Getenv("TLD"),
		"AUTHELIA_SESSION_SECRET":         os.Getenv("AUTHELIA_SESSION_SECRET"),
		"AUTHELIA_STORAGE_ENCRYPTION_KEY": os.Getenv("AUTHELIA_STORAGE_ENCRYPTION_KEY"),
		"AUTHELIA_AUTHENTICATION_BACKEND_LDAP_PASSWORD": os.Getenv("AUTHELIA_AUTHENTICATION_BACKEND_LDAP_PASSWORD"),
	}
	if err := parseConfigFile(configurationYMLFile, data); err != nil {
		return err
	}

	return nil
}

func initTraefik(serviceConfig, localConfig string) error {
	// Create acme.json
	acmeFile := fmt.Sprintf("%s/%s/acme.json", localConfig, "certificates")
	if err := _file.Touch(acmeFile, os.FileMode(0600)); err != nil {
		return err
	}

	// Copy config folder
	traefikConfig := fmt.Sprintf("%s/config", localConfig)
	if err := _dir.Copy(serviceConfig, traefikConfig); err != nil {
		return err
	}

	// Parse traefik.yml
	treafikYMLFile := fmt.Sprintf("%s/traefik.yml", traefikConfig)
	data := map[string]string{
		"DOMAIN":             os.Getenv("DOMAIN"),
		"TRAEFIK_CERT_EMAIL": os.Getenv("TRAEFIK_CERT_EMAIL"),
	}
	if err := parseConfigFile(treafikYMLFile, data); err != nil {
		return err
	}

	return nil
}

func initLldap(localConfig string) error {
	dataDir := fmt.Sprintf("%s/data/", localConfig)
	err := _fs.Create(dataDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}

func parseConfigFile(file string, data interface{}) error {
	fileContent, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(string(fileContent))
	if err != nil {
		return err
	}

	outputFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	if err := tmpl.Execute(outputFile, data); err != nil {
		return fmt.Errorf("executing template for %s: %w", filepath.Base(file), err)
	}

	return nil
}
