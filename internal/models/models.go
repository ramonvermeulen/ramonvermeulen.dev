package models

import "html/template"

type PageData struct {
	Title   string
	Path    string
	Content template.HTML
}
