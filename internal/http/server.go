package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/textwire/textwire/v3"
)

type server struct {
	mux *http.ServeMux
	tpl *textwire.Template
}

func NewServer() server {
	return server{
		mux: http.NewServeMux(),
		tpl: newTpl(),
	}
}

func (s *server) ListenAndServe(addr string) error {
	s.servePublic("GET /", "./public")
	s.mux.HandleFunc("GET /about-me", s.aboutHandler)

	fmt.Printf("Server is running on http://localhost%s\n", addr)

	return http.ListenAndServe(addr, s.mux)
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
