package http

import (
	"fmt"
	"log"
	"net/http"
	"serhii/internal/model"
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

func (s *server) aboutMeHandler(w http.ResponseWriter, _ *http.Request) {
	data := map[string]any{
		"technologies": s.conf.Technologies,
	}

	s.tpl.Response(w, "~about-me", data)
}

func (s *server) postsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.Post.List()
	if err != nil {
		log.Fatalln(err)
	}

	title := "Articles for developers"
	urlTag := r.URL.Query().Get("tag")
	var tag *model.Tag

	if urlTag != "" {
		tag, err = s.db.Tag.ByName(urlTag)
		if err != nil {
			log.Fatalln(err)
		}
		title = fmt.Sprintf("Read about %s", tag.Title)
	}

	var series []model.Series
	if len(posts) > 0 {
		series = posts[0].Series
	}

	data := map[string]any{
		"title":  title,
		"posts":  posts,
		"series": series,
		"tag":    tag,
	}

	s.tpl.Response(w, "~posts/list", data)
}

func (s *server) postHandler(w http.ResponseWriter, r *http.Request) {
	post, err := s.db.Post.Single("php-revival-extension-a-new-look-at-php-documentation")
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]any{
		"post": post,
	}

	s.tpl.Response(w, "~posts/single", data)
}

func (s *server) seriesHandler(w http.ResponseWriter, r *http.Request) {
	series, err := s.db.Series.List()
	if err != nil {
		log.Fatalln(err)
	}

	data := map[string]any{
		"series": series,
	}

	s.tpl.Response(w, "~series/list", data)
}
