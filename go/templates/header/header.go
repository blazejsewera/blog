package header

import (
	"github.com/blazejsewera/blog/times"
	"html/template"
	"io/fs"
	"strings"
)

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

func Header(templateFS fs.FS, t *template.Template) error {
	t, err := t.ParseFS(templateFS, "header/*.html.tmpl")
	return err
}

func (p Props) ISODate() string {
	return p.Date.ISODate()
}

func (p Props) CommaSeparatedKeywords() string {
	return strings.Join(p.Keywords, ", ")
}
