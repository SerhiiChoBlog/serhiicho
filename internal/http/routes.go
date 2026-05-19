package http

import (
	"net/http"
	"os"
)

func (s *server) registerRoutes() {
	s.servePublic("GET /", "./public")

	s.mux.HandleFunc("GET /about-me", s.aboutMeHandler)
	s.mux.HandleFunc("GET /posts", s.listPostsHandler)
	s.mux.HandleFunc("GET /posts/{slug}", s.singlePostHandler)
	s.mux.HandleFunc("GET /series", s.listSeriesHandler)
	s.mux.HandleFunc("GET /series/{slug}", s.singleSeriesHandler)
	s.mux.HandleFunc("GET /projects", s.listProjectsHandler)
	s.mux.HandleFunc("GET /projects/{name}", s.singleProjectHandler)

	// API routes
	s.mux.HandleFunc("POST /api/post-likes/is-liked", s.postLikesIsLikedHandler)
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
