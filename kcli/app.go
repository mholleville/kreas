package kcli

import (
	"bytes"
	"encoding/json"
	"github.com/otiai10/copy"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func initProject(projectType string, projectName string) {
	dest, _ := os.Getwd()
	err := copy.Copy(rootProjectPath()+"/resources/"+projectType, dest+"/"+projectName)
	if err != nil {
		log.Fatal(err.Error())
	}
	applyTemplate(dest + "/" + projectName)
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
		"ansible":   "as",
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

func applyTemplate(rootDir string) {
	var files []string
	type Profile struct {
		Package         string
		ApplicationName string
		Version         string
		FirstClass      string
	}
	jsonProfile, err := ioutil.ReadFile("profile.json")
	if err != nil {
		log.Fatalln(err)
	}
	var profile Profile
	json.Unmarshal([]byte(jsonProfile), &profile)
	filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		file, err := os.Stat(path)
		if !file.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	for _, f := range files {
		tpl, err := template.ParseFiles(f)
		var output bytes.Buffer
		err = tpl.Execute(&output, profile)
		ioutil.WriteFile(f, output.Bytes(), 0644)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
