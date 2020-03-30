package kcli

import (
	"github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func initProject(projectType string, projectName string){
	dest, _ := os.Getwd()
	err := copy.Copy(rootProjectPath() + "/resources/" + projectType, dest + "/" + projectName)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func ResourcesList() []string {
	var modules []string
	files, err := ioutil.ReadDir(rootProjectPath() + "/resources")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			modules = append(modules, f.Name())
		}
	}

	return modules
}

func GenerateAlias(module string) string {
	aliasMap := map[string]string{
		"terraform": "tf",
		"ansible": "as",
		"terragrunt": "tg",
	}
	splittedModule := strings.SplitN(module, "-", 2)
	if val, ok := aliasMap[splittedModule[0]]; ok {
		suffix := ""
		if len(splittedModule) > 1 {
			suffix = splittedModule[1]
		}
		return val + suffix
	} else {
		return module
	}
}
