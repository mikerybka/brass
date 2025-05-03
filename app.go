package brass

import "net/http"

type App struct {
	Icon []byte // 1024x1024 pixel .png file
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()
	mux.ServeHTTP(w, r)
}
