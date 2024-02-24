package homelab

import (
	"os"

	_fs "github.com/vekio/fs"
	_dir "github.com/vekio/fs/dir"
	_file "github.com/vekio/fs/file"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/secrets"
	"github.com/vekio/homelab/internal/services"
)

func (h Homelab) initTraefik() (err error) {
	err = initTraefikConfig(h.Config, h.Secrets)
	if err != nil {
		return err
	}

	// TODO investigate DO token as secret

	return nil
}

func initTraefikConfig(c config.HomelabConfig, s secrets.HomelabSecrets) (err error) {
	volumeDir := config.Manager.DirPath() + "/" + services.TRAEFIK

	err = _fs.CreateDir(volumeDir, os.FileMode(_fs.DefaultDirPerms))
	if err != nil {
		return err
	}

	acmeFile := volumeDir + "/certificates/acme.json"
	if err = _file.Touch(acmeFile, _fs.RestrictedFilePerms); err != nil {
		return err
	}

	src := c.Repository + "/" + services.TRAEFIK + "/config/dynamic"
	dst := volumeDir + "/config/dynamic"
	if err := _dir.Copy(src, dst); err != nil {
		return err
	}

	traefikData := map[string]string{
		"DOMAIN":                  os.Getenv("DOMAIN"),
		"TRAEFIK_CERT_EMAIL":      os.Getenv("TRAEFIK_CERT_EMAIL"),
		"TRAEFIK_LETSENCRYPT_API": os.Getenv("TRAEFIK_LETSENCRYPT_API"),
	}

	src = c.Repository + "/" + services.TRAEFIK + "/config/traefik.yml"
	dst = volumeDir + "/config/traefik.yml"
	if err := ParseTemplate(src, dst, traefikData); err != nil {
		return err
	}

	unraidData := map[string]string{
		"DOMAIN": os.Getenv("DOMAIN"),
		"UNRAID": os.Getenv("UNRAID"),
	}

	src = c.Repository + "/" + services.TRAEFIK + "/config/dynamic/unraid.yml"
	dst = volumeDir + "/config/dynamic/unraid.yml"
	if err := ParseTemplate(src, dst, unraidData); err != nil {
		return err
	}

	piholeData := map[string]string{
		"DOMAIN": os.Getenv("DOMAIN"),
		"PIHOLE": os.Getenv("PIHOLE"),
	}

	src = c.Repository + "/" + services.TRAEFIK + "/config/dynamic/pihole.yml"
	dst = volumeDir + "/config/dynamic/pihole.yml"
	if err := ParseTemplate(src, dst, piholeData); err != nil {
		return err
	}

	return nil
}
