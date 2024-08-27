package main

import (
	"fmt"
	"os"

	"github.com/vekio/homelab/pkg/cmd/root"
)

func main() {
	rootCmd, _ := root.NewCmdRoot()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
