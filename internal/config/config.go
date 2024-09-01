package config

type Config struct {
	Services   map[string]ServiceConfig `yaml:",flow"`
	Repository RepositoryConfig         `yaml:"repo"`
	Contexts   []string                 `yaml:"contexts"`
}

type RepositoryConfig struct {
	URL    string
	Branch string
}

type ServiceConfig struct {
	Server string `yaml:"context"`
}
