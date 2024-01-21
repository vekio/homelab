package config

import (
	"fmt"

	"github.com/joho/godotenv"
)

var Settings homeConfig

// TODO validar si Settings es correcto tras cargar el fichero
func (s homeConfig) isValid() (bool, error) {
	return false, nil
}

func (s homeConfig) LoadEnvVars() error {
	for _, ctx := range s.Context.Available {
		if ctx.Name == s.Context.Current {
			if err := godotenv.Load(ctx.EnvFile); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("error loading environment: %s doesn't exists", s.Context.Current)
}

func (s homeConfig) ServiceRepo() string {
	return s.Service.Repository
}
