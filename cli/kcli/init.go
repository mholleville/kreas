package kcli

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GenerateProjectArchive(serverUrl string, projectType string, projectName string) {
	path := fmt.Sprintf(serverUrl+"/init?type=%s&name=%s", projectType, projectName)
	//dest, _ := os.Getwd()
	//target := dest + "/" + projectName
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
		if zipFile.FileInfo().IsDir() {
			fmt.Println(zipFile.Mode())
			//os.MkdirAll(target + zipFile.Name, 0777)
			os.MkdirAll("subdir/parent/child", 0755)
			continue
		}
		unzippedFileBytes, err := readZipFile(zipFile)
		if err != nil {
			log.Println(err)
			continue
		}

		_ = unzippedFileBytes // this is unzipped file bytes
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

