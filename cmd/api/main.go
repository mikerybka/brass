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
	api := &brass.API{
		DataDir: requireEnvVar("SRC_DIR"),
		SrcDir:  requireEnvVar("DATA_DIR"),
	}
	log.Fatal(http.ListenAndServe(":"+requireEnvVar("PORT"), api))
}
