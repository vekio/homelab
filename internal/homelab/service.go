package homelab

import (
	"fmt"
	"os"
	"os/exec"

	_dir "github.com/vekio/fs/dir"
	"github.com/vekio/homelab/internal/utils"
	"github.com/vekio/homelab/pkg/config"
)

type Services = map[string]*Service

type Service struct {
	Name         string
	Context      string
	ComposeFiles []string
}

// NewService initializes a new Service instance, ensures its directory exists,
// builds the URLs for its compose files, and downloads those files.
func NewService(name, context string, composeFiles []string) (*Service, error) {
	// Create a new service instance with the given name, context, and compose files.
	s := &Service{
		Name:         name,
		Context:      context,
		ComposeFiles: composeFiles,
	}

	// Ensure that the directory for the service exists. If not, create it.
	if err := _dir.EnsureDir(s.ServicePath(), _dir.DefaultDirPerms); err != nil {
		return nil, err
	}

	// Concatenate the service name with each compose file to generate the correct file paths.
	var filePaths []string
	for _, file := range s.ComposeFiles {
		filePaths = append(filePaths, fmt.Sprintf("%s/%s", s.Name, file))
	}

	// Build the GitHub URLs for each compose file using the repository URL and branch from settings.
	urls, err := utils.GenerateGithubURLs(settings.Repository.URL, settings.Repository.Branch, filePaths)
	if err != nil {
		// Return an error if URL generation fails.
		return nil, err
	}

	// Download the compose files using the generated URLs, saving them in the service's directory.
	if err := utils.DownloadFiles(urls, s.ServicePath()); err != nil {
		// Return an error if the download fails.
		return nil, err
	}

	// Return the newly created service instance.
	return s, nil
}

// ServicePath returns the directory path where the docker compose files for the service are stored.
func (s Service) ServicePath() string {
	return fmt.Sprintf("%s/%s/%s", composeFilesBasePath(), config.DirPath(), s.Name)
}

// ComposeFilePaths returns a slice of file paths for the service's Docker Compose files.
func (s Service) ComposeFilePaths() []string {
	var composeFilesPaths []string
	for _, file := range s.ComposeFiles {
		// Append the full path for each compose file by combining the service path and file name.
		composeFilesPaths = append(composeFilesPaths, fmt.Sprintf("%s/%s", s.ServicePath(), file))
	}
	return composeFilesPaths
}

// execComposeCmd constructs and executes a docker compose command for the service.
// It takes a variable number of command arguments and executes the Docker Compose command.
func (s Service) execComposeCmd(command ...string) error {
	// Build the list of arguments for the docker compose command by adding the -f flag for each compose file.
	var composeFileArgs []string
	for _, file := range s.ComposeFilePaths() {
		composeFileArgs = append(composeFileArgs, "-f", file)
	}

	// Prepend "docker compose" to the arguments, followed by the compose file arguments.
	cmdArgs := append([]string{"docker", "compose"}, composeFileArgs...)
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
