package services

import "github.com/vekio/homelab/internal/pkg/context"

var traefik = &Service{
	Name:        TRAEFIK,
	ComposeFile: composeFile,
	Context:     context.DEFAULT,
	Init: func() error {
		return nil
	},
}
