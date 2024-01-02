package utils

import (
	"os"

	_secretgen "github.com/vekio/rand/secretgen"
)

type Secret struct {
	Name   string
	Length int
}

// TODO 0600 permissions
func CreateAlphanumericSecret(filename string, n int) error {
	s, err := _secretgen.GenerateRandomAlphaNumeric(n)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(s + "\r\n")
	if err != nil {
		return err
	}

	return nil
}
