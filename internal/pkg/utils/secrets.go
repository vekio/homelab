package utils

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)

func WriteSecret(filename, secret string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(secret)
	if err != nil {
		return err
	}

	return nil
}

func Bcrypt(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	hashedString := string(hashedBytes)
	return hashedString, nil
}
