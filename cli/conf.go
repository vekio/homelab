package homelab

import (
	_ "embed"
	"log"

	"github.com/vekio/homelab/cli/conf"
	"gopkg.in/yaml.v3"
)

type Environment string

type ContextProp struct {
	Name        string      `yaml:"name"`
	Environment Environment `yaml:"environment"`
	EnvFile     string      `yaml:"env_file"`
}

type Context struct {
	Current   string        `yaml:"current"`
	Available []ContextProp `yaml:",flow"`
}

type Service struct {
	Repository string `yaml:"repository"`
}

type Settings struct {
	Service Service `yaml:"service"`
	Context Context `yaml:"context"`
}

const (
	DEV Environment = "dev"
	PRO Environment = "pro"
)

var settings Settings

func init() {
	err := conf.Config.SoftInit()
	if err != nil {
		log.Fatal(err)
	}

	data, err := conf.Config.Data()
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal([]byte(data), &settings)
	if err != nil {
		log.Fatal(err)
	}
}

// func (c *Config) isValid() bool {
// 	return c.Context.Current != "" && len(c.Context.Available) > 0 && c.Service.Repository != ""
// }

// func GetCurrentEnvFile() (string, error) {
// 	for _, context := range Settings.Context.Available {
// 		if context.Name == Settings.Context.Current {
// 			return context.EnvFile, nil
// 		}
// 	}
// 	return "", fmt.Errorf("Error current context doesn't exists: %s", Settings.Context.Current)
// }

// func GetCurrentEnv() (Environment, error) {
// 	for _, context := range Settings.Context.Available {
// 		if context.Name == Settings.Context.Current {
// 			return context.Environment, nil
// 		}
// 	}
// 	return "", fmt.Errorf("Error current context doesn't exists: %s", Settings.Context.Current)
// }

// func GetServiceRepo() string {
// 	return Settings.Service.Repository
// }
