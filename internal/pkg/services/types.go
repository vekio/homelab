package services

import (
	"fmt"
	"log/slog"
	"os"
	"os/exec"
)

type Service struct {
	Name        string
	ComposeFile string
	Context     string
	Init        Initialize
	Priority    int
}

type Initialize func() error

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
	slog.Debug("docker compose config",
		slog.String("service", s.Name))
	if err := s.exComposeCmd("config"); err != nil {
		return err
	}
	return nil
}

func (s Service) Down() error {
	slog.Debug("docker compose down",
		slog.String("service", s.Name))
	if err := s.exComposeCmd("down", "-v"); err != nil {
		return err
	}
	return nil
}

func (s Service) Logs() error {
	slog.Debug("docker compose logs",
		slog.String("service", s.Name))
	if err := s.exComposeCmd("logs", "-f"); err != nil {
		return err
	}
	return nil
}

func (s Service) Pull() error {
	slog.Debug("docker compose pull",
		slog.String("service", s.Name))
	if err := s.exComposeCmd("pull"); err != nil {
		return err
	}
	return nil
}

func (s Service) Restart() error {
	slog.Debug("docker compose restart",
		slog.String("service", s.Name))
	if err := s.exComposeCmd("restart"); err != nil {
		return err
	}
	return nil
}

func (s Service) Stop() error {
	slog.Debug("docker compose stop",
		slog.String("service", s.Name))
	if err := s.exComposeCmd("stop"); err != nil {
		return err
	}
	return nil
}

func (s Service) Up() error {
	slog.Debug("docker compose up",
		slog.String("service", s.Name))
	if err := s.exComposeCmd("up", "-d"); err != nil {
		return err
	}
	return nil
}

func (s Service) Upgrade() error {
	slog.Debug("docker compose upgrade",
		slog.String("service", s.Name))
	if err := s.Pull(); err != nil {
		return err
	}
	if err := s.Up(); err != nil {
		return err
	}
	return nil
}

func (s Service) exComposeCmd(command ...string) error {
	cmdArgs := append([]string{"docker", "compose", "-f", s.ComposeFile}, command...)
	slog.Debug("executing docker compose command",
		slog.String("command", fmt.Sprint(cmdArgs)))

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		err := fmt.Errorf("error executing docker compose command: %w", err)

		slog.Error(err.Error(),
			slog.String("command", fmt.Sprint(cmdArgs)))

		return err
	}
	return nil
}
