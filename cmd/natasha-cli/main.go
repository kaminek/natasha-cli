package main

import (
	"os"

	"github.com/kaminek/natasha-cli/pkg/handlers"
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
				Name:        "status",
				Usage:       "use it to see natasha server status",
				Description: "Checks natasha server status",
				Action:      handlers.NatashaStatus,
			},
			{
				Name:        "reload",
				Usage:       "use it to see natasha server status",
				Description: "Checks natasha server status",
				Action:      handlers.NatashaReload,
			},
			{
				Name:        "exit",
				Usage:       "use it to see natasha server status",
				Description: "Checks natasha server status",
				Action:      handlers.NatashaExit,
			},
			{
				Name:        "reset-stats",
				Usage:       "use it to see natasha server status",
				Description: "Checks natasha server status",
				Action:      handlers.NatashaResetStats,
			},
			{
				Name:        "version",
				Usage:       "use it to see natasha server status",
				Description: "Checks natasha server status",
				Action:      handlers.NatashaVersion,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
