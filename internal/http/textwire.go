package http

import (
	"os"

	"github.com/textwire/textwire/v4"
	"github.com/textwire/textwire/v4/config"
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
	err.FatalOnError()

	return tpl
}
