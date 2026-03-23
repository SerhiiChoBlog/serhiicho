package http

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *server) postLikesIsLikedHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		PostID int    `json:"postID"`
		Secret string `json:"secret"` // can be empty string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likes, err := s.db.PostLikeRepo.CountPostLikes(req.PostID)
	if err != nil {
		log.Fatalln(err)
	}

	resp := map[string]any{"liked": false, "likes": likes}

	if req.Secret == "" {
		jsonResponse(w, 200, resp)
		return
	}

	user, err := s.db.UserRepo.BySecret(req.Secret)
	if err != nil {
		log.Fatalln(err)
	}

	if user == nil {
		log.Fatalf("user with secret '%s' doesn't exist", req.Secret)
	}

	resp["liked"], err = s.db.PostLikeRepo.UserLikedPost(req.PostID, user.ID)
	if err != nil {
		log.Fatalln(err)
	}

	jsonResponse(w, 200, resp)
}
