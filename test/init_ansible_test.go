package test

import (
	"kreas/kcli"
	"os"
	"testing"
)

func TestInitAnsibleProjectCLI(t *testing.T){
	testingArg := "init"
	projectType := "ansible-project"
	projectName := "test-ansible-project"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}

func TestInitAnsibleRoleCLI(t *testing.T){
	testingArg := "init"
	projectType := "ansible-role"
	projectName := "test-ansible-role"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg, projectType, "--name=" + projectName) // Append a flag
	err := kcli.Run(args)
	dest, _ := os.Getwd()
	os.RemoveAll(dest + "/" + projectName)
	ok(t, err)
}
