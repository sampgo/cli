package main

import (
	"log"
	"os"
	"sampgo-cli/handler"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		EnableBashCompletion: true,
		Name:                 "sampgo",
		Usage:                "Interacts with the SA-MP server via Go",
		Commands: []*cli.Command{
			{
				Name:   "init",
				Usage:  "Create a new gomode",
				Action: handler.Init,
			},
			{
				Name:  "build",
				Usage: "Builds your gomode",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "verbose",
						Aliases: []string{"v"},
						Usage:   "Enable verbose mode for enhanced debugging.",
					},
				},
				Action: handler.Build,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
