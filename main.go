package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mikerybka/brass"
)

func requireEnvVar(name string) string {
	v := os.Getenv(name)
	if v == "" {
		fmt.Println(name, "required")
		os.Exit(1)
	}
	return v
}

func main() {
	app := &brass.API{
		DataDir: requireEnvVar("DATA_DIR"),
		SrcDir:  requireEnvVar("SRC_DIR"),
	}
	addr := ":" + requireEnvVar("PORT")
	err := http.ListenAndServe(addr, app)
	if err != nil {
		log.Fatal(err)
	}
}
