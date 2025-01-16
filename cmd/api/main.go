package main

import (
	"fmt"
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
	http.ListenAndServe(":3000", api)
}
