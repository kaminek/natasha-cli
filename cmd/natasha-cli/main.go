package main

import (
	"os"

	"github.com/kaminek/natasha-cli/pkg/handlers"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:                  "greet",
		EnableShellCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Amine KHERBOUCHE",
				Email: "akherbouche@scaleway.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "status",
				Usage:       "use it to see natasha server status",
				Description: "Checks natasha server status",
				Action:      handlers.NatashaStatusHandler,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
