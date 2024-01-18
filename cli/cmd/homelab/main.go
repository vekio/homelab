package main

import (
	"log"
	"os"

	homelab "github.com/vekio/homelab/cli/internal/homelab"
)

func main() {
	if err := homelab.Cmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
