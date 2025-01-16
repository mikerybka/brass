package brass

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
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
	tables := []string{}
	for _, t := range api.types() {
		tables = append(tables, t.PluralName.ID())
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

// mutator returns the path of the executable used in POST, PUT, PATCH and DELETE requests.
func (api *API) mutator(typeID string) string {
	// mkdir {datadir}/cmd/{typeID}
	cmdPath := filepath.Join(api.DataDir, "cmd", typeID)
	err := os.MkdirAll(cmdPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// go mod init
	cmd := exec.Command("go", "mod", "init", typeID)
	cmd.Dir = cmdPath
	cmd.Env = os.Environ()
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(b))
		panic(err)
	}

	// generate main.go
	t := api.typ(typeID)
	mainGo := t.mutatorCmd()
	err = os.WriteFile(filepath.Join(cmdPath, "main.go"), []byte(mainGo), os.ModePerm)
	if err != nil {
		panic(err)
	}

	// go build -o binpath main.go
	binPath := filepath.Join(api.DataDir, "bin", typeID)
	cmd = exec.Command("go", "build", "-o", binPath, "main.go")
	cmd.Dir = cmdPath
	cmd.Env = os.Environ()
	b, err = cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(b))
		panic(err)
	}

	return binPath
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
	mux.HandleFunc("POST /{tableID}", func(w http.ResponseWriter, r *http.Request) {})

	// Update row
	mux.HandleFunc("PUT /{tableID}/{rowID}", func(w http.ResponseWriter, r *http.Request) {})
	mux.HandleFunc("PATCH /{tableID}/{rowID}", func(w http.ResponseWriter, r *http.Request) {})

	// Delete row
	mux.HandleFunc("DELETE /{tableID}/{rowID}", func(w http.ResponseWriter, r *http.Request) {})

	mux.ServeHTTP(w, r)
}
