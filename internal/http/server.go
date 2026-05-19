package http

import (
	"fmt"
	"net/http"
	"serhii/internal/config"
	"serhii/internal/database"

	"github.com/textwire/textwire/v5"
)

type server struct {
	mux  *http.ServeMux
	tpl  *textwire.Template
	conf *config.Config
	db   *database.Database
}

func NewServer(conf *config.Config, db *database.Database) server {
	return server{
		mux:  http.NewServeMux(),
		tpl:  newTpl(),
		conf: conf,
		db:   db,
	}
}

func (s *server) ListenAndServe(addr string) error {
	s.registerRoutes()

	fmt.Printf("[Server ] running on http://localhost:%s\n", addr)

	handler := chain(s.mux, stripTrailingSlash)

	return http.ListenAndServe(":"+addr, handler)
}
