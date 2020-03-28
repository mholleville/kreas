package test

import (
	"core/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListHandler(t *testing.T) {
	path := fmt.Sprintf(basePath + "/list")
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc(basePath + "/list", handlers.ListHandler)
	router.ServeHTTP(rr, req)

	equals(t, rr.Code, 202)
}
