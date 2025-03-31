package brass

import (
	"encoding/json"
	"net/http"
)

type Metadata struct {
	Types    map[string]*Type
	RootType string
}

func (meta *Metadata) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(meta)
}
