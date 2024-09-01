package config

import (
	"fmt"
	"os"
	"path/filepath"

	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"gopkg.in/yaml.v3"
)

// ConfigManager manages configuration files for an application.
type ConfigManager[T Validatable] struct {
	appName string // Name of the application
	dir     string // Directory where the config files are stored
	file    string // Configuration file name
	Data    T      // Stored configuration data
}

// Validatable is an interface that should be implemented by all config types
// that will be managed by ConfigManager.
type Validatable interface {
	Validate() error
}

// NewConfigManager creates a new instance of ConfigManager for an application.
func NewConfigManager[T Validatable](appName, configName string) (*ConfigManager[T], error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user config directory: %w", err)
	}

	cm := &ConfigManager[T]{
		dir:     dir,
		appName: appName,
		file:    configName,
	}

	// Initialize the configuration file.
	if err := cm.SoftInit(); err != nil {
		return nil, fmt.Errorf("failed to initialize configuration: %w", err)
	}

	// Deserialize the configuration file.
	buf, err := cm.Content()
	if err != nil {
		return nil, err
	}

	var config T
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	cm.Data = config

	return cm, nil
}

// DirPath returns the path to the directory where the configuration file is stored.
func (cm *ConfigManager[T]) DirPath() string {
	return filepath.Join(cm.dir, cm.appName)
}

// Path returns the full path to the configuration file.
func (cm *ConfigManager[T]) Path() string {
	return filepath.Join(cm.DirPath(), cm.file)
}

// SoftInit checks for the existence of the config file and initializes it if it does not exist.
func (cm *ConfigManager[T]) SoftInit() error {
	exists, err := _file.Exists(cm.Path())
	if err != nil {
		return err
	}
	if !exists {
		return cm.Init()
	}
	return nil
}

// Init creates the configuration file and its directory if they do not exist.
func (cm *ConfigManager[T]) Init() error {
	err := _dir.EnsureDir(cm.DirPath(), _dir.DefaultDirPerms)
	if err != nil {
		return err
	}

	file, err := _file.Create(cm.Path(), _file.DefaultFilePerms)
	if err != nil {
		return err
	}
	defer file.Close()

	// TODO Writing default configuration as YAML
	// defaultConfig := new(T) // Create a zero value for T to marshal into YAML
	// data, err := yaml.Marshal(defaultConfig)
	// if err != nil {
	// 	return fmt.Errorf("failed to marshal default config: %w", err)
	// }
	// _, err = file.Write(data)
	// if err != nil {
	// 	return fmt.Errorf("failed to write default config to file %s: %w", cm.Path(), err)
	// }

	return nil
}

// Content reads and returns the contents of the configuration file.
func (cm *ConfigManager[T]) Content() ([]byte, error) {
	return os.ReadFile(cm.Path())
}
