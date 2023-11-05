package section

import (
	"github.com/blazejsewera/blog/domain"
	"html/template"
)

type FooterProps struct {
	Previous domain.ArticleMetadata
	Current  domain.ArticleMetadata
	Next     domain.ArticleMetadata
}

func (p FooterProps) Site() domain.Site {
	return domain.DefaultSite
}

func (p FooterProps) Source() template.URL {
	if p.Current.SourceFile == "" {
		return ""
	}
	return template.URL(p.Site().BlogSourceRootURL + p.Current.SourceFile)
}
