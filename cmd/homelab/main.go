package main

import (
	"log"

	cmdHomelab "github.com/vekio/homelab/internal/cmd/homelab"
	"github.com/vekio/homelab/internal/homelab"
)

func main() {
	homelab, err := homelab.NewHomelab()
	if err != nil {
		log.Fatalf("error homelab: %v", err)
	}

	// Homelab root command.
	rootCmd := cmdHomelab.NewCmdHomelab(homelab)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("error homelab: %v", err)
	}
}
