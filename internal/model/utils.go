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
