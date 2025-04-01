package brass

import (
	"encoding/json"
	"net/http"
)

type Metadata struct {
	Types    map[string]*Type `json:"types"`
	RootType string           `json:"rootType"`
}

func (meta *Metadata) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(meta)
}
