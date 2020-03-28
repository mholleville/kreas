package kcli

import (
	"github.com/urfave/cli"
	"log"
)

func getProjectName(c *cli.Context) string {
	name := c.String("name")
	if name == "" {
		log.Fatal("Missing or empty --name flag")
	}
	return name
}
