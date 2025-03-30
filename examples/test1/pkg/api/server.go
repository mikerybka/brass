package api

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/mikerybka/brass/examples/test1/pkg/api/auth"
	"github.com/mikerybka/brass/examples/test1/pkg/api/data"
)

type Server struct {
	Workdir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/auth") {
		http.StripPrefix("/auth", &auth.Server{
			Workdir: filepath.Join(s.Workdir, "auth"),
		})
	}
	if strings.HasPrefix(r.URL.Path, "/data") {
		http.StripPrefix("/data", &data.Server{
			Workdir: filepath.Join(s.Workdir, "data"),
		})
	}
}
