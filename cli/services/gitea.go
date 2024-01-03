package services

import (
	"fmt"
	"os"

	_fs "github.com/vekio/fs"
)

func InitGitea(envConfig string) error {
	// Create data folder
	dataDir := fmt.Sprintf("%s/data/", envConfig)
	err := _fs.CreateDir(dataDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	return nil
}
