package handler

import (
	"path/filepath"
	"text/template"
)

func renderTemplate(templateFile string) *template.Template {
	layout := filepath.Join("static", "templates", "layout.html")
	footer := filepath.Join("static", "templates", "footer.html")
	header := filepath.Join("static", "templates", "header.html")
	page := filepath.Join("static", "templates", templateFile)

	return template.Must(template.ParseFiles(layout, header, footer, page))
}
