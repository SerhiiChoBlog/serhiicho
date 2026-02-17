package http

import (
	"log"
	"net/http"
)

func (s *server) homeHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]any{}

	if err := s.tpl.Response(w, "views/home", data); err != nil {
		log.Println(err)
	}
}

func (s *server) aboutHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]any{
		"technologies": s.conf.Technologies,
	}

	if err := s.tpl.Response(w, "views/about-me", data); err != nil {
		log.Println(err)
	}
}
