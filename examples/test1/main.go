package main

import (
	"net/http"

	"github.com/mikerybka/brass/examples/test1/pkg/app"
	"github.com/mikerybka/util"
)

func main() {
	port := util.RequireEnvVar("PORT")
	addr := ":" + port
	s := &app.Server{}
	http.ListenAndServe(addr, s)
}
