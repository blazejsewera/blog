package section

import (
	"github.com/blazejsewera/blog/page/component/molecule"
	"html/template"
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
