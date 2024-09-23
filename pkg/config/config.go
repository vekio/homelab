package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// Config manages configuration files for an application.
type Config struct {
	appName string // Name of the application
	dir     string // Directory where the config files are stored
	file    string // Configuration file name
}

var config *Config

func init() {
	config = newConfig()
}

// Validatable is an interface that should be implemented by all config types
// that will be managed by Config.
type Validatable interface {
	Validate() error
}

// newConfig creates a new instance of ConfigManager for an application.
func newConfig() *Config {
	// Extract the executable name from the first argument.
	appName := filepath.Base(os.Args[0])

	// User config directory.
	dir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("failed to get user config directory: %s", err)
	}

	// Check for the app environment variable.
	envValue := os.Getenv(fmt.Sprintf("%s_ENV", strings.ToUpper(appName)))

	configFile := "config.yml"
	if envValue != "" && envValue == "develop" {
		configFile = "config.dev.yml"
	}

	cm := &Config{
		dir:     dir,
		appName: appName,
		file:    configFile,
	}

	return cm
}

// DirPath returns the path to the directory where the configuration file is stored.
func DirPath() string {
	return config.dirPath()
}

func (cm *Config) dirPath() string {
	return filepath.Join(cm.dir, cm.appName)
}

// Path returns the full path to the configuration file.
func Path() string {
	return config.path()
}

func (cm *Config) path() string {
	return filepath.Join(cm.dirPath(), cm.file)
}

// Content reads and returns the contents of the configuration file.
func Content() ([]byte, error) {
	return config.content()
}

func (cm *Config) content() ([]byte, error) {
	return os.ReadFile(cm.path())
}

func Load(data Validatable) error {
	buf, err := config.content()
	if err != nil {
		return err
	}
	// Deserialize the configuration file.
	if err := yaml.Unmarshal(buf, data); err != nil {
		return err
	}
	// Validate configuration data.
	if err := data.Validate(); err != nil {
		return err
	}
	return nil
}

// // SoftInit checks for the existence of the config file and initializes it if it does not exist.
// func (cm *ConfigManager) SoftInit() error {
// 	exists, err := _file.Exists(cm.Path())
// 	if err != nil {
// 		return err
// 	}
// 	if !exists {
// 		return cm.Init()
// 	}
// 	return nil
// }

// // Init creates the configuration file and its directory if they do not exist.
// func (cm *ConfigManager) Init() error {
// 	err := _dir.EnsureDir(cm.DirPath(), _dir.DefaultDirPerms)
// 	if err != nil {
// 		return err
// 	}

// 	file, err := _file.Create(cm.Path(), _file.DefaultFilePerms)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// TODO Writing default configuration as YAML
// 	// defaultConfig := new(T) // Create a zero value for T to marshal into YAML
// 	// data, err := yaml.Marshal(defaultConfig)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to marshal default config: %w", err)
// 	// }
// 	// _, err = file.Write(data)
// 	// if err != nil {
// 	// 	return fmt.Errorf("failed to write default config to file %s: %w", cm.Path(), err)
// 	// }

// 	return nil
// }
