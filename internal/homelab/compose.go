package homelab

import (
	"fmt"
	"os"
	"os/exec"
)

func (s Service) ExecComposeCmd(command ...string) error {
	var composeArgs []string
	for _, file := range s.ComposeFilesPath() {
		composeArgs = append(composeArgs, "-f", file)
	}
	composeArgs = append(composeArgs, command...)

	cmdArgs := []string{"docker", "compose"}
	cmdArgs = append(cmdArgs, composeArgs...)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing: %w", err)
	}
	return nil
}

func (s Service) Config() error {
	return s.ExecComposeCmd("config")
}

func (s Service) Down() error {
	return s.ExecComposeCmd("down", "-v")
}

func (s Service) Logs() error {
	return s.ExecComposeCmd("logs", "-f")
}

func (s Service) Pull() error {
	return s.ExecComposeCmd("pull")
}

func (s Service) Restart() error {
	return s.ExecComposeCmd("restart")
}

func (s Service) Stop() error {
	return s.ExecComposeCmd("stop")
}

func (s Service) Up() error {
	return s.ExecComposeCmd("up", "-d")
}

func (s Service) Upgrade() error {
	// if err := s.Pull(); err != nil {
	// 	return err
	// }
	// if err := s.Up(); err != nil {
	// 	return err
	// }
	return nil
}
