package http

import (
	"fmt"
	"os"

	"github.com/SerhiiCho/timeago/v3"
	"github.com/textwire/textwire/v4"
	"github.com/textwire/textwire/v4/config"
)

func newTpl() *textwire.Template {
	registerCustomFuncs()

	isDev := os.Getenv("APP_ENV") == "development"

	twConf := &config.Config{
		DebugMode:     isDev,
		FileWatcher:   isDev,
		ErrorPagePath: "error-page",
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

	return tpl
}

func registerCustomFuncs() {
	textwire.RegisterStrFunc("_timeago", func(s string, args ...any) any {
		out, err := timeago.Parse(s)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		return out
	})
}
