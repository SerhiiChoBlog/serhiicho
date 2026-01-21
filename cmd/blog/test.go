package main

import (
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tpl.Response(w, "home", map[string]any{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
