package kcli

import (
	"fmt"
	"github.com/urfave/cli"
	"strings"
)

var App = cli.NewApp()

func Run(args []string) error {
	info()
	commands()
	err := App.Run(args)
	if err != nil {
		return err
	}
	return nil

}

func info() {
	App.Name = "** Kreas **"
	App.Usage = "Kreas is an open-source tool to help you organize your DevOps project"
	App.Author = "mholleville"
	App.Version = "0.0.1"
}

func loadSubCommands() []cli.Command {
	var arrayCommand []cli.Command
	modules := ResourcesList()
	for i := 0; i < len(modules); i++ {
		projectType := modules[i]
		arrayCommand = append(arrayCommand, cli.Command{
			Name:  modules[i],
			Usage: "create new " + modules[i] + " project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "name of the " + modules[i] + " project",
				},
			},
			Action: func(c *cli.Context) error {
				name := getProjectName(c)
				InitProject(projectType, name)
				return nil
			},
		})


	}

	return arrayCommand
}

func commands() {
	App.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "options for project init",
			Subcommands: loadSubCommands(),
		},
		{
			Name:    "list",
			Usage:   "options for project init",
			Action: func(c *cli.Context) error {
				modulesArray := ResourcesList()
				modulesString := strings.Join(modulesArray,"\n")
				fmt.Println(modulesString)
				return nil
			},
		},
	}
}
