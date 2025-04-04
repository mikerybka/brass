package brass

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/util"
)

func NewServer(dir string) http.Handler {
	return &Server{dir}
}

type Server struct {
	Dir string
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serve(s.meta(), s.Dir, w, r)
}

func serve(meta *Metadata, dir string, w http.ResponseWriter, r *http.Request) {
	t := meta.Types[meta.RootType]

	first, rest, isRoot := util.PopPath(r.URL.Path)
	if isRoot {
		if t.IsMap || t.IsArray || t.IsStruct {
			if r.Method == http.MethodGet {
				res, err := util.ReadDir(dir)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				json.NewEncoder(w).Encode(res)
			} else {
				http.NotFound(w, r)
				return
			}
		} else {
			if r.Method == http.MethodGet {
				b, err := os.ReadFile(dir)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				w.Write(b)
			} else if r.Method == http.MethodPut {
				b, err := io.ReadAll(r.Body)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				err = util.WriteFile(dir, b)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else if r.Method == http.MethodDelete {
				err := os.RemoveAll(dir)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			} else {
				http.NotFound(w, r)
				return
			}
		}
		return
	}

	if t.IsArray || t.IsMap {
		meta.RootType = t.UnderlyingTypeID
		dir = filepath.Join(dir, first)
		r.URL.Path = rest
		serve(meta, dir, w, r)
		return
	}

	if t.IsStruct {
		for _, f := range t.Fields {
			if f.Name.SnakeCase() == first {
				meta.RootType = f.TypeID
				dir = filepath.Join(dir, first)
				r.URL.Path = rest
				serve(meta, dir, w, r)
				return
			}
		}
		http.NotFound(w, r)
		return
	}

	http.NotFound(w, r)
}

func (s *Server) meta() *Metadata {
	m := &Metadata{}
	json.Unmarshal(DefaultMeta, m)
	util.ReadJSONFile(filepath.Join(s.Dir, "meta"), m)
	return m
}
