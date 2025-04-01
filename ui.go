package brass

import (
	"net/http"
)

func NewUI() *UI {
	return &UI{
		s: http.NewServeMux(),
	}
}

type UI struct {
	s *http.ServeMux
}

func (ui *UI) Handler() http.Handler {
	return ui.s
}
