package brass

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/util"
)

type List struct {
	Path string
	Dir  string
}

func (l *List) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		entries, err := os.ReadDir(l.Dir)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "<ul>")
		for _, e := range entries {
			fmt.Fprintf(w, "<li><a href=\"%s\">%s</a></lil>", filepath.Join(l.Path, e.Name()), e.Name())
		}
		fmt.Fprintf(w, "<li><a href=\"%s/new\">New</a></lil></ul>", l.Path)
	})
	mux.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		f := &util.SimpleForm{
			TitleText: "New",
			Fields: []util.Field{
				{
					Name: util.NewName("ID"),
				},
			},
			SubmitText: "Create",
		}
		f.HandlePOST = func(w http.ResponseWriter, r *http.Request) {
			id := r.FormValue("id")
			path := filepath.Join(l.Dir, id)
			_, err := os.Stat(path)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					err = os.WriteFile(path, []byte("{}"), os.ModePerm)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					return
				}
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			f.Error = fmt.Errorf("ID %s already in use", id)
			f.ServeHTTP(w, r)
			return
		}
		f.ServeHTTP(w, r)
	})
	mux.HandleFunc("PUT /{objID}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("GET /{objID}/", func(w http.ResponseWriter, r *http.Request) {})
	mux.ServeHTTP(w, r)
}
