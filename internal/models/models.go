package models

import "html/template"

// PageData t.b.d. until API stable
type PageData struct {
	Title    string
	Path     string
	BlogPost *BlogPost
}

// BlogPost t.b.d. until API stable
type BlogPost struct {
	Meta    *BlogPostMeta
	Content template.HTML
}

// BlogPostMeta t.b.d. until API stable
type BlogPostMeta struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Image       string `yaml:"image"`
	Date        string `yaml:"date"`
}
