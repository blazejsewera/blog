package page_test

import (
	"github.com/blazejsewera/blog/renderer/domain"
	"github.com/blazejsewera/blog/renderer/internal/times"
	"html/template"
)

var headerData = map[string]string{
	"blogRoot":       "blog.sewera.dev",
	"title":          "title",
	"date":           "2023-10-01",
	"author":         "author",
	"license":        "license",
	"language":       "language",
	"abstract":       "abstract",
	"keyword1":       "keyword1",
	"keyword2":       "keyword2",
	"imgURL":         "https://example.com/img.jpg",
	"imgDescription": "imgDescription",
}

var listingData = map[string]string{
	"description": "description",

	"articleTitle":  "articleTitle",
	"articleDate":   "October 2, 2023",
	"articleURL":    "https://example.com/article",
	"articleImgURL": "https://example.com/article/img.jpg",

	"articleDate2": "October 3, 2023",
}

var articleData = map[string]string{
	"draftDescription": "draftDescription",
	"abstract":         "This is an abstract.",
	"content1":         "An article",
	"content2":         "Sample text.",
}

var rawContent = []byte(`# An article

Sample text.`)

var articleMetadata = domain.ArticleMetadata{
	Title:            headerData["title"],
	Date:             times.Parse(headerData["date"]),
	Author:           headerData["author"],
	License:          headerData["license"],
	Language:         headerData["language"],
	Abstract:         articleData["abstract"],
	Keywords:         []string{headerData["keyword1"], headerData["keyword2"]},
	ImgURL:           template.URL(headerData["imgURL"]),
	ImgDescription:   headerData["imgDescription"],
	Draft:            true,
	DraftDescription: articleData["draftDescription"],
}

var articles = []domain.ArticleMetadata{{
	Title:  listingData["articleTitle"],
	Date:   times.Parse("2023-10-02"),
	URL:    template.URL(listingData["articleURL"]),
	Draft:  false,
	ImgURL: template.URL(listingData["articleImgURL"]),
}, {
	Title:  listingData["articleTitle"],
	Date:   times.Parse("2023-10-03"),
	URL:    template.URL(listingData["articleURL"]),
	Draft:  true,
	ImgURL: template.URL(listingData["articleImgURL"]),
}}
