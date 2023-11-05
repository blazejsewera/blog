package domain

import (
	"github.com/blazejsewera/blog/internal/times"
	"html/template"
)

type ArticleMetadata struct {
	Title            string
	Date             times.Time
	URL              template.URL
	SourceFile       string
	TargetFile       string
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

func (m ArticleMetadata) Equal(other ArticleMetadata) bool {
	return m.SourceFile == other.SourceFile
}

type Update struct {
	Date    times.Time
	DiffURL string
}

func (m ArticleMetadata) ShortDate() string {
	return m.Date.ShortDate()
}

var defaultValues = ArticleMetadata{
	Author:   "Blazej Sewera",
	Language: "en-US",
	License:  "CC BY 4.0",
}

func FillDefaultIfEmpty(metadata ArticleMetadata) ArticleMetadata {
	if metadata.Author == "" {
		metadata.Author = defaultValues.Author
	}
	if metadata.Language == "" {
		metadata.Language = defaultValues.Language
	}
	if metadata.License == "" {
		metadata.License = defaultValues.License
	}
	return metadata
}
