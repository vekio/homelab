package homelab

import (
	"github.com/joho/godotenv"
	"github.com/vekio/homelab/internal/config"
	"github.com/vekio/homelab/internal/secrets"
	"github.com/vekio/homelab/internal/services"
)

func New() (*Homelab, error) {
	// Create HomelabConfig
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	// Create HomelabSecrets
	s, err := secrets.New()
	if err != nil {
		return nil, err
	}

	// Create HomelabServices
	srvs := services.New(conf)
	if err != nil {
		return nil, err
	}

	h := Homelab{
		Config:   conf,
		Secrets:  s,
		Services: srvs,
	}

	// Init
	if _, ok := srvs[services.AUTHELIA]; ok {
		srvs[services.AUTHELIA].Init = h.initAuthelia
	}
	if _, ok := srvs[services.GITEA]; ok {
		srvs[services.GITEA].Init = h.initGitea
	}
	if _, ok := srvs[services.IMMICH]; ok {
		srvs[services.IMMICH].Init = h.initImmich
	}
	if _, ok := srvs[services.JELLYFIN]; ok {
		srvs[services.JELLYFIN].Init = h.initJellyfin
	}
	if _, ok := srvs[services.LLDAP]; ok {
		srvs[services.LLDAP].Init = h.initLldap
	}
	if _, ok := srvs[services.PROTONMAIL_BRIDGE]; ok {
		srvs[services.PROTONMAIL_BRIDGE].Init = h.initProtonmailBridge
	}
	if _, ok := srvs[services.TRAEFIK]; ok {
		srvs[services.TRAEFIK].Init = h.initTraefik
	}

	// Load envvars
	if err := godotenv.Load(h.Config.EnvFile); err != nil {
		return nil, err
	}

	return &h, nil
}
