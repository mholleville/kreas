package test

import (
	"kreas/kcli"
	"testing"
)

// func ResourcesList() []string //
func TestModuleArrayLen(t *testing.T){
	modulesArray := kcli.ResourcesList()
	equals(t, len(modulesArray), 6)
}

// GenerateAlias(module string) string //
func TestTFProjectAlias(t *testing.T){
	result := kcli.GenerateAlias("terraform-project")
	equals(t, "tfproject", result)
}

func TestTFAlias(t *testing.T){
	result := kcli.GenerateAlias("terraform")
	equals(t, "tf", result)
}

func TestASRoleAlias(t *testing.T){
	result := kcli.GenerateAlias("ansible-role")
	equals(t, "asrole", result)
}

func TestASAlias(t *testing.T){
	result := kcli.GenerateAlias("ansible")
	equals(t, "as", result)
}

func TestHelmAlias(t *testing.T){
	result := kcli.GenerateAlias("helm")
	equals(t, "helm", result)
}

func TestTGProjectAlias(t *testing.T){
	result := kcli.GenerateAlias("terragrunt-project")
	equals(t, "tgproject", result)
}

func TestTGAlias(t *testing.T){
	result := kcli.GenerateAlias("terragrunt")
	equals(t, "tg", result)
}

func TestUndefinedAlias(t *testing.T){
	result := kcli.GenerateAlias("xseop")
	equals(t, "xseop", result)
}
