package http

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/textwire/textwire/v3"
	"github.com/textwire/textwire/v3/config"
)

var tpl *textwire.Template

type server struct {
}

func NewServer() server {
	return server{}
}

// addr like :8080
func (s *server) ListenAndServe(addr string) error {
	initTextwire()

	s.registerRoutes()

	fmt.Println("Server is running on", addr)

	return http.ListenAndServe(addr, nil)
}

func (s *server) registerRoutes() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about-me", aboutHandler)
}

func initTextwire() {
	twConf := &config.Config{
		DebugMode: true,
		GlobalData: map[string]any{
			"env": os.Getenv("APP_ENV"),
			"manifest": map[string]string{
				"js":  "",
				"css": "",
			},
		},
	}

	twTpl, err := textwire.NewTemplate(twConf)
	if err != nil {
		log.Fatal(err)
	}

	tpl = twTpl
}
