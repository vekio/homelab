package config

type environment string

type contextProp struct {
	Name        string      `yaml:"name"`
	Environment environment `yaml:"environment"`
	EnvFile     string      `yaml:"env_file"`
}

type context struct {
	Current   string        `yaml:"current"`
	Available []contextProp `yaml:",flow"`
}

type service struct {
	Repository string `yaml:"repository"`
}

type homeConfig struct {
	Service service `yaml:"service"`
	Context context `yaml:"context"`
}
