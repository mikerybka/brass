package brass

import "net/http"

type API struct {
	DataDir string
	SrcDir  string
}

func (api *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}
