package section

import (
	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/page/component"

	"html/template"
)

type FooterProps struct {
	Metadata article.Metadata
}

func FooterPropsFromDomain(metadata article.Metadata) FooterProps {
	return FooterProps{metadata}
}

func (p FooterProps) Site() component.SiteInfo {
	return component.DefaultSite
}

func (p FooterProps) Source() template.URL {
	if p.Metadata.SourceFile == "" {
		return ""
	}
	return template.URL(p.Site().BlogSourceRootURL + p.Metadata.SourceFile)
}

func (p FooterProps) Previous() article.PartialMetadata {
	return p.Metadata.Previous
}

func (p FooterProps) Next() article.PartialMetadata {
	return p.Metadata.Next
}
