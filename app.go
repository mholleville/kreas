package main

import (
	"github.com/otiai10/copy"
	"io/ioutil"
	"log"
	"os"
)

func InitProject(projectType string, projectName string){
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
