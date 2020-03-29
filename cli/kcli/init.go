package kcli

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func GenerateProjectArchive(serverUrl string, projectType string, projectName string){
	path := fmt.Sprintf(serverUrl + "/init?type=%s&name=%s", projectType, projectName)
	dest, _ := os.Getwd()
	target := dest + "/" + projectName
	tempZip := target + ".zip"
	resp, err := http.Get(path)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("status", resp.Status)
	if resp.StatusCode != 200 {
		return
	}

	//zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))

	// Create the file
	out, err := os.Create(tempZip)
	if err != nil {
		fmt.Printf("err: %s", err)
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	err = unzip(tempZip, target)
	fmt.Println(err.Error())

	//for _, f := range zipReader.File {
	//	bodyString := string(body)
	//	d1 := []byte(bodyString)
	//	fmt.Println(f.Name)
	//	ioutil.WriteFile(target + f.Name, d1, 0777)
	//	if err != nil {
	//		log.Fatal(err.Error())
	//	}
		////fmt.Println("Reading file:", zipFile.Name)
		//fmt.Println(target + zipFile.Name)
		//io.Copy(target, resp.Body)
		//
		//os.MkdirAll(target + "/subdir/parent/child", 0755)
		//if zipFile.FileInfo().IsDir() {
		//	fmt.Println(zipFile.Mode())
		//	//os.MkdirAll(target + zipFile.Name, 0777)
		//	os.MkdirAll("subdir/parent/child", 0755)
		//	continue
		//}
		//unzippedFileBytes, err := readZipFile(zipFile)
		//if err != nil {
		//	log.Println(err)
		//	continue
		//}
		//
		//_ = unzippedFileBytes // this is unzipped file bytes
	}

func readZipFile(zf *zip.File) ([]byte, error) {
	f, err := zf.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}

func unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0777); err != nil {
		return err
	}

	for _, file := range reader.File {
		fmt.Println(file.Name)
		path := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			fmt.Println("1")
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
