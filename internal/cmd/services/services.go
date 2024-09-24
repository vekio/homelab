package services

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vekio/homelab/internal/homelab"
)

func NewCmdServices(h homelab.Homelab) *cobra.Command {
	servicesCmd := &cobra.Command{
		Use:     "services",
		Aliases: []string{"srv"},
		Short:   "Manage homelab services",
	}

	// Subcommands
	servicesCmd.AddCommand(newCmdList(h))

	// Compose Subcommands
	servicesCmd.AddGroup(&cobra.Group{ID: "compose", Title: "Compose Commands"})
	servicesCmd.AddCommand(
		newCmdConfig(h), newCmdDown(h),
		newCmdLogs(h), newCmdRestart(h),
		newCmdStop(h), newCmdPull(h),
		newCmdUp(h))

	return servicesCmd
}

func newCmdCompose(name, shortDesc string, h homelab.Homelab, exec func(executor homelab.DockerComposeExecutor) error) *cobra.Command {
	cmd := &cobra.Command{
		Use:       fmt.Sprintf("%s SERVICE", name),
		Short:     shortDesc,
		Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		GroupID:   "compose",
		ValidArgs: h.ServicesNames(),
		RunE: func(cmd *cobra.Command, args []string) error {
			// Extract service from first arg.
			service, err := h.ServiceByName(args[0])
			if err != nil {
				return err
			}

			executor := homelab.NewDockerComposeExecutor(service)

			if err := exec(executor); err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func newCmdConfig(h homelab.Homelab) *cobra.Command {
	return newCmdCompose("config", "Parse, resolve and render compose file in canonical format", h,
		func(exec homelab.DockerComposeExecutor) error {
			return exec.Config()
		})
}

func newCmdDown(h homelab.Homelab) *cobra.Command {
	return newCmdCompose("down", "Stop and remove containers, networks", h,
		func(exec homelab.DockerComposeExecutor) error {
			return exec.Down()
		})
}

func newCmdLogs(h homelab.Homelab) *cobra.Command {
	return newCmdCompose("logs", "View output from containers", h,
		func(exec homelab.DockerComposeExecutor) error {
			return exec.Logs()
		})
}

func newCmdPull(h homelab.Homelab) *cobra.Command {
	return newCmdCompose("pull", "Pull service images", h,
		func(exec homelab.DockerComposeExecutor) error {
			return exec.Pull()
		})
}

func newCmdRestart(h homelab.Homelab) *cobra.Command {
	return newCmdCompose("restart", "Restart service containers", h,
		func(exec homelab.DockerComposeExecutor) error {
			return exec.Restart()
		})
}

func newCmdStop(h homelab.Homelab) *cobra.Command {
	return newCmdCompose("stop", "Stop services", h,
		func(exec homelab.DockerComposeExecutor) error {
			return exec.Stop()
		})
}

func newCmdUp(h homelab.Homelab) *cobra.Command {
	return newCmdCompose("up", "Create and start containers", h,
		func(exec homelab.DockerComposeExecutor) error {
			return exec.Up()
		})
}
