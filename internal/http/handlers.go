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
