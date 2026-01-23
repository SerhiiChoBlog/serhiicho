package http

import (
	"fmt"
	"net/http"

	"github.com/textwire/textwire/v2"
	"github.com/textwire/textwire/v2/config"
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
	http.HandleFunc("/about", aboutHandler)
}

func initTextwire() {
	twConf := &config.Config{
		TemplateExt: ".tw",
		DebugMode:   true,
	}

	twTpl, err := textwire.NewTemplate(twConf)
	if err != nil {
		fmt.Println(err)
	}

	tpl = twTpl
}
