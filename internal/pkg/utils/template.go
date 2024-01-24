package utils

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func ParseTemplate(src, dst string, data interface{}) error {
	content, err := os.ReadFile(src)
	if err != nil {
		return fmt.Errorf("ParseTemplate: error reading source %s: %w", src, err)
	}

	tmpl, err := template.New("").Parse(string(content))
	if err != nil {
		return fmt.Errorf("ParseTemplate: error creating template %s: %w", dst, err)
	}

	output, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("ParseTemplate: error creating destiny %s: %w", dst, err)
	}
	defer output.Close()

	if err := tmpl.Execute(output, data); err != nil {
		return fmt.Errorf("ParseTemplate: error executing template for %s: %w", filepath.Base(dst), err)
	}

	return nil
}
