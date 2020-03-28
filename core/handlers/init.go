package handlers

import (
	"archive/zip"
	"bytes"
	"core/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var ZipHandler = func (w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	projectType := params.Get("type")
	projectName := params.Get("name")

	if projectType == "" || projectName == "" {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(utils.Error(400, "Bad request. Missing parameters."))
		return
	}

	var pathToZip = utils.RootProjectPath() + "/resources/" + projectType
	buf := new(bytes.Buffer)
	myZip := zip.NewWriter(buf)
	err := filepath.Walk(pathToZip, func(filePath string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if err != nil {
			return err
		}
		relPath := strings.TrimPrefix(filePath, filepath.Dir(pathToZip))
		zipFile, err := myZip.Create(relPath)
		if err != nil {
			return err
		}
		fsFile, err := os.Open(filePath)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Error(404, err.Error()))
		return
	}
	err = myZip.Close()
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(utils.Error(500, err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", projectName))
	//io.Copy(w, buf)
	w.Write(buf.Bytes())
}
