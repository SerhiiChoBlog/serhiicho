package http

import (
	"log"
	"os"
	"time"

	"github.com/SerhiiCho/timeago/v3"
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
			"appURL":          os.Getenv("APP_URL"),
			"darkThemeCookie": "dark",
			"manifest": map[string]string{
				"js":  "",
				"css": "",
			},
		},
	}

	tpl, err := textwire.NewTemplate(twConf)
	err.FatalOnError()

	registerCustomFuncs()

	return tpl
}

func registerCustomFuncs() {
	textwire.RegisterStrFunc("timeago", func(s string, args ...any) any {
		t, err := time.Parse("2006-01-02 15:04:05", s)
		if err != nil {
			log.Fatal(err)
		}

		out, err := timeago.Parse(t)
		if err != nil {
			log.Fatal(err)
		}

		return out
	})
}
