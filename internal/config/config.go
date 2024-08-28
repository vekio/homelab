package config

import (
	"os"
	"path/filepath"

	_fs "github.com/vekio/fs"
	_file "github.com/vekio/fs/file"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Services map[string]ServiceConfig `yaml:",flow"`
}

type ServiceConfig struct {
	Server string `yaml:"server"`
}

// ConfigManager manages configuration files for an application.
type ConfigManager struct {
	appName string // Name of the application
	dir     string // Directory where the config files are stored
	file    string // Configuration file name
}

// NewConfigManager creates a new instance of ConfigManager for an application.
// It sets the directory to the user's config directory and initializes the configuration file name to "config.yml".
func NewConfigManager(appName string) ConfigManager {
	dir, _ := os.UserConfigDir() // TODO Handle error

	conf := ConfigManager{
		dir:     dir,
		appName: appName,
		file:    "config.yml",
	}
	return conf
}

// DirPath returns the path to the directory where the configuration file is stored.
func (cm ConfigManager) DirPath() string {
	return filepath.Join(cm.dir, cm.appName)
}

// Path returns the full path to the configuration file.
func (cm ConfigManager) Path() string {
	return filepath.Join(cm.dir, cm.appName, cm.file)
}

// SoftInit checks for the existence of the config file and initializes it if it does not exist.
func (cm ConfigManager) SoftInit() error {
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
// If the file already exists, Init truncates the content in the file.
func (cm ConfigManager) Init() error {
	file, err := _fs.CreateFileWithDirs(cm.Path(), _fs.DefaultFilePerms)
	if err != nil {
		return err
	}
	defer file.Close()

	// TODO Escribir configuración predeterminada en el archivo.
	// It handles file permissions and prepares the file to be written with default configuration data.
	// _, err = file.WriteString("default config data")
	// if err != nil {
	// 	return fmt.Errorf("failed to write to config file %s: %w", cm.Path(), err)
	// }

	return nil
}

// Content reads and returns the contents of the configuration file.
func (cm ConfigManager) Content() ([]byte, error) {
	return os.ReadFile(cm.Path())
}

func (cm ConfigManager) Data() (Config, error) {
	buf, err := cm.Content()
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
