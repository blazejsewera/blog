package section

import (
	"html/template"

	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/page/component/molecule"
)

type ArticleProps struct {
	Draft            bool
	DraftDescription string
	Abstract         string
	RawContent       []byte
}

func (p ArticleProps) Content() template.HTML {
	return template.HTML(p.RawContent)
}

func (p ArticleProps) DraftProps() molecule.DraftProps {
	return molecule.DraftProps{DraftDescription: p.DraftDescription}
}

func ArticlePropsFromDomainAndRaw(metadata article.Metadata, rawContent []byte) ArticleProps {
	return ArticleProps{
		Draft:            metadata.Draft,
		DraftDescription: metadata.DraftDescription,
		Abstract:         metadata.Abstract,
		RawContent:       rawContent,
	}
}
