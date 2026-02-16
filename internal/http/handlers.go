package http

import (
	"log"
	"net/http"
	"os"
)

func (s *server) homeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"appName": os.Getenv("APP_NAME"),
		"env":     os.Getenv("APP_ENV"),
	}

	if err := s.tpl.Response(w, "views/home", data); err != nil {
		log.Println(err)
	}
}

func (s *server) aboutHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"technologies": []string{},
	}

	if err := s.tpl.Response(w, "views/about-me", data); err != nil {
		log.Println(err)
	}
}
