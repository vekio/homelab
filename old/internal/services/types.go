package services

import (
	"fmt"
	"os"
	"os/exec"
)

type Services map[string]*Service

type Initialize func() error

type Service struct {
	Name               string
	Context            string
	ComposeFile        string
	TraefikComposeFile string
	Priority           int
	Init               Initialize
	AllInitCmd         bool
	AllUpCmd           bool
	AllDownCmd         bool
}

type Compose interface {
	Config()
	Down()
	Logs()
	Pull()
	Restart()
	Stop()
	Up()
	Upgrade()
}

func (s Service) Config() error {
	if err := s.ExComposeCmd("config"); err != nil {
		return err
	}
	return nil
}

func (s Service) Down() error {
	if err := s.ExComposeCmd("down", "-v"); err != nil {
		return err
	}
	return nil
}

func (s Service) Logs() error {
	if err := s.ExComposeCmd("logs", "-f"); err != nil {
		return err
	}
	return nil
}

func (s Service) Pull() error {
	if err := s.ExComposeCmd("pull"); err != nil {
		return err
	}
	return nil
}

func (s Service) Restart() error {
	if err := s.ExComposeCmd("restart"); err != nil {
		return err
	}
	return nil
}

func (s Service) Stop() error {
	if err := s.ExComposeCmd("stop"); err != nil {
		return err
	}
	return nil
}

func (s Service) Up() error {
	if err := s.ExComposeCmd("up", "-d"); err != nil {
		return err
	}
	return nil
}

func (s Service) Upgrade() error {
	if err := s.Pull(); err != nil {
		return err
	}
	if err := s.Up(); err != nil {
		return err
	}
	return nil
}

func (s Service) ExComposeCmd(command ...string) error {
	cmdArgs := append([]string{"docker", "compose", "-f", s.ComposeFile}, command...)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing docker compose command: %w", err)
	}
	return nil
}
