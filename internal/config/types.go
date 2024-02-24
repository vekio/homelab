package config

type configManager struct {
	name string
	dir  string
	file string
}

type serviceConfig struct {
	Context  string `yaml:"context"`
	Priority int    `yaml:"priority"`
	AllUp    bool   `yaml:"allup"`
	AllDown  bool   `yaml:"alldown"`
	AllInit  bool   `yaml:"allinit"`
}

type HomelabConfig struct {
	Repository string                   `yaml:"repository"`
	EnvFile    string                   `yaml:"env_file"`
	Services   map[string]serviceConfig `yaml:",flow"`
}
