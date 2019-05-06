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
				Name:   "status",
				Usage:  "Use it to see Natasha server status",
				Action: handlers.NatashaStatus,
			},
			{
				Name:   "reload",
				Usage:  "Triggers Natasha server rules reload",
				Action: handlers.NatashaReload,
			},
			{
				Name:   "exit",
				Usage:  "Forces Natasha to exit",
				Action: handlers.NatashaExit,
			},
			{
				Name:   "reset-stats",
				Usage:  "Reset all Natasha Stats (App + DPDK)",
				Action: handlers.NatashaResetStats,
			},
			{
				Name:   "version",
				Usage:  "Show Natasha server version",
				Action: handlers.NatashaVersion,
			},
			{
				Name:   "dpdk-stats",
				Usage:  "Show DPDK stats per port",
				Action: handlers.NatashaDPDKStats,
			},
			{
				Name:   "app-stats",
				Usage:  "Display the application per core statistics",
				Action: handlers.NatashaAppStats,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
