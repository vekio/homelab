package homelab

import (
	"log"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/cli/conf"
	"github.com/vekio/homelab/cli/secrets"
)

func init() {
	filename := conf.Config.DirPath() + "/secrets.yml"
	if err := secrets.SoftInitSecrets(filename); err != nil {
		log.Fatalf("failed to init secrets %s", err)
	}
}

var testCmd = &cli.Command{
	Name:  "test",
	Usage: "testing",
	Action: func(cCtx *cli.Context) error {

		return nil
	},
}

// var initCmd = &cli.Command{
// 	Name:    "init",
// 	Aliases: []string{"i"},
// 	Usage:   "Initialize required folders and config files",
// 	Action: func(cCtx *cli.Context) error {
// 		service := utils.ParentCommandName(cCtx)
// 		repository := settings.getRepository()

// 		envConfig := fmt.Sprintf("%s/%s", conf.Config.DirPath(), service)
// 		repoConfig := fmt.Sprintf("%s/%s/config", repository, service)

// 		switch service {
// 		case s.AUTHELIA:
// 			if err := s.InitAuthelia(repoConfig, envConfig); err != nil {
// 				return err
// 			}
// 		case s.GITEA:
// 			if err := s.InitGitea(envConfig); err != nil {
// 				return err
// 			}
// 		case s.LLDAP:
// 			if err := s.InitLldap(envConfig); err != nil {
// 				return err
// 			}
// 		case s.TRAEFIK:
// 			if err := s.InitTraefik(repoConfig, envConfig); err != nil {
// 				return err
// 			}
// 		}

// 		return nil
// 	},
// }
