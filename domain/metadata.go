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
	LicenseURL       template.URL
	ImgURL           template.URL
	ImgDescription   string
	Previous         PartialMetadata
	Next             PartialMetadata
}

func (m ArticleMetadata) EqualSource(markdownSourceFile string) bool {
	return m.SourceFile == markdownSourceFile
}

func (m ArticleMetadata) ShortDate() string {
	return m.Date.ShortDate()
}

func (m ArticleMetadata) Year() int {
	return m.Date.Year()
}

type Update struct {
	Date    times.Time
	DiffURL string
}

type PartialMetadata struct {
	Title          string
	Subtitle       string
	Author         string
	ImgURL         template.URL
	ImgDescription string
	Date           times.Time
	URL            template.URL
}

func PartialFromArticleMetadata(m ArticleMetadata) PartialMetadata {
	return PartialMetadata{
		Title:          m.Title,
		Subtitle:       m.Subtitle,
		Author:         m.Author,
		ImgURL:         m.ImgURL,
		ImgDescription: m.ImgDescription,
		Date:           m.Date,
		URL:            m.URL,
	}
}

var defaultValues = ArticleMetadata{
	Author:     "Blazej Sewera",
	Language:   "en-US",
	License:    "CC BY-SA 4.0",
	LicenseURL: "https://creativecommons.org/licenses/by-sa/4.0/",
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
	if metadata.LicenseURL == "" {
		metadata.LicenseURL = defaultValues.LicenseURL
	}
	return metadata
}
