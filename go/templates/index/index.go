package index

import (
	"html/template"
	"io/fs"
)

func Index(templateFS fs.FS, t *template.Template) error {
	t, err := t.ParseFS(templateFS, "index/*.html.tmpl")
	return err
}
