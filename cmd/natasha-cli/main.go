package main

import (
	"os"

	"github.com/kaminek/natasha-cli/pkg/client"
	// "github.com/kaminek/natasha-cli/pkg/handlers"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:                  "natasha-cli",
		EnableShellCompletion: true,
		Authors: []*cli.Author{
			{
				Name:  "Amine KHERBOUCHE",
				Email: "akherbouche@scaleway.com",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "address",
				Value: "127.0.0.1",
				Usage: "Server address",
			},
			&cli.StringFlag{
				Name:  "port",
				Value: "4242",
				Usage: "Server port",
			},
		},
		Commands: []*cli.Command{
			{
				Name:        "cli",
				Usage:       "use it to see natasha server status",
				Description: "Checks natasha server status",
				// Action:      handlers.NatashaStatusHandler,
				Action: client.Connect,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
