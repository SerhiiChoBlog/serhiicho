package http

import (
	"log"
	"net/http"
)

func (s *server) homeHandler(w http.ResponseWriter, _ *http.Request) {
	latest, err := s.db.Post.Latest()
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]any{
		"latest": latest,
	}

	s.tpl.Response(w, "~home", data)
}

func (s *server) aboutHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]any{
		"technologies": s.conf.Technologies,
	}

	s.tpl.Response(w, "~about-me", data)
}
