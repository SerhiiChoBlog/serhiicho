package http

import (
	"fmt"
	"net/http"
	"os"
	"serhii/internal/config"
	"serhii/internal/database"

	"github.com/textwire/textwire/v4"
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
	s.mux.HandleFunc("GET /about-me", s.aboutMeHandler)
	s.mux.HandleFunc("GET /posts", s.postsHandler)
	s.mux.HandleFunc("GET /posts/{slug}", s.postHandler)
	s.mux.HandleFunc("GET /series", s.seriesHandler)

	// API routes
	s.mux.HandleFunc("POST /api/post-likes/is-liked", s.postLikesIsLikedHandler)

	fmt.Printf("[Server ] running on http://localhost:%s\n", addr)

	handler := chain(s.mux, stripTrailingSlash)

	return http.ListenAndServe(":"+addr, handler)
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
