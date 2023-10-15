package header

import (
	"github.com/blazejsewera/blog/templates"
	"github.com/blazejsewera/blog/times"
	"html/template"
	"strings"
)

type Header struct {
	Title          string
	Date           times.Time
	Author         string
	License        string
	Language       string
	SiteName       string
	Abstract       string
	Keywords       []string
	ImgURL         string
	ImgDescription string
}

func WithHeader(t *template.Template) {
	template.Must(t.ParseFS(templates.TemplateFS, "header/*.html.tmpl"))
}

func (h Header) ISODate() string {
	return h.Date.ISODate()
}

func (h Header) CommaSeparatedKeywords() string {
	return strings.Join(h.Keywords, ", ")
}
