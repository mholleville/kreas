package utils

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"runtime"
	"strings"
)

func RootProjectPath() string {
	_, b, _, _ := runtime.Caller(0)
	basepath   := filepath.Dir(b)
	rootProjectPath := strings.Replace(basepath, "/utils", "", 3)
	return rootProjectPath
}

func Message(message string) (map[string]interface{}) {
	return map[string]interface{} {"data": message}
}

func Respond(w http.ResponseWriter, data map[string] interface{}) {
	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}
