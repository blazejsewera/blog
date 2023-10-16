package header

import (
	"github.com/blazejsewera/blog/internal/times"
	"github.com/blazejsewera/blog/templates"
	"strings"
)

var templateNames = []string{"header/article-header.html.tmpl", "header/base-header.html.tmpl"}

type Props struct {
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

func Header(t *templates.Template) {
	t.ParseTFS(templateNames...)
}

func (p Props) ISODate() string {
	return p.Date.ISODate()
}

func (p Props) CommaSeparatedKeywords() string {
	return strings.Join(p.Keywords, ", ")
}
