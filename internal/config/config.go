package config

import (
	Z "github.com/rwxrob/bonzai/z"
	"gopkg.in/yaml.v3"
)

func New() (HomelabConfig, error) {
	var config HomelabConfig

	data, err := Z.Conf.Data()
	if err != nil {
		return HomelabConfig{}, err
	}

	if err := yaml.Unmarshal([]byte(data), &config); err != nil {
		return HomelabConfig{}, err
	}

	return config, nil
}
