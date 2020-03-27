package test

import (
	"kreas/kcli"
	"os"
	"testing"
)

func TestInitHelmCLI(t *testing.T){
	testingArg := "init"
	projectType := "helm"
	projectName := "test-helm"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}
