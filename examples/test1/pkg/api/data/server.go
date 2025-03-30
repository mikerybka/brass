package data

import "net/http"

type Server struct {
	Workdir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
