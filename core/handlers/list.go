package handlers

import (
	"core/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var ListHandler = func (w http.ResponseWriter, r *http.Request) {
	modules := ResourcesList()

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	response := map[string][]string{"data": modules}
	json.NewEncoder(w).Encode(response)
	return
}

func ResourcesList() []string {
	var modules []string
	files, err := ioutil.ReadDir(utils.RootProjectPath() + "/resources")
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
