package brass

import (
	"net/http"

	"github.com/mikerybka/english"
)

func NewApp(name *english.Name) *App {
	return &App{
		ID:       name.KebabCase(),
		Name:     name,
		Types:    DefaultTypes,
		BaseType: "type",
	}
}

type App struct {
	ID       string          `json:"id"`
	Name     *english.Name   `json:"name"`
	Types    map[string]Type `json:"types"`
	BaseType string          `json:"baseType"`
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", a.ServeRoot)
	mux.HandleFunc("/types/", a.ServeTypes)
	mux.HandleFunc("/baseType/", a.ServeBaseType)
	mux.ServeHTTP(w, r)
}
func (a *App) ServeRoot(w http.ResponseWriter, r *http.Request)     {}
func (a *App) ServeTypes(w http.ResponseWriter, r *http.Request)    {}
func (a *App) ServeBaseType(w http.ResponseWriter, r *http.Request) {}
