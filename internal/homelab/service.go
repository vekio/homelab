package homelab

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	_dir "github.com/vekio/fs/dir"
	"github.com/vekio/homelab/internal/utils"
)

type Services = map[string]*Service

type Service struct {
	Name         string
	Context      string
	ComposeFiles []string
	ExtraFiles   []string
}

// NewService initializes a new Service instance, ensures its directory exists,
// builds the URLs for its compose files, and downloads those files.
func NewService(name, context string, composeFiles, extraFiles []string) (*Service, error) {
	// Check if the context is set to "local". If so, switch it to "default".
	// "local" is treated as an alias for the "default" docker context.
	if context == "local" {
		context = "default"
	}

	// Create a new service instance with the given name, context, and compose files.
	s := &Service{
		Name:         name,
		Context:      context,
		ComposeFiles: composeFiles,
		ExtraFiles:   extraFiles,
	}

	// Ensure that the directory for the service exists. If not, create it.
	if err := _dir.EnsureDir(s.ServicePath(), _dir.DefaultDirPerms); err != nil {
		return nil, err
	}

	// Download compose files.
	if err := s.DownloadComposeFiles(); err != nil {
		return nil, err
	}

	// Download extra files.
	if err := s.DownloadExtraFiles(); err != nil {
		return nil, err
	}

	// Return the newly created service instance.
	return s, nil
}

func (s Service) DownloadExtraFiles() error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(s.ExtraFiles)) // Channel to handle any errors that occur concurrently.

	// Iterate over the extra files and process each in a separate goroutine.
	for _, file := range s.ExtraFiles {
		wg.Add(1)

		go func(file string) {
			defer wg.Done()

			// Concatenate the service name with the extra file to generate the correct file path.
			fileURLPath := fmt.Sprintf("%s/%s", s.Name, file)

			// Build the GitHub URL for the extra file using the repository URL and branch from settings.
			url, err := utils.GenerateGithubURL(settings.Repository.URL, settings.Repository.Branch, fileURLPath)
			if err != nil {
				// Send the error to the error channel if URL generation fails.
				errCh <- fmt.Errorf("error generating URL for file %s: %w", file, err)
				return
			}

			extraPath := fmt.Sprintf("%s/%s", s.ServicePath(), filepath.Dir(file))

			if err := _dir.EnsureDir(extraPath, _dir.DefaultDirPerms); err != nil {
				errCh <- err
			}

			// Download the extra file using the generated URL, saving it in the service's directory.
			if err := utils.DownloadFile(url, extraPath); err != nil {
				// Send the error to the error channel if the download fails.
				errCh <- fmt.Errorf("error downloading file %s: %w", file, err)
				return
			}
		}(file)
	}

	// Wait for all goroutines to complete.
	wg.Wait()
	close(errCh) // Close the error channel after all tasks have completed.

	// Check for any errors that may have occurred during execution.
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

// Concatenate the service name with each compose file to generate the correct file paths and download them in parallel.
func (s Service) DownloadComposeFiles() error {
	var wg sync.WaitGroup
	errCh := make(chan error, len(s.ComposeFiles)) // Channel to handle any errors that occur concurrently.

	// Iterate over the compose files and process each in a separate goroutine.
	for _, file := range s.ComposeFiles {
		wg.Add(1)

		go func(file string) {
			defer wg.Done()

			// Concatenate the service name with the compose file to generate the correct file path.
			fileURLPath := fmt.Sprintf("%s/%s", s.Name, file)

			// Build the GitHub URL for the compose file using the repository URL and branch from settings.
			url, err := utils.GenerateGithubURL(settings.Repository.URL, settings.Repository.Branch, fileURLPath)
			if err != nil {
				// Send the error to the error channel if URL generation fails.
				errCh <- fmt.Errorf("error generating URL for file %s: %w", file, err)
				return
			}

			// Download the compose file using the generated URL, saving it in the service's directory.
			if err := utils.DownloadFile(url, s.ServicePath()); err != nil {
				// Send the error to the error channel if the download fails.
				errCh <- fmt.Errorf("error downloading file %s: %w", file, err)
				return
			}
		}(file)
	}

	// Wait for all goroutines to complete.
	wg.Wait()
	close(errCh) // Close the error channel after all tasks have completed.

	// Check for any errors that may have occurred during execution.
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

// ServicePath returns the service directory path where service's files are stored.
func (s Service) ServicePath() string {
	return fmt.Sprintf("%s/%s", composeFilesBasePath(), s.Name)
}

// ComposeFilePaths returns a slice of file paths for the service's docker compose files.
func (s Service) ComposeFilePaths() []string {
	var composeFilesPaths []string
	for _, file := range s.ComposeFiles {
		// Append the full path for each compose file by combining the service path and file name.
		composeFilesPaths = append(composeFilesPaths, fmt.Sprintf("%s/%s", s.ServicePath(), file))
	}
	return composeFilesPaths
}

// ExtraFilePaths returns a slice of file paths for the service's extra files.
func (s Service) ExtraFilePaths() []string {
	var extraFilesPaths []string
	for _, file := range s.ExtraFiles {
		// Append the full path for each extra file by combining the service path and file name.
		extraFilesPaths = append(extraFilesPaths, fmt.Sprintf("%s/%s", s.ServicePath(), file))
	}
	return extraFilesPaths
}

// execComposeCmd constructs and executes a docker compose command for the service.
// It takes a variable number of command arguments and executes the Docker Compose command.
func (s Service) execComposeCmd(command ...string) error {
	// Change to the service context
	// TODO mejorar el cambio de contexto
	// Change Docker context using "docker context use"
	contextCmd := exec.Command("docker", "context", "use", s.Context)
	contextCmd.Stdin = os.Stdin
	contextCmd.Stdout = os.Stdout
	contextCmd.Stderr = os.Stderr

	// Execute the context change command
	if err := contextCmd.Run(); err != nil {
		return err
	}

	// Build the list of arguments for the docker compose command by adding the -f flag for each compose file.
	var composeFileArgs []string
	for _, file := range s.ComposeFilePaths() {
		composeFileArgs = append(composeFileArgs, "-f", file)
	}

	// Prepend "docker compose" to the arguments, followed by the compose file arguments.
	cmdArgs := append([]string{"docker", "compose"}, composeFileArgs...)

	// Append the env-file path.
	cmdArgs = append(cmdArgs, "--env-file", envFilePath())

	// Append the actual command (e.g., "up", "down", "logs") passed as a variadic argument.
	cmdArgs = append(cmdArgs, command...)

	// Create the command to be executed.
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Execute the command and return any error if the execution fails.
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("error executing: %w", err)
	}
	return nil
}

// Config runs the `docker compose config` command to validate and view the Compose file configuration.
func (s Service) Config() error {
	return s.execComposeCmd("config")
}

// Down runs the `docker compose down -v` command to stop and remove containers, networks, and volumes.
func (s Service) Down() error {
	return s.execComposeCmd("down", "-v")
}

// Logs runs the `docker compose logs -f` command to follow the logs of all running services.
func (s Service) Logs() error {
	return s.execComposeCmd("logs", "-f")
}

// Pull runs the `docker compose pull` command to pull the latest images for the services.
func (s Service) Pull() error {
	return s.execComposeCmd("pull")
}

// Restart runs the `docker compose restart` command to restart all running services.
func (s Service) Restart() error {
	return s.execComposeCmd("restart")
}

// Stop runs the `docker compose stop` command to stop all running services without removing containers.
func (s Service) Stop() error {
	return s.execComposeCmd("stop")
}

// Up runs the `docker compose up -d` command to start services in detached mode.
func (s Service) Up() error {
	return s.execComposeCmd("up", "-d")
}
