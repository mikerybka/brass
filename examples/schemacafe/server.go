package schemacafe

import (
	"net/http"
	"path/filepath"

	"github.com/mikerybka/brass"
	"github.com/mikerybka/util"
)

type Server struct {
	DataDir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", s.landingPage)
	mux.HandleFunc("POST /{$}", s.createOrg)
	mux.HandleFunc("GET /favicon.ico", s.favicon)
	mux.HandleFunc("/login", s.login)
	mux.HandleFunc("/logout", s.logout)
	mux.HandleFunc("/join", s.join)
	mux.HandleFunc("/{orgID}", s.org)
	mux.HandleFunc("/{orgID}/{objID}/", s.obj)
	mux.ServeHTTP(w, r)
}

func (s *Server) auth() *brass.Auth {
	return &brass.Auth{
		DataDir: filepath.Join(s.DataDir, "auth"),
	}
}
func (s *Server) landingPage(w http.ResponseWriter, r *http.Request) {}
func (s *Server) createOrg(w http.ResponseWriter, r *http.Request) {
}
func (s *Server) favicon(w http.ResponseWriter, r *http.Request) {}
func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	form := &util.SimpleForm{
		TitleText: "Join",
		Fields: []util.Field{
			{
				Name: util.NewName("Username"),
			},
			{
				Name: util.NewName("Password"),
				Type: "password",
			},
		},
		SubmitText: "Enter",
	}
	form.HandlePOST = func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		token, err := s.auth().Login(
			username,
			r.FormValue("password"),
		)
		if err != nil {
			form.Error = err
			form.ServeHTTP(w, r)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "UserID",
			Value: username,
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "SessionID",
			Value: token,
		})
		http.Redirect(w, r, "/"+username, http.StatusSeeOther)
	}
	form.ServeHTTP(w, r)
}
func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	form := &util.SimpleForm{
		SubmitText: "Logout",
	}
	form.HandlePOST = func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:  "UserID",
			Value: "",
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "SessionID",
			Value: "",
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	form.ServeHTTP(w, r)
}
func (s *Server) join(w http.ResponseWriter, r *http.Request) {
	form := &util.SimpleForm{
		TitleText: "Join",
		Fields: []util.Field{
			{
				Name: util.NewName("Username"),
			},
			{
				Name: util.NewName("Password"),
				Type: "password",
			},
			{
				Name: util.NewName("Confirm Password"),
				Type: "password",
			},
		},
		SubmitText: "Enter",
	}
	form.HandlePOST = func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		token, err := s.auth().Join(
			username,
			r.FormValue("password"),
			r.FormValue("confirm-password"),
		)
		if err != nil {
			form.Error = err
			form.ServeHTTP(w, r)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:  "UserID",
			Value: username,
		})
		http.SetCookie(w, &http.Cookie{
			Name:  "SessionID",
			Value: token,
		})
		http.Redirect(w, r, "/"+username, http.StatusSeeOther)
	}
	form.ServeHTTP(w, r)
}
func (s *Server) org(w http.ResponseWriter, r *http.Request) {
	ok, err := s.auth().Allowed(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
}
func (s *Server) obj(w http.ResponseWriter, r *http.Request) {
	ok, err := s.auth().Allowed(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !ok {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
}
