package http

import (
	"fmt"
	"net/http"

	"github.com/textwire/textwire/v2"
	"github.com/textwire/textwire/v2/config"
)

var tpl *textwire.Template

type server struct {
	addr string // like :8080
}

// addr like :8080
func NewServer(addr string) server {
	return server{addr}
}

func (s *server) ListenAndServe() error {
	initTextwire()

	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running...")

	return http.ListenAndServe(s.addr, nil)
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
