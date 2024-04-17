package templates

import (
	"embed"
	"html/template"
)

//go:embed *.html
var tmplFS embed.FS

// Templates holds all parsed templates
var Templates *template.Template

func init() {
	// Parse all HTML templates on initialization
	var err error
	Templates, err = template.ParseFS(tmplFS, "*.html")
	if err != nil {
		panic(err) // In production, handle this error more gracefully
	}
}
