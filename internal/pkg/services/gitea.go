package services

import "github.com/vekio/homelab/internal/pkg/context"

var gitea = &Service{
	Name:        GITEA,
	ComposeFile: composeFile,
	Context:     context.DEFAULT,
	Init: func() error {
		return nil
	},
}
