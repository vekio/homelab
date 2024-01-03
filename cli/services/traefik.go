package services

import (
	"os"

	_fs "github.com/vekio/fs"
	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/cli/conf"
	"github.com/vekio/homelab/cli/utils"
)

func InitTraefik() error {
	traefikConf := conf.Config.DirPath() + "/" + TRAEFIK

	err := _fs.CreateDir(traefikConf, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	acmeFile := traefikConf + "/certificates/acme.json"
	if err := _file.Touch(acmeFile, _fs.RestrictedFilePerms); err != nil {
		return err
	}

	err = initTraefikConfig(traefikConf)
	if err != nil {
		return err
	}

	// err = initTraefikSecrets(traefikConf)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func initTraefikConfig(traefikConf string) error {
	src := repoConfig + "/" + TRAEFIK + "/config/dynamic"
	dst := traefikConf + "/config/dynamic"
	if err := _dir.Copy(src, dst); err != nil {
		return err
	}

	data := map[string]string{
		"DOMAIN":             os.Getenv("DOMAIN"),
		"TRAEFIK_CERT_EMAIL": os.Getenv("TRAEFIK_CERT_EMAIL"),
	}

	src = repoConfig + "/" + TRAEFIK + "/config/traefik.yml"
	dst = traefikConf + "/config/traefik.yml"
	if err := utils.ParseConfig(src, dst, data); err != nil {
		return err
	}

	return nil
}

// func initTraefikSecrets(traefikConf string) error {
// 	secretsDir := traefikConf + "/secrets"
// 	err := _fs.CreateDir(secretsDir, os.FileMode(_fs.DefaultDirPerms))
// 	if err != nil {
// 		return err
// 	}

// 	if err = utils.WriteSecret(secretsDir+"/DO_AUTH_TOKEN",
// 		os.Getenv("TRAEFIK_DO_TOKEN")); err != nil {
// 		return err
// 	}

// 	return nil
// }
