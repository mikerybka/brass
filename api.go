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

func (api *API) types() []*Type {
	path := filepath.Join(api.SrcDir, "types")
	entries, _ := os.ReadDir(path)
	types := []*Type{}
	for _, e := range entries {
		id := e.Name()
		types = append(types, api.typ(id))
	}
	return types
}

func (api *API) typ(id string) *Type {
	t := &Type{}
	path := filepath.Join(api.SrcDir, "types", id)
	b, _ := os.ReadFile(path)
	json.Unmarshal(b, t)
	return t
}

func (api *API) tables() []string {
	path := filepath.Join(api.SrcDir, "types")
	entries, _ := os.ReadDir(path)
	tables := []string{}
	for _, e := range entries {
		id := e.Name()
		tables = append(tables, id)
	}
	return tables
}

func (api *API) rows(tableID string) []string {
	path := filepath.Join(api.DataDir, tableID)
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
		b = append(b, '\n')
		w.Write(b)
	})

	// List rows
	mux.HandleFunc("GET /{tableID}", func(w http.ResponseWriter, r *http.Request) {
		b, err := json.MarshalIndent(api.rows(r.PathValue("tableID")), "", "  ")
		if err != nil {
			panic(err)
		}
		b = append(b, '\n')
		w.Write(b)
	})

	// Fetch row
	mux.HandleFunc("GET /{tableID}/{rowID}", func(w http.ResponseWriter, r *http.Request) {})

	// Create row
	mux.HandleFunc("POST /{tableID}", func(w http.ResponseWriter, r *http.Request) {
		panic("todo")
	})

	// Set row
	mux.HandleFunc("PUT /{tableID}/{rowID}", func(w http.ResponseWriter, r *http.Request) {
		panic("todo")
	})

	// Update row
	mux.HandleFunc("PATCH /{tableID}/{rowID}", func(w http.ResponseWriter, r *http.Request) {
		panic("todo")
	})

	// Delete row
	mux.HandleFunc("DELETE /{tableID}/{rowID}", func(w http.ResponseWriter, r *http.Request) {
		panic("todo")
	})

	mux.ServeHTTP(w, r)
}
