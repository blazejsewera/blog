package meta

import (
	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/internal/times"
	"github.com/blazejsewera/blog/renderer/page/component"

	"html/template"
	"strings"
)

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

func HeaderPropsFromDomain(metadata article.Metadata) HeaderProps {
	site := component.DefaultSite

	return HeaderProps{
		Title:          metadata.Title,
		Date:           metadata.Date,
		Author:         metadata.Author,
		License:        metadata.License,
		Language:       metadata.Language,
		SiteName:       site.Name,
		Abstract:       metadata.Abstract,
		Keywords:       metadata.Keywords,
		ImgURL:         metadata.ImgURL,
		ImgDescription: metadata.ImgDescription,
	}
}
