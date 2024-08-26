package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/vekio/homelab/pkg/homelab/cmd"
)

func main() {
	app := &cli.App{
		Name:  "homelab",
		Usage: "Manage my homelab services",
		Commands: []*cli.Command{
			cmd.ConfigCmd,
		},
		// Action: func(*cli.Context) error {
		// 	fmt.Println("homelab 😊!")
		// 	return nil
		// },
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
