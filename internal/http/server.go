package http

import (
	"fmt"
	"net/http"
	"os"
	"serhii/internal/config"
	"serhii/internal/database"

	"github.com/textwire/textwire/v3"
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
	s.servePublic("GET /", "./public")
	s.mux.HandleFunc("GET /about-me", s.aboutHandler)
	s.mux.HandleFunc("GET /posts", s.postsHandler)

	fmt.Printf("Server is running on http://localhost:%s\n", addr)

	return http.ListenAndServe(":"+addr, s.mux)
}

func (s *server) servePublic(pattern, pubDir string) {
	fileServer := http.FileServer(http.Dir(pubDir))

	s.mux.Handle(pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := pubDir + r.URL.Path
		info, err := os.Stat(path)
		if err == nil && !info.IsDir() {
			fileServer.ServeHTTP(w, r)
			return
		}

		if r.URL.Path == "/" {
			s.homeHandler(w, r)
			return
		}

		http.NotFound(w, r)
	}))
}
