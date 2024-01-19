package services

import "github.com/vekio/homelab/internal/pkg/context"

var lldap = &Service{
	Name:        LLDAP,
	ComposeFile: composeFile,
	Context:     context.DEFAULT,
	Init: func() error {
		return nil
	},
}
