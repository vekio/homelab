package services

import (
	"fmt"

	"github.com/spf13/cobra"
	H "github.com/vekio/homelab/internal/homelab"
)

// NewCmdServices initializes the root 'services' command for managing homelab services.
func NewCmdServices(homelab H.Homelab) *cobra.Command {
	servicesCmd := &cobra.Command{
		Use:     "services",
		Aliases: []string{"srv"}, // Optional alias for the command.
		Short:   "Manage homelab services",
	}

	// Add subcommands.
	servicesCmd.AddCommand(newCmdList(homelab))

	// Add docker compose-related subcommands to the 'services' command.
	servicesCmd.AddCommand(
		newCmdConfig(homelab),
		newCmdDown(homelab),
		newCmdLogs(homelab),
		newCmdRestart(homelab),
		newCmdStop(homelab),
		newCmdPull(homelab),
		newCmdUp(homelab),
	)

	// Return the constructed 'services' command with all subcommands attached.
	return servicesCmd
}

// newCmdCompose creates a generic command for a Docker Compose action.
// It accepts the command name, short description, homelab instance, and an exec function that will be run on a specific service.
func newCmdCompose(name, shortDesc string, homelab H.Homelab, exec func(s *H.Service) error) *cobra.Command {
	// Define the command with usage, short description, and argument validation.
	cmd := &cobra.Command{
		Use:       fmt.Sprintf("%s SERVICE", name),
		Short:     shortDesc,
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		GroupID:   "compose",               // Grouping for docker compose-related commands.
		ValidArgs: homelab.ServicesNames(), // Set the valid service names as arguments.
		RunE: func(cmd *cobra.Command, args []string) error {
			// Get the service by its name from the homelab.
			service, err := homelab.ServiceByName(args[0])
			if err != nil {
				return err
			}

			// Execute the provided function (like config, down, etc.) on the service.
			if err := exec(service); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

// newCmdConfig creates the 'config' command, which parses, resolves, and renders the Compose file.
func newCmdConfig(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("config", "Parse, resolve and render compose file in canonical format", homelab,
		func(s *H.Service) error {
			// Executes the Config function on the service.
			return s.Config()
		})
}

// newCmdDown creates the 'down' command to stop and remove containers, networks, and volumes.
func newCmdDown(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("down", "Stop and remove containers, networks", homelab,
		func(s *H.Service) error {
			// Executes the Down function on the service.
			return s.Down()
		})
}

// newCmdLogs creates the 'logs' command to view output from the service containers.
func newCmdLogs(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("logs", "View output from containers", homelab,
		func(s *H.Service) error {
			// Executes the Logs function on the service.
			return s.Logs()
		})
}

// newCmdPull creates the 'pull' command to pull the latest service images.
func newCmdPull(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("pull", "Pull service images", homelab,
		func(s *H.Service) error {
			// Executes the Pull function on the service.
			return s.Pull()
		})
}

// newCmdRestart creates the 'restart' command to restart service containers.
func newCmdRestart(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("restart", "Restart service containers", homelab,
		func(s *H.Service) error {
			// Executes the Restart function on the service.
			return s.Restart()
		})
}

// newCmdStop creates the 'stop' command to stop running service containers.
func newCmdStop(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("stop", "Stop services", homelab,
		func(s *H.Service) error {
			// Executes the Stop function on the service.
			return s.Stop()
		})
}

// newCmdUp creates the 'up' command to start service containers in detached mode.
func newCmdUp(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("up", "Create and start containers", homelab,
		func(s *H.Service) error {
			// Executes the Up function on the service.
			return s.Up()
		})
}
