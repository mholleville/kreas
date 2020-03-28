package kcli

import (
	"cli/kcli/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func LoadResourcesList(serverUrl string) []string {
	resp, err := http.Get(serverUrl + "/list")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var modules models.List
	json.Unmarshal(body, &modules)

	return modules.Data
}
