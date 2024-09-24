package homelab

import (
	"fmt"
	"os"
	"os/exec"
)

type Compose interface {
	Config() error
	Down() error
	Logs() error
	Pull() error
	Restart() error
	Stop() error
	Up() error
}

type DockerComposeExecutor struct {
	service *Service
}

func NewDockerComposeExecutor(service *Service) DockerComposeExecutor {
	dockerComposeExecutor := DockerComposeExecutor{
		service: service,
	}
	return dockerComposeExecutor
}

func (d *DockerComposeExecutor) composeFilesArgs() []string {
	var args []string
	for _, file := range d.service.ComposeFilesPath() {
		args = append(args, "-f", file)
	}
	return args
}

func (d DockerComposeExecutor) execComposeCmd(command ...string) error {
	cmdArgs := append([]string{"docker", "compose"}, d.composeFilesArgs()...)
	cmdArgs = append(cmdArgs, command...)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing: %w", err)
	}
	return nil
}

func (d DockerComposeExecutor) Config() error {
	return d.execComposeCmd("config")
}

func (d DockerComposeExecutor) Down() error {
	return d.execComposeCmd("down", "-v")
}

func (d DockerComposeExecutor) Logs() error {
	return d.execComposeCmd("logs", "-f")
}

func (d DockerComposeExecutor) Pull() error {
	return d.execComposeCmd("pull")
}

func (d DockerComposeExecutor) Restart() error {
	return d.execComposeCmd("restart")
}

func (d DockerComposeExecutor) Stop() error {
	return d.execComposeCmd("stop")
}

func (d DockerComposeExecutor) Up() error {
	return d.execComposeCmd("up", "-d")
}

// func (d DockerComposeExecutor) Upgrade() error {
// 	// if err := s.Pull(); err != nil {
// 	// 	return err
// 	// }
// 	// if err := s.Up(); err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }
