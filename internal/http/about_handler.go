package http

import (
	"net/http"
)

func (s *server) aboutMeHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]any{
		"technologies": s.conf.Technologies,
	}

	s.tpl.Response(w, "~about-me", data)
}
