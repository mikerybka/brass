package brass

import "net/http"

type App struct{}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {}
