package model

func AttachPostsToSeries(posts []*Post, series []*Series) error {
	// Group posts by series ID
	postsBySeries := make(map[int][]*Post)
	for _, p := range posts {
		postsBySeries[p.SeriesID] = append(postsBySeries[p.SeriesID], p)
	}

	// Attach series to posts
	for _, s := range series {
		s.Posts = postsBySeries[s.ID]
	}

	return nil
}

func AttachTagsToPosts(tags []*Tag, posts []*Post) error {
	// Group tags by post ID
	tagsByPost := make(map[int][]*Tag)
	for _, tag := range tags {
		tagsByPost[tag.PostID] = append(tagsByPost[tag.PostID], tag)
	}

	// Attach tags to posts
	for _, post := range posts {
		post.Tags = tagsByPost[post.ID]
	}

	return nil
}

func AttachSeriesToPosts(series []*Series, posts []*Post) error {
	// Group series by post ID
	seriesByPost := make(map[int][]*Series)
	for _, s := range series {
		seriesByPost[s.PostID] = append(seriesByPost[s.PostID], s)
	}

	// Attach posts to series
	for _, post := range posts {
		post.Series = seriesByPost[post.ID]
	}

	return nil
}
