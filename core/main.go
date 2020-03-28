package main

import (
	"core/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func main() {
	basePath := "/api/v0.1"
	var router = mux.NewRouter().StrictSlash(true)
	router.HandleFunc(basePath + "/list", handlers.ListHandler).Methods("GET")
	router.HandleFunc(basePath + "/init", handlers.ZipHandler).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}
