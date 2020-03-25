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
					Usage: "create a new helm project",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "name of the helm project",
						},
					},
					Action: func(c *cli.Context) error {
						name := getProjectName(c)
						InitProject("helm", name)
						return nil
					},
				},
				{
					Name:  "terraform-project",
					Aliases: []string{"tfp"},
					Usage: "create a new terraform project",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "name of the terraform project",
						},
					},
					Action: func(c *cli.Context) error {
						name := getProjectName(c)
						InitProject("terraform-project", name)
						return nil
					},
				},
				{
					Name:  "terraform-module",
					Aliases: []string{"tfm"},
					Usage: "create a new terraform module",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "name of the terraform module",
						},
					},
					Action: func(c *cli.Context) error {
						name := getProjectName(c)
						InitProject("terraform-module", name)
						return nil
					},
				},
				{
					Name:  "ansible-project",
					Aliases: []string{"asp"},
					Usage: "create a new ansible project",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "name of the ansible project",
						},
					},
					Action: func(c *cli.Context) error {
						name := getProjectName(c)
						InitProject("ansible-project", name)
						return nil
					},
				},
				{
					Name:  "ansible-role",
					Aliases: []string{"asr"},
					Usage: "create a new ansible role",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name",
							Usage: "name of the ansible role",
						},
					},
					Action: func(c *cli.Context) error {
						name := getProjectName(c)
						InitProject("ansible-role", name)
						return nil
					},
				},
			},
		},
	}
}
