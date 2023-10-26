package templates

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"html/template"
	"strings"
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

	tt, err := template.New("index.html.tmpl").
		Funcs(template.FuncMap{"require": require}).
		ParseFiles(toParse...)
	if err != nil {
		panic(fmt.Errorf("templates: parse all: %w", err))
	}
	return &Template{tt}
}

func (t *Template) Render(templateName string, data any) ([]byte, error) {
	buf := &bytes.Buffer{}
	err := t.ExecuteTemplate(buf, templateName, data)
	if err != nil {
		err = fmt.Errorf("render: %w%s", err, t.definedTemplates())
	}
	return buf.Bytes(), err
}

func (t *Template) definedTemplates() string {
	t.DefinedTemplates()
	s := strings.Replace(t.DefinedTemplates(), "; ", ";\n", 1)
	return strings.ReplaceAll(s, "\", ", "\",\n")
}
