package brass

import "net/http"

func NewOrg(id string) *Org {
	return &Org{
		ID:        id,
		Libraries: map[string]Library{},
	}
}

type Org struct {
	ID        string
	Libraries map[string]Library
}

func (org *Org) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
