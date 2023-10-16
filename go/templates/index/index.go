package index

import (
	"github.com/blazejsewera/blog/templates"
	"github.com/blazejsewera/blog/templates/header"
)

var templateNames = []string{"index/index.html.tmpl"}

type Template struct {
	t *templates.Template
}

type Props struct {
	Header header.Props
	Posts  []string
}

func Index() *Template {
	return &Template{templates.
		ParseTFS(templateNames...).
		With(header.Header)}
}

func (t *Template) Render(props Props) []byte {
	return t.t.Render(props)
}
