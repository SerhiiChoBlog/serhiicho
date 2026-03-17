package http

import (
	"fmt"
	"log"
	"net/http"
	"serhii/internal/model"
)

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

	s.tpl.Response(w, "~posts/index", data)
}
