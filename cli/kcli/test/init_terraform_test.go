package test

import (
	kcli "cli/kcli"
	"os"
	"testing"
)

func TestInitTerraformProjectCLI(t *testing.T){
	testingArg := "init"
	projectType := "terraform-project"
	projectName := "test-terraform-project"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}

func TestInitTerraformProjectWithAliasCLI(t *testing.T){
	testingArg := "init"
	projectType := "tfproject"
	projectName := "test-tfproject"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}

func TestInitTerraformModuleCLI(t *testing.T){
	testingArg := "init"
	projectType := "terraform-module"
	projectName := "test-terraform-module"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}

func TestInitTerraformModuleWithAliasCLI(t *testing.T){
	testingArg := "init"
	projectType := "tfmodule"
	projectName := "test-terraform-module"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}

