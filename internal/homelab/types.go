package homelab

import (
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/secrets"
	"github.com/vekio/homelab/internal/services"
)

type Homelab struct {
	Config   config.HomelabConfig
	Secrets  secrets.HomelabSecrets
	Services services.Services
}
