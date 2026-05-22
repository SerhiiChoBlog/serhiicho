package http

import (
	"net/http"
	"serhii/internal/model"
)

func (s *server) register(w http.ResponseWriter, r *http.Request) {
	user := &model.User{}

	user.Email = r.FormValue("email")
	user.Name = r.FormValue("name")

	if !validateUserEmail(user.Email) || !validateUserName(user.Name) {
		return
	}

	rawPass := r.FormValue("password")
}

func validateUserEmail(email string) bool {
	return false
}

func validateUserName(name string) bool {
	return false
}
