package config

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

type C struct {
	Service Service `yaml:"service"`
	Context Context `yaml:"context"`
}
