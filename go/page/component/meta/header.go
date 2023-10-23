package meta

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/internal/times"
	"html/template"
	"strings"
)

func Header(t *templates.Template) {
	templateNames := templates.Components("meta/header", "meta/base-header")
	t.ParseTFS(templateNames)
}

type HeaderProps struct {
	Title          string
	Date           times.Time
	Author         string
	License        string
	Language       string
	SiteName       string
	Abstract       string
	Keywords       []string
	ImgURL         template.URL
	ImgDescription string
}

func (p HeaderProps) ISODate() string {
	return p.Date.ISODate()
}

func (p HeaderProps) CommaSeparatedKeywords() string {
	return strings.Join(p.Keywords, ", ")
}
