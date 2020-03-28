package test

import (
	"core/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var basePath = "/api/v0.1"

func TestInitHandler(t *testing.T) {
	tt := []struct{
		projectType string
		projectName string
	}{
		{"helm", "helm-projectName"},
		{"ansible-project", "ansible-projectName"},
		{"ansible-role", "ansible-projectName"},
		{"terraform-module", "terraform-moduleName"},
		{"terraform-project", "terraform-projectName"},
	}

	for _, tc := range tt {
		path := fmt.Sprintf(basePath + "/init?type=%s&name=%s", tc.projectType, tc.projectName)
		req, err := http.NewRequest("GET", path, nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		// Need to create a router that we can pass the request through so that the vars will be added to the context
		router := mux.NewRouter()
		router.HandleFunc(basePath + "/init", handlers.ZipHandler)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			fmt.Println(path)
			t.Errorf("handler should have failed on projectType %s and projectName %s: got %v want %v",
				tc.projectType, tc.projectName, rr.Code, http.StatusOK)
		}
	}
}

func TestEmptyParamsInitHandler(t *testing.T) {
	path := fmt.Sprintf(basePath + "/init?type=&name=")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(basePath + "/init", handlers.ZipHandler)
	router.ServeHTTP(rr, req)

	equals(t, rr.Code, 400)
}

func TestInvalidParamsInitHandler(t *testing.T) {
	var projectType = "twop"
	path := fmt.Sprintf(basePath + "/init?type=%s&name=test", projectType)
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(basePath + "/init", handlers.ZipHandler)
	router.ServeHTTP(rr, req)

	equals(t, rr.Code, 500)
}
