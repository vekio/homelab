package homelab

import (
	"fmt"
	"log"

	Z "github.com/rwxrob/bonzai/z"
	"github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/vars"
	"github.com/vekio/homelab/internal/homelab"
	"github.com/vekio/homelab/internal/secrets"
	"github.com/vekio/homelab/internal/services"
)

var Version string

func init() {
	err := Z.Conf.SoftInit()
	if err != nil {
		log.Fatal(err)
	}

	err = Z.Vars.SoftInit()
	if err != nil {
		log.Fatal(err)
	}

	err = secrets.Manager.SoftInit()
	if err != nil {
		log.Fatal(err)
	}
}

func Cmd() *Z.Cmd {
	return &Z.Cmd{
		Name:        `homelab`,
		Commands:    buildCmds(),
		Version:     Version,
		License:     `Apache-2.0`,
		Copyright:   `(c) Alberto Castañeiras <alberto@casta.me>`,
		Summary:     `Manage my homelab server, services and configs`,
		Description: `Enough of multiple scripts`,
	}
}

func buildCmds() []*Z.Cmd {
	var cmds []*Z.Cmd

	// default commands
	defaultCmds := buildDefaultCmds()
	cmds = append(cmds, defaultCmds...)

	// services commands
	srvsCmds := buildServicesCmd()
	cmds = append(cmds, srvsCmds...)

	return cmds
}

func buildDefaultCmds() []*Z.Cmd {
	var cmds []*Z.Cmd

	// help command
	cmds = append(cmds, help.Cmd)

	// conf command
	cmds = append(cmds, conf.Cmd)

	// vars command
	cmds = append(cmds, vars.Cmd)

	return cmds
}

func buildServicesCmd() []*Z.Cmd {
	var cmds []*Z.Cmd

	app, err := homelab.New()
	if err != nil {
		log.Fatal(err)
	}

	for _, srv := range app.Services {
		cmd := &Z.Cmd{
			Name:    srv.Name,
			Summary: help.S(fmt.Sprintf("manage %s service", srv.Name)),
		}

		var initCmd = &Z.Cmd{
			Name:    "init",
			Summary: "initialize service's config files",
			Call: func(caller *Z.Cmd, args ...string) error {
				err := srv.Init()
				if err != nil {
					return err
				}
				return nil
			},
		}

		cmd.Commands = buildDefaultCmds()
		cmd.Commands = append(cmd.Commands, buildServiceCmds(srv)...)
		cmd.Commands = append(cmd.Commands, initCmd)

		cmds = append(cmds, cmd)
	}

	var allUpCmd = &Z.Cmd{
		Name:    "allup",
		Summary: "fires up all services tagged with allup",
		Call: func(caller *Z.Cmd, args ...string) error {
			app.Services.AllUp()
			return nil
		},
	}
	cmds = append(cmds, allUpCmd)

	var allDownCmd = &Z.Cmd{
		Name:    "alldown",
		Summary: "down all services tagged with alldown",
		Call: func(caller *Z.Cmd, args ...string) error {
			app.Services.AllDown()
			return nil
		},
	}
	cmds = append(cmds, allDownCmd)

	var allInitCmd = &Z.Cmd{
		Name:    "allinit",
		Summary: "initializes all services tagged with allinit",
		Call: func(caller *Z.Cmd, args ...string) error {
			app.Services.AllInit()
			return nil
		},
	}
	cmds = append(cmds, allInitCmd)

	return cmds
}

func buildServiceCmds(srv *services.Service) []*Z.Cmd {
	var cmds []*Z.Cmd

	var configCmd = &Z.Cmd{
		Name:    "config",
		Summary: "render compose file in canonical format",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Config()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, configCmd)

	var downCmd = &Z.Cmd{
		Name:    "down",
		Summary: "stop and remove containers, networks",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Down()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, downCmd)

	var logsCmd = &Z.Cmd{
		Name:    "logs",
		Summary: "view output from containers",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Logs()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, logsCmd)

	var pullCmd = &Z.Cmd{
		Name:    "pull",
		Summary: "pull service images",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Pull()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, pullCmd)

	var restartCmd = &Z.Cmd{
		Name:    "restart",
		Summary: "restart service containers",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Restart()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, restartCmd)

	var stopCmd = &Z.Cmd{
		Name:    "stop",
		Summary: "stop services",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Stop()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, stopCmd)

	var upCmd = &Z.Cmd{
		Name:    "up",
		Summary: "create and start containers",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Up()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, upCmd)

	var upgradeCmd = &Z.Cmd{
		Name:    "upgrade",
		Summary: "create and start containers",
		Call: func(caller *Z.Cmd, args ...string) error {
			err := srv.Up()
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmds = append(cmds, upgradeCmd)

	return cmds
}
