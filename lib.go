package brass

import "net/http"

func NewLib(id string) *Lib {
	return &Lib{}
}

type Lib struct {
	Types map[string]Type
}

func (lib *Lib) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
