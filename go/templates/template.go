package templates

import (
	"fmt"
	"github.com/blazejsewera/blog/templates/whitespace"
	"html/template"
	"io"
	"io/fs"
)

type AdditionalTemplate func(fs.FS, *template.Template) error

type Template struct {
	templateFS fs.FS
	template   *template.Template
}

func New() *Template {
	return NewWithFS(TemplateFS)
}

func NewWithFS(templateFS fs.FS) *Template {
	return &Template{
		templateFS: templateFS,
		template:   template.New("index.html.tmpl"),
	}
}

func (t *Template) With(fn AdditionalTemplate) *Template {
	err := fn(t.templateFS, t.template)
	if err != nil {
		panic(fmt.Sprintf("parsing template: %v%s", err, t.template.DefinedTemplates()))
	}
	return t
}

func (t *Template) Execute(wr io.Writer, data any) error {
	return t.template.Execute(whitespace.Collapse(wr), data)
}
