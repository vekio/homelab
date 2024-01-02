package utils

import (
	"os"

	_rsa "github.com/vekio/crypto/rsa"
)

func GenerateRSAPrivateKey(filename string) error {

	bits := 4096

	privateKey, err := _rsa.GenerateKeyPair(bits)
	if err != nil {
		return err
	}

	data, err := _rsa.PrivateKeyData(privateKey)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data + "\r\n")
	if err != nil {
		return err
	}

	return nil
}
