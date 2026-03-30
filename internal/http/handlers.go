package http

import (
	"fmt"
	"log"
	"net/http"
	"serhii/internal/model"
	"strings"
)

func (s *server) homeHandler(w http.ResponseWriter, _ *http.Request) {
	latest, err := s.db.PostRepo.LatestWithTags(s.db.TagRepo)
	if err != nil {
		log.Fatalln(err)
	}

	latestVideos, err := s.db.VideoRepo.Latest()
	if err != nil {
		log.Fatalln(err)
	}

	comingSoonPosts, err := s.db.PostRepo.ComingSoonWithTagsAndSeries(s.db.TagRepo, s.db.SeriesRepo)
	if err != nil {
		log.Fatalln(err)
	}

	s.tpl.Response(w, "~home", map[string]any{
		"latest":          latest,
		"latestVideos":    latestVideos,
		"comingSoonPosts": comingSoonPosts,
	})
}

func (s *server) aboutMeHandler(w http.ResponseWriter, _ *http.Request) {
	s.tpl.Response(w, "~about-me", map[string]any{
		"technologies": s.conf.Technologies,
	})
}

func (s *server) listPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.PostRepo.AllWithTags(s.db.TagRepo)
	if err != nil {
		log.Fatalln(err)
	}

	title := "Articles for developers"
	urlTag := r.URL.Query().Get("tag")
	var tag *model.Tag

	if urlTag != "" {
		tag, err = s.db.TagRepo.SingleByName(urlTag)
		if err != nil {
			log.Fatalln(err)
		}
		title = fmt.Sprintf("Read about %s", tag.Title)
	}

	var series []*model.Series
	if len(posts) > 0 {
		series = posts[0].Series
	}

	s.tpl.Response(w, "~posts/list", map[string]any{
		"title":  title,
		"posts":  posts,
		"series": series,
		"tag":    tag,
	})
}

func (s *server) singlePostHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	if strings.Contains(slug, ".") {
		http.NotFound(w, r)
		return
	}

	post, err := s.db.PostRepo.SingleWithTagsAndSeries(slug, s.db.TagRepo, s.db.SeriesRepo)
	if err != nil {
		log.Fatalln(err)
	}

	var series *model.Series
	if len(post.Series) > 0 {
		series = post.Series[0]
	}

	postTitle := post.Title
	if series != nil {
		postTitle = series.Title
	}

	s.tpl.Response(w, "~posts/single", map[string]any{
		"post":          post,
		"postTitle":     postTitle,
		"titleFontSize": getPostTitleFontSize(post.Title),
		"series":        series,
	})
}

func (s *server) listSeriesHandler(w http.ResponseWriter, r *http.Request) {
	series, err := s.db.SeriesRepo.AllWithPosts(s.db.PostRepo)
	if err != nil {
		log.Fatalln(err)
	}

	s.tpl.Response(w, "~series/list", map[string]any{
		"series": series,
	})
}

func (s *server) singleSeriesHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	if strings.Contains(slug, ".") {
		http.NotFound(w, r)
		return
	}

	series, err := s.db.SeriesRepo.SingleWithPosts(slug, s.db.PostRepo)
	if err != nil {
		log.Fatalln(err)
	}

	s.tpl.Response(w, "~series/single", map[string]any{
		"series": series,
	})
}
