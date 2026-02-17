package config

type Technology struct {
	Image   string `json:"image"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Config struct {
	Technologies []Technology `json:"technologies"`
}
