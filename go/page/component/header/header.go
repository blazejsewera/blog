package header

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/internal/times"
	"html/template"
	"strings"
)

var templateNames = templates.Components("header/header", "header/base-header")

func Header(t *templates.Template) {
	t.ParseTFS(templateNames)
}

type Props struct {
	RawCSS         []byte
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

func (p Props) CSS() template.CSS {
	return template.CSS(p.RawCSS)
}

func (p Props) ISODate() string {
	return p.Date.ISODate()
}

func (p Props) CommaSeparatedKeywords() string {
	return strings.Join(p.Keywords, ", ")
}
