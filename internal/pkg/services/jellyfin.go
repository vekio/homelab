package services

import "github.com/vekio/homelab/internal/pkg/context"

var jellyfin = &Service{
	Name:        JELLYFIN,
	ComposeFile: composeFile,
	Context:     context.DEFAULT,
	Init: func() error {
		return nil
	},
}
