package homelab

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/vekio/homelab/cli/conf"
	_c "github.com/vekio/homelab/cli/config"
)

func init() {
	err := conf.Config.SoftInit()
	if err != nil {
		log.Fatal(err)
	}

	_, err = conf.Config.Data()
	if err != nil {
		log.Fatal(err)
	}

	err = conf.Config.ReadConfig(&_c.Config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(_c.Config)

	// err = yaml.Unmarshal([]byte(data), &settings)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if settings.isValid() {
	// 	env, err := settings.getCurrentEnv()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	switch env {
	// 	case DEV:
	// 		settings.loadEnvVariables(DEV)
	// 	case PRO:
	// 		settings.loadEnvVariables(PRO)
	// 	}
	// }
}

// func (s Settings) getCurrentEnv() (Environment, error) {
// 	for _, context := range s.Context.Available {
// 		if context.Name == s.Context.Current {
// 			return context.Environment, nil
// 		}
// 	}
// 	return "", fmt.Errorf("current context doesn't exists: %s", s.Context.Current)
// }

// func (s Settings) loadEnvVariables(e Environment) error {
// 	for _, context := range s.Context.Available {
// 		if context.Environment == e {
// 			if err := godotenv.Load(context.EnvFile); err != nil {
// 				return err
// 			}
// 		}
// 	}
// 	return fmt.Errorf("environment doesn't exists: %s", e)
// }

// func (s Settings) isValid() bool {
// 	return s.Context.Current != "" && len(s.Context.Available) > 0 && s.Service.Repository != ""
// }

// func (s Settings) getRepository() string {
// 	return s.Service.Repository
// }
