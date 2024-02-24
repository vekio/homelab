package homelab

import (
	"os"
	"text/template"

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

func ParseTemplate(src, dst string, data interface{}) error {
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
		return err
	}

	return nil
}
