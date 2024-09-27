package services

import (
	"fmt"

	"github.com/spf13/cobra"
	H "github.com/vekio/homelab/internal/homelab"
)

func NewCmdServices(homelab H.Homelab) *cobra.Command {
	servicesCmd := &cobra.Command{
		Use:     "services",
		Aliases: []string{"srv"},
		Short:   "Manage homelab services",
	}

	// Subcommands
	servicesCmd.AddCommand(newCmdList(homelab))

	// Compose Subcommands
	servicesCmd.AddCommand(
		newCmdConfig(homelab),
		newCmdDown(homelab),
		newCmdLogs(homelab),
		newCmdRestart(homelab),
		newCmdStop(homelab),
		newCmdPull(homelab),
		newCmdUp(homelab),
	)

	return servicesCmd
}

func newCmdCompose(name, shortDesc string, homelab H.Homelab, exec func(s *H.Service) error) *cobra.Command {
	cmd := &cobra.Command{
		Use:       fmt.Sprintf("%s SERVICE", name),
		Short:     shortDesc,
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		GroupID:   "compose",
		ValidArgs: homelab.ServicesNames(),
		RunE: func(cmd *cobra.Command, args []string) error {
			service, err := homelab.ServiceByName(args[0])
			if err != nil {
				return err
			}

			if err := exec(service); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func newCmdConfig(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("config", "Parse, resolve and render compose file in canonical format", homelab,
		func(s *H.Service) error {
			return s.Config()
		})
}

func newCmdDown(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("down", "Stop and remove containers, networks", homelab,
		func(s *H.Service) error {
			return s.Down()
		})
}

func newCmdLogs(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("logs", "View output from containers", homelab,
		func(s *H.Service) error {
			return s.Logs()
		})
}

func newCmdPull(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("pull", "Pull service images", homelab,
		func(s *H.Service) error {
			return s.Pull()
		})
}

func newCmdRestart(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("restart", "Restart service containers", homelab,
		func(s *H.Service) error {
			return s.Restart()
		})
}

func newCmdStop(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("stop", "Stop services", homelab,
		func(s *H.Service) error {
			return s.Stop()
		})
}

func newCmdUp(homelab H.Homelab) *cobra.Command {
	return newCmdCompose("up", "Create and start containers", homelab,
		func(s *H.Service) error {
			return s.Up()
		})
}
