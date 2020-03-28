package test

import (
	"kreas/kcli"
	"testing"
)

// GenerateAlias(module string) string //
func TestTFAlias(t *testing.T){
	result := kcli.GenerateAlias("terraform-project")
	equals(t, "tfproject", result)
}

func TestASAlias(t *testing.T){
	result := kcli.GenerateAlias("ansible-role")
	equals(t, "asrole", result)
}

func TestHelmAlias(t *testing.T){
	result := kcli.GenerateAlias("helm")
	equals(t, "helm", result)
}

func TestUndefinedAlias(t *testing.T){
	result := kcli.GenerateAlias("xseop")
	equals(t, "xseop", result)
}
