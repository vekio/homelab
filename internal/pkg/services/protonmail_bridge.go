package services

import "github.com/vekio/homelab/internal/pkg/context"

var protonmailBridge = &Service{
	Name:        PROTONMAIL_BRIDGE,
	ComposeFile: composeFile,
	Context:     context.DEFAULT,
	Init: func() error {
		return nil
	},
}
