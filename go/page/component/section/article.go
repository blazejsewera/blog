package section

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/molecule"
	"html/template"
)

func Article(t *templates.Template) {
	templateNames := templates.Components("section/article")
	t.With(molecule.Draft).ParseTFS(templateNames)
}

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
