package services

import "github.com/vekio/homelab/internal/pkg/context"

var immich = &Service{
	Name:        IMMICH,
	ComposeFile: composeFile,
	Context:     context.DEFAULT,
	Init: func() error {
		return nil
	},
}
