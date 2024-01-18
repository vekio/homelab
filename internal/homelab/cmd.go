package homelab

// func commands() []*cli.Command {
// 	var commands []*cli.Command

// 	for _, srv := range services.AvailableServices() {
// 		commands = append(commands, serviceCmdFactory(srv))
// 	}

// 	commands = append(commands, initCmd, allUpCmd, allDownCmd)

// 	return commands
// }

// var initCmd = &cli.Command{
// 	Name:    "init",
// 	Aliases: []string{"i"},
// 	Usage:   "Initialize required folders and config files",
// 	Action: func(cCtx *cli.Context) error {
// 		if err := services.InitAuthelia(); err != nil {
// 			return err
// 		}

// 		if err := services.InitGitea(); err != nil {
// 			return err
// 		}

// 		if err := services.InitImmich(); err != nil {
// 			return err
// 		}

// 		if err := services.InitJellyfin(); err != nil {
// 			return err
// 		}

// 		if err := services.InitLldap(); err != nil {
// 			return err
// 		}

// 		if err := services.InitTraefik(); err != nil {
// 			return err
// 		}

// 		return nil
// 	},
// }

// var allUpCmd = &cli.Command{
// 	Name:    "allup",
// 	Aliases: []string{"au"},
// 	Usage:   "Create and start all service containers",
// 	Action: func(cCtx *cli.Context) (err error) {
// 		// Order by priority
// 		err = execDockerCompose(services.TRAEFIK, "up", "-d")
// 		if err != nil {
// 			return err
// 		}

// 		if err = os.Setenv("PROTONMAIL_BRIDGE_COMMAND", ""); err != nil {
// 			return err
// 		}
// 		if err = execDockerCompose(services.PROTONMAIL_BRIDGE, "up", "-d"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.LLDAP, "up", "-d"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.AUTHELIA, "up", "-d"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.GITEA, "up", "-d"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.JELLYFIN, "up", "-d"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.IMMICH, "up", "-d"); err != nil {
// 			return err
// 		}

// 		return
// 	},
// }

// var allDownCmd = &cli.Command{
// 	Name:    "alldown",
// 	Aliases: []string{"ad"},
// 	Usage:   "Stop and remove services containers, networks and volumes",
// 	Action: func(cCtx *cli.Context) (err error) {
// 		// Order by less priority
// 		if err = execDockerCompose(services.IMMICH, "down", "-v"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.JELLYFIN, "down", "-v"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.GITEA, "down", "-v"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.AUTHELIA, "down", "-v"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.LLDAP, "down", "-v"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.TRAEFIK, "down", "-v"); err != nil {
// 			return err
// 		}

// 		if err = execDockerCompose(services.PROTONMAIL_BRIDGE, "down", "-v"); err != nil {
// 			return err
// 		}

// 		return
// 	},
// }
