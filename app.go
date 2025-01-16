package brass

import "net/http"

type App struct {
	DataDir string
	SrcDir  string
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok\n"))
}
