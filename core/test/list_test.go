package test

import (
	"core/handlers"
	"testing"
)

// func ResourcesList() []string //
func TestModuleArrayLen(t *testing.T){
	modulesArray := handlers.ResourcesList()
	equals(t, len(modulesArray), 5)
}
