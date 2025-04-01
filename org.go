package brass

import "net/http"

func NewOrg(id string) *Org {
	return &Org{
		ID:        id,
		Libraries: map[string]Lib{},
	}
}

type Org struct {
	ID        string
	Libraries map[string]Lib
}

func (org *Org) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
