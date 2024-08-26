package secrets

import (
	"gopkg.in/yaml.v3"
)

func New() (HomelabSecrets, error) {
	var secrets HomelabSecrets

	data, err := Manager.Data()
	if err != nil {
		return HomelabSecrets{}, err
	}

	if err := yaml.Unmarshal([]byte(data), &secrets); err != nil {
		return HomelabSecrets{}, err
	}

	return secrets, nil
}
