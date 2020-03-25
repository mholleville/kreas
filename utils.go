package main

import (
	"github.com/urfave/cli"
	"log"
	"path/filepath"
	"runtime"
)

func rootProjectPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)
	return basepath
}

func getProjectName(c *cli.Context) string {
	name := c.String("name")
	if name == "" {
		log.Fatal("Missing or empty --name flag")
	}
	return name
}
