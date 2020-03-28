package main

import (
	"cli/kcli"
	"log"
	"os"
)

func main() {
	err := kcli.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}

}
