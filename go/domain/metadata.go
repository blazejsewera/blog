package domain

import (
	"github.com/blazejsewera/blog/internal/times"
	"html/template"
)

type ArticleMetadata struct {
	Title            string
	Date             times.Time
	URL              template.URL
	Subtitle         string
	Updates          []Update
	Draft            bool
	DraftDescription string
	Author           string
	Abstract         string
	Keywords         []string
	Language         string
	License          string
	ImgURL           template.URL
	ImgDescription   string
}

type Update struct {
	Date    times.Time
	DiffURL string
}

func (i ArticleMetadata) ShortDate() string {
	return i.Date.ShortDate()
}
