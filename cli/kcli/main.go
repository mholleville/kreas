package kcli

import (
	"fmt"
	"github.com/urfave/cli"
	"strings"
)

var App = cli.NewApp()
var ServerURL = "http://127.0.0.1:8000/api/v0.1"

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
	modules := LoadResourcesList(ServerURL)
	for i := 0; i < len(modules); i++ {
		projectType := modules[i]
		alias := GenerateAlias(projectType)
		arrayCommand = append(arrayCommand, cli.Command{
			Name:  modules[i],
			Usage: "create new " + modules[i] + " project",
			Aliases: []string{alias},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "name",
					Usage: "name of the " + modules[i] + " project",
				},
			},
			Action: func(c *cli.Context) error {
				name := getProjectName(c)
				GenerateProjectArchive(ServerURL, projectType, name)
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
			Usage:   "list of project initialization choices",
			Action: func(c *cli.Context) error {
				modulesArray := LoadResourcesList(ServerURL)
				modulesString := strings.Join(modulesArray,"\n")
				fmt.Println(modulesString)
				return nil
			},
		},
	}
}
