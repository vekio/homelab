package homelab

import (
	"fmt"
	"os"
	"text/template"

	"github.com/urfave/cli/v2"
	_dir "github.com/vekio/fs/dir"
	"github.com/vekio/homelab/cli/conf"
)

var initCmd = &cli.Command{
	Name:    "init",
	Aliases: []string{"i"},
	Usage:   "Initialize required folders and config files",
	Action: func(cCtx *cli.Context) (err error) {
		service := getService(cCtx)

		switch service {
		case TRAEFIK:
			err = initTraefik(service)
			return err
		}

		return
	},
}

func initTraefik(service string) error {
	serviceRepo := settings.getRepository()
	traefikConfig := fmt.Sprintf("%s/%s/config", serviceRepo, service)

	// Copy config folder
	localConfig := fmt.Sprintf("%s/%s", conf.Config.DirPath(), service)
	err := _dir.Copy(traefikConfig, localConfig)
	if err != nil {
		return err
	}

	// Parse traefik.yml
	treafikYMLFile := fmt.Sprintf("%s/traefik.yml", localConfig)
	data := map[string]string{
		"DOMAIN":             os.Getenv("DOMAIN"),
		"TRAEFIK_CERT_EMAIL": os.Getenv("TRAEFIK_CERT_EMAIL"),
	}
	if err := parseConfigFile(treafikYMLFile, data); err != nil {
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

	outputFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	// Execute the template and write to the output file
	if err := tmpl.Execute(outputFile, data); err != nil {
		return err
	}

	return nil
}
