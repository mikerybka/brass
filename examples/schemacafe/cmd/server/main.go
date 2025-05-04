package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mikerybka/schemacafe"
	"github.com/mikerybka/util"
)

func main() {
	dataDir := util.RequireEnvVar("DATA_DIR")
	addr := util.RequireEnvVar("LISTEN_ADDR")
	s := &schemacafe.Server{
		DataDir: dataDir,
	}
	err := http.ListenAndServe(addr, s)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
