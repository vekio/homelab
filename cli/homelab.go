package homelab

import (
	"log"

	"github.com/vekio/homelab/cli/conf"
	"github.com/vekio/homelab/cli/secrets"
)

func init() {
	filename := conf.Config.DirPath() + "/secrets.yml"
	if err := secrets.SoftInitSecrets(filename); err != nil {
		log.Fatalf("failed to init secrets %s", err)
	}
}
