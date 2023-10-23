package templates

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"html/template"
)

type Template struct {
	*template.Template
}

const pageDir = "page"

func ParseAll() *Template {
	toParse, err := files.FindBySuffix(pageDir, ".html.tmpl")
	if err != nil {
		panic(fmt.Errorf("templates: parse all: %w", err))
	}

	tt, err := template.ParseFiles(toParse...)
	if err != nil {
		panic(fmt.Errorf("templates: parse all: %w", err))
	}
	return &Template{tt}
}

func (t *Template) Render(data any) []byte {
	buf := &bytes.Buffer{}
	err := t.Execute(buf, data)
	if err != nil {
		panic(fmt.Errorf("render: %w%s", err, t.DefinedTemplates()))
	}
	return buf.Bytes()
}
