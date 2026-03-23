package model

func AttachPostsToSeries(posts []*Post, series []*Series) error {
	for _, s := range series {
		for _, post := range posts {
			if post.SeriesID != s.ID {
				continue
			}
			s.Posts = append(s.Posts, post)
		}
	}
	return nil
}

func AttachTagsToPosts(tags []*Tag, posts []*Post) error {
	// Group tags by post ID
	tagsByPost := make(map[int][]Tag)
	for _, tag := range tags {
		tagsByPost[tag.PostID] = append(tagsByPost[tag.PostID], tag)
	}

	// Attach tags to posts
	for _, post := range posts {
		post.Tags = tagsByPost[post.ID]
	}

	return nil
}
