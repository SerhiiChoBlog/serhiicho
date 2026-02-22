package http

import (
	"log"
	"os"

	"github.com/textwire/textwire/v3"
	"github.com/textwire/textwire/v3/config"
)

func newTpl() *textwire.Template {
	isDev := os.Getenv("APP_ENV") == "development"

	twConf := &config.Config{
		DebugMode:   isDev,
		FileWatcher: isDev,
		GlobalData: map[string]any{
			"env":             os.Getenv("APP_ENV"),
			"appName":         os.Getenv("APP_NAME"),
			"darkThemeCookie": "", // dark
			"manifest": map[string]string{
				"js":  "",
				"css": "",
			},
		},
	}

	tpl, err := textwire.NewTemplate(twConf)
	if err != nil {
		log.Fatal(err)
	}

	return tpl
}
