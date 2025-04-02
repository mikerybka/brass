package brass

import (
	"net/http"

	"github.com/mikerybka/english"
)

func NewLib(name *english.Name) *Lib {
	return &Lib{
		ID:   name.KebabCase(),
		Name: name,
	}
}

type Lib struct {
	ID    string          `json:"id"`
	Name  *english.Name   `json:"name"`
	Types map[string]Type `json:"types"`
}

func (lib *Lib) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
