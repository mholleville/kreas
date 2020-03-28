package kcli

import (
	"github.com/urfave/cli"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func rootProjectPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)
	rootProjectPath := strings.Replace(basepath, "/kcli", "", 3)
	return rootProjectPath
}

func getProjectName(c *cli.Context) string {
	name := c.String("name")
	if name == "" {
		log.Fatal("Missing or empty --name flag")
	}
	return name
}
