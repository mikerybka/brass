package schemacafe

import (
	"net/http"

	"github.com/mikerybka/brass"
)

type Server struct {
	DataDir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", s.root)
	mux.HandleFunc("/favicon", s.favicon)
	mux.HandleFunc("/login", s.login)
	mux.HandleFunc("/logout", s.logout)
	mux.HandleFunc("/join", s.join)
	mux.HandleFunc("/{orgID}", s.org)
	mux.HandleFunc("/{orgID}/{objID}/", s.obj)
	mux.ServeHTTP(w, r)
}

func (s *Server) auth() *brass.Auth {
	return nil
}
func (s *Server) root(w http.ResponseWriter, r *http.Request)    {}
func (s *Server) favicon(w http.ResponseWriter, r *http.Request) {}
func (s *Server) login(w http.ResponseWriter, r *http.Request)   {}
func (s *Server) logout(w http.ResponseWriter, r *http.Request)  {}
func (s *Server) join(w http.ResponseWriter, r *http.Request)    {}
func (s *Server) org(w http.ResponseWriter, r *http.Request)     {}
func (s *Server) obj(w http.ResponseWriter, r *http.Request)     {}
