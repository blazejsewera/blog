package listing

import (
	"github.com/blazejsewera/blog/internal/times"
	"github.com/blazejsewera/blog/templates"
)

var templateNames = []string{"listing/article-listing.html.tmpl"}

type Update struct {
	Date    times.Time
	DiffURL string
}

type ArticleInfo struct {
	Title            string
	Date             times.Time
	URL              string
	File             string
	Subtitle         string
	Updates          []Update
	Draft            bool
	DraftDescription string
	Author           string
	Abstract         string
	Keywords         []string
	Language         string
	License          string
	ImgURL           string
	MinImgURL        string
	ImgDescription   string
}

func (i ArticleInfo) ShortDate() string {
	return i.Date.ShortDate()
}

type Props struct {
	Title          string
	TabTitle       string
	Subtitle       string
	Description    string
	ImgURL         string
	ImgDescription string
	Articles       []ArticleInfo
}

func Listing(t *templates.Template) {
	t.ParseTFS(templateNames...)
}
