package brass

import "net/http"

func NewServer() http.Handler {
	return http.NotFoundHandler()
}
