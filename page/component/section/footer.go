package section

import (
	"github.com/blazejsewera/blog/domain"
	"html/template"
)

type FooterProps struct {
	Metadata domain.ArticleMetadata
}

func FooterPropsFromDomain(metadata domain.ArticleMetadata) FooterProps {
	return FooterProps{metadata}
}

func (p FooterProps) Site() domain.Site {
	return domain.DefaultSite
}

func (p FooterProps) Source() template.URL {
	if p.Metadata.SourceFile == "" {
		return ""
	}
	return template.URL(p.Site().BlogSourceRootURL + p.Metadata.SourceFile)
}

func (p FooterProps) Previous() domain.PartialMetadata {
	return p.Metadata.Previous
}

func (p FooterProps) Next() domain.PartialMetadata {
	return p.Metadata.Next
}
