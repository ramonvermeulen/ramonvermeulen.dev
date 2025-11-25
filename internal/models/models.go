package models

import (
	"fmt"
	"html/template"
	"strings"
	"time"
)

// PageData is a generic struct for passing data to templates.
type PageData[T any] struct {
	Title       string
	Description string
	Path        string
	CdnURL      string
	Content     T
}

// NoContent is used for pages without specific content.
type NoContent struct{}

// BlogIndex represents the blog index page content.
type BlogIndex struct {
	Posts []*BlogPostMeta
}

// BlogPost represents a single blog post with metadata and content.
type BlogPost struct {
	Meta    *BlogPostMeta
	Content template.HTML
}

// BlogPostMeta represents metadata for a blog post.
type BlogPostMeta struct {
	Title       string    `yaml:"title"`
	Description string    `yaml:"description"`
	Image       string    `yaml:"image"`
	Date        time.Time `yaml:"date"`
	Slug        string
}

// Technology struct represents a technology used in a position
type Technology struct {
	Name string
	URL  string
}

// Position struct represents a work experience position
type Position struct {
	CompanyName    string
	CompanyWebsite string
	StartDate      time.Time
	EndDate        *time.Time
	Title          string
	Location       string
	Description    string
	Technologies   []Technology
}

// StartDateShort returns the start date formatted as "Jan 2006".
func (p *Position) StartDateShort() string {
	if p.StartDate.IsZero() {
		return ""
	}
	return p.StartDate.Format("Jan 2006")
}

// EndDateShort returns "Present" if ongoing, else formatted date.
func (p *Position) EndDateShort() string {
	if p.EndDate == nil || p.EndDate.IsZero() {
		return "Present"
	}
	return p.EndDate.Format("Jan 2006")
}

// Duration returns a human-readable duration between start and end (or now).
func (p *Position) Duration() string {
	end := time.Now()
	if p.EndDate != nil && !p.EndDate.IsZero() {
		end = *p.EndDate
	}
	years, months := diffTime(p.StartDate, end)
	if years < 0 || months < 0 {
		return "Less than a month"
	}
	var parts []string
	if years > 0 {
		if years == 1 {
			parts = append(parts, "1 year")
		} else {
			parts = append(parts, fmt.Sprintf("%d years", years))
		}
	}
	if months > 0 {
		if months == 1 {
			parts = append(parts, "1 month")
		} else {
			parts = append(parts, fmt.Sprintf("%d months", months))
		}
	}
	if len(parts) == 0 {
		return "Less than a month"
	}
	return strings.Join(parts, ", ")
}

// diffTime computes year/month difference using the same logic as LinkedIn.
func diffTime(a, b time.Time) (years, months int) {
	if a.After(b) {
		a, b = b, a
	}
	years = b.Year() - a.Year()
	months = int(b.Month()) - int(a.Month())
	if b.Day() >= a.Day() {
		months++
	}
	if months < 0 {
		years--
		months += 12
	}
	if months >= 12 {
		years += months / 12
		months = months % 12
	}
	return
}
