package test

import (
	kcli "cli/kcli"
	"os"
	"testing"
)

func TestListCLI(t *testing.T){
	testingArg := "list"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg) // Append a flag
	err := kcli.Run(args)
	ok(t, err)
}

func TestInitCLI(t *testing.T){
	testingArg := "init"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg) // Append a flag
	err := kcli.Run(args)
	ok(t, err)
}