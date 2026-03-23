package http

import (
	"fmt"
	"log"
	"net/http"
	"serhii/internal/model"
	"strings"
)

func (s *server) homeHandler(w http.ResponseWriter, _ *http.Request) {
	latest, err := s.db.PostRepo.Latest()
	if err != nil {
		log.Fatalln(err)
	}

	s.tpl.Response(w, "~home", map[string]any{
		"latest": latest,
	})
}

func (s *server) aboutMeHandler(w http.ResponseWriter, _ *http.Request) {
	s.tpl.Response(w, "~about-me", map[string]any{
		"technologies": s.conf.Technologies,
	})
}

func (s *server) postsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.PostRepo.All()
	if err != nil {
		log.Fatalln(err)
	}

	title := "Articles for developers"
	urlTag := r.URL.Query().Get("tag")
	var tag *model.Tag

	if urlTag != "" {
		tag, err = s.db.TagRepo.ByName(urlTag)
		if err != nil {
			log.Fatalln(err)
		}
		title = fmt.Sprintf("Read about %s", tag.Title)
	}

	var series []*model.Series
	if len(posts) > 0 {
		series = posts[0].Series
	}

	s.tpl.Response(w, "~posts/posts", map[string]any{
		"title":  title,
		"posts":  posts,
		"series": series,
		"tag":    tag,
	})
}

func (s *server) postHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	if strings.Contains(slug, ".") {
		http.NotFound(w, r)
		return
	}

	post, err := s.db.PostRepo.Single(slug)
	if err != nil {
		log.Fatalln(err)
	}

	postSeries, err := s.db.SeriesRepo.PostSeries(post.ID)
	if err != nil {
		log.Fatalln(err)
	}

	post.Series = postSeries

	var series *model.Series
	if len(post.Series) > 0 {
		series = post.Series[0]
	}

	postTitle := post.Title
	if series != nil {
		postTitle = series.Title
	}

	s.tpl.Response(w, "~posts/post", map[string]any{
		"post":          post,
		"postTitle":     postTitle,
		"titleFontSize": getPostTitleFontSize(post.Title),
		"series":        series,
	})
}

func (s *server) seriesHandler(w http.ResponseWriter, r *http.Request) {
	series, err := s.db.SeriesRepo.WithPosts(s.db.PostRepo)
	if err != nil {
		log.Fatalln(err)
	}

	s.tpl.Response(w, "~series/series", map[string]any{
		"series": series,
	})
}
