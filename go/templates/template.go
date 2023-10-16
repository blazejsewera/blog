package templates

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/internal/whitespace"
	"html/template"
	"os"
)

type Template struct {
	*template.Template
}

var templateFS = os.DirFS("templates")

func ParseTFS(name ...string) *Template {
	tt, err := template.ParseFS(templateFS, name...)
	if err != nil {
		panic(fmt.Errorf("parse with default fs:%w", err))
	}
	return &Template{tt}
}

func (t *Template) ParseTFS(name ...string) *Template {
	tt, err := t.ParseFS(templateFS, name...)
	if err != nil {
		panic(fmt.Errorf("parse with default fs:%w", err))
	}
	return &Template{tt}
}

func (t *Template) With(component func(*Template)) *Template {
	component(t)
	return t
}

func (t *Template) Render(data any) []byte {
	buf := &bytes.Buffer{}
	err := t.Execute(buf, data)
	if err != nil {
		panic(fmt.Errorf("render: %w%s", err, t.DefinedTemplates()))
	}
	return whitespace.Collapse(buf.Bytes())
}
