package main

import (
	"fmt"
	"net/http"

	"github.com/textwire/textwire/v2"
	"github.com/textwire/textwire/v2/config"
)

var tpl *textwire.Template

func main() {
	var err error

	tpl, err = textwire.NewTemplate(&config.Config{
		TemplateExt: ".tw",
		DebugMode:   true,
	})
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/", homeHandler)

	fmt.Println("Server is running...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}
