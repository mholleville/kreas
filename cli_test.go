package main

import (
	"os"
	"testing"
)

func TestListCLI(t *testing.T){
	testingArg := "list"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg) // Append a flag
	err := Run(args)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Arg " + testingArg + " ok")
	}
}

func TestInitCLI(t *testing.T){
	testingArg := "init"
	args := os.Args[0:1] // Name of the program.
	args = append(args, testingArg) // Append a flag
	err := Run(args)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Log("Arg " + testingArg + " ok")
	}
}
