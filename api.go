package brass

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type API struct {
	DataDir string
	SrcDir  string
}

func (api *API) types() []Type {
	path := filepath.Join(api.SrcDir, "types")
	entries, _ := os.ReadDir(path)
	types := []Type{}
	for _, e := range entries {
		id := e.Name()
		path := filepath.Join(path, id)
		b, _ := os.ReadFile(path)
		var t Type
		json.Unmarshal(b, &t)
		types = append(types, t)
	}
	return types
}

func (api *API) tables() []string {
	tables := []string{}
	for _, t := range api.types() {
		tables = append(tables, t.PluralName)
	}
	return tables
}

func (api *API) rows(table string) []string {
	path := filepath.Join(api.DataDir, table)
	entries, _ := os.ReadDir(path)
	rows := []string{}
	for _, e := range entries {
		id := e.Name()
		rows = append(rows, id)
	}
	return rows
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	// List tables
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(api.tables(), "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(b)
	})

	// List rows
	mux.HandleFunc("GET /{table}", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(api.rows(r.PathValue("table")), "", "  ")
		if err != nil {
			panic(err)
		}
		w.Write(b)
	})

	mux.ServeHTTP(w, r)
}
