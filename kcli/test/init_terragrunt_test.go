package test

import (
	"kreas/kcli"
	"os"
	"testing"
)

func TestInitTerragruntProjectCLI(t *testing.T){
	testingArg := "init"
	projectType := "terragrunt-project"
	projectName := "test-terragrunt-project"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}

func TestInitTerragruntProjectWithAliasCLI(t *testing.T){
	testingArg := "init"
	projectType := "tgproject"
	projectName := "test-tgproject"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}
