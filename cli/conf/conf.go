package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/vekio/fs"
	"github.com/vekio/fs/dir"
	"github.com/vekio/fs/file"
)

type C struct {
	name string
	dir  string
	file string
}

var Config C

func init() {
	// Get user default config folder
	dir, _ := os.UserConfigDir()

	// Check command config in user config
	exePath, _ := os.Executable()
	exePath, _ = filepath.EvalSymlinks(exePath)
	exeName := strings.TrimSuffix(
		filepath.Base(exePath), filepath.Ext(exePath),
	)

	Config = C{
		name: exeName,
		dir:  dir,
		file: "config.yml",
	}
}

// DirPath is the Dir and Name joined.
func (c C) DirPath() string { return filepath.Join(c.dir, c.name) }

// Path returns the combined Dir and File.
func (c C) Path() string { return filepath.Join(c.dir, c.name, c.file) }

// Exists returns true if a configuration file exists at Path.
func (c C) Exists() (bool, error) {
	exists, err := file.Exists(c.Path())
	if err != nil {
		return false, fmt.Errorf("conf: %w", err)
	}
	return exists, nil
}

// Data returns a string buffer containing all of the configuration file
// data for the given configuration. An empty string is returned and an
// error logged if any error occurs.
func (c C) Data() (string, error) {
	buf, err := os.ReadFile(c.Path())
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// Print prints the Data to standard output with an additional line
// return.
func (c C) Print() error {
	data, err := c.Data()
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

// Init initializes the configuration directory (Dir) for the current
// user and given application name (Name). The directory is completely
// removed and new configuration file(s) are created.
func (c C) Init() error {
	d := c.DirPath()

	exists, err := dir.Exists(d)
	if err != nil {
		return fmt.Errorf("error init config: %w", err)
	}

	if exists {
		if err := os.RemoveAll(d); err != nil {
			return fmt.Errorf("error init config: %w", err)
		}
	}

	if err := fs.CreateDir(d, fs.DefaultDirPerms); err != nil {
		return fmt.Errorf("error init config: %w", err)
	}

	if err := file.Touch(c.Path(), fs.DefaultFilePerms); err != nil {
		return fmt.Errorf("error init config: %w", err)
	}

	return nil
}

// SoftInit calls Init if not Exists.
func (c C) SoftInit() error {
	exists, err := c.Exists()
	if err != nil {
		return fmt.Errorf("error at soft-init config: %w", err)
	}

	if !exists {
		err = c.Init()
		if err != nil {
			return fmt.Errorf("error at soft-init config: %w", err)
		}
	}
	return nil
}
