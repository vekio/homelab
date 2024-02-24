package services

import (
	"fmt"

	"github.com/vekio/homelab/internal/config"
)

func New(c config.HomelabConfig) Services {
	srvs := make(Services)

	for srvName, srvConf := range c.Services {
		srv, ok := ServicesMap[srvName]
		if ok {

			srv.ComposeFile = fmt.Sprintf("%s/%s/compose.yml", c.Repository, srvName)
			srv.TraefikComposeFile = fmt.Sprintf("%s/%s/compose.traefik.yml", c.Repository, srvName)

			context := "default"
			if srvConf.Context != "" {
				context = srvConf.Context
			}

			srv.Context = context
			srv.AllInitCmd = srvConf.AllInit
			srv.AllUpCmd = srvConf.AllUp
			srv.AllDownCmd = srvConf.AllDown

			srv.Init = func() error {
				return nil
			}

			srvs[srvName] = srv
		}

	}

	return srvs
}

func (srvs Services) AllDown() {
	for _, srv := range srvs {
		if srv.AllDownCmd {
			srv.Down()
		}
	}
}

func (srvs Services) AllUp() {
	for _, srv := range srvs {
		if srv.AllUpCmd {
			srv.Down()
		}
	}
}

func (srvs Services) AllInit() {
	for _, srv := range srvs {
		if srv.AllInitCmd {
			srv.Down()
		}
	}
}
