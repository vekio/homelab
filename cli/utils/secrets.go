package utils

import (
	"os"

	_secretgen "github.com/vekio/rand/secretgen"
)

type Secret struct {
	Name   string
	Length int
}

func CreateAlphanumericSecret(filename string, n int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	s, err := _secretgen.GenerateRandomAlphaNumeric(n)
	if err != nil {
		return err
	}

	_, err = file.WriteString(s + "\r\n")
	if err != nil {
		return err
	}

	return nil
}
