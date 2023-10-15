package index

import (
	"github.com/blazejsewera/blog/templates/header"
	"html/template"
	"io/fs"
)

type Props struct {
	Header header.Props
	Posts  []string
}

func Index(templateFS fs.FS, t *template.Template) error {
	t, err := t.ParseFS(templateFS, "index/*.html.tmpl")
	return err
}
