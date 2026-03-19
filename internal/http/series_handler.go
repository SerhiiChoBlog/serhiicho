package http

import (
	"log"
	"net/http"
)

func (s *server) seriesHandler(w http.ResponseWriter, r *http.Request) {
	series, err := s.db.Series.List()
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]any{
		"series": series,
	}

	s.tpl.Response(w, "~series/index", data)
}
