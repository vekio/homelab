package utils

import "os"

func WriteSecret(filename, secret string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(secret + "\r\n")
	if err != nil {
		return err
	}

	return nil
}
