package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func ParseConfig(src, dst string, data interface{}) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	tmpl, err := template.New("").Parse(string(content))
	if err != nil {
		return err
	}

	output, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer output.Close()

	if err := tmpl.Execute(output, data); err != nil {
		return fmt.Errorf("executing template for %s: %w", filepath.Base(dst), err)
	}

	return nil
}
