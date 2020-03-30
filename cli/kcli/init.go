package kcli

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func GenerateProjectArchive(serverUrl string, projectType string, projectName string) {
	path := fmt.Sprintf(serverUrl+"/init?type=%s&name=%s", projectType, projectName)
	resp, err := http.Get(path)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))

	if err != nil {
		fmt.Printf("err: %s", err)
	}

	for _, zipFile := range zipReader.File {
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Reading file:", zipFile.Name)
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}

		file := unzippedFileBytes

		dest, _ := os.Getwd()
		target := dest + "/"

		createProjectDirectory(projectName, zipFile, target)

		ioutil.WriteFile(target + strings.Replace(zipFile.Name, projectType, projectName, 1), file, 0644)
	}
}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func createProjectDirectory(projectName string, zipFile *zip.File, target string) {
	res1 := strings.Split(strings.TrimPrefix(zipFile.Name, "/"), "/")
	if len(res1) > 0 {
		res1 = res1[:len(res1)-1]
	}
	res1[0] = projectName
	newPath := filepath.Join(target, strings.Join(res1, "/"))
	os.MkdirAll(newPath, os.ModePerm)

}

