package main

import (
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func info() {
	app.Name = "** Kreas **"
	app.Usage = "Kreas is an open-source tool to help you organize your DevOps project"
	app.Author = "mholleville"
	app.Version = "0.0.1"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "options for project init",
			Subcommands: []cli.Command{
				{
					Name:  "helm",
					Usage: "create a new app project",
					Action: func(c *cli.Context) error {
						InitProject("helm")
						return nil
					},
				},
				{
					Name:  "terraform",
					Usage: "create a new terraform project",
					Action: func(c *cli.Context) error {
						InitProject("terraform")
						return nil
					},
				},
				{
					Name:  "ansible",
					Usage: "create a new ansible project",
					Action: func(c *cli.Context) error {
						InitProject("ansible")
						return nil
					},
				},
			},
		},
	}
}
