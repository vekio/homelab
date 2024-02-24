package config

import (
	"os"
	"path/filepath"
)

var Manager configManager

func init() {
	dir, _ := os.UserConfigDir()

	Manager = configManager{
		dir:  dir,
		file: "config.yml",
		name: "homelab",
	}
}

func (cm configManager) DirPath() string {
	return filepath.Join(cm.dir, cm.name)
}

func (cm configManager) Path() string {
	return filepath.Join(cm.dir, cm.name, cm.file)
}
