package models

import (
	"html/template"
	"time"
)

// PageData t.b.d. until API stable
type PageData[T any] struct {
	Title   string
	Path    string
	CdnURL  string
	Content *T
}

// NoContent t.b.d. until API stable
type NoContent struct{}

// BlogIndex t.b.d. until API stable
type BlogIndex struct {
	Posts []*BlogPostMeta
}

// BlogPost t.b.d. until API stable
type BlogPost struct {
	Meta    *BlogPostMeta
	Content template.HTML
}

// BlogPostMeta t.b.d. until API stable
type BlogPostMeta struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Image       string    `yaml:"image"`
	Date        time.Time `yaml:"date"`
	Slug        string
}

// Experience t.b.d. until API stable
type Experience struct {
	Positions []*Position
}

// Position t.b.d. until API stable
type Position struct {
	StartDate   string
	EndDate     string
	Title       string
	Company     string
	Description string
	IsCurrent   bool
}
