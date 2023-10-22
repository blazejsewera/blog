package article

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/draft"
	"html/template"
)

var templateNames = templates.Components("article/article")

func Article(t *templates.Template) {
	t.With(draft.Draft).ParseTFS(templateNames)
}

type Props struct {
	Draft            bool
	DraftDescription string
	Abstract         string
	RawContent       []byte
}

func (p Props) Content() template.HTML {
	return template.HTML(p.RawContent)
}

func (p Props) DraftProps() draft.Props {
	return draft.Props{DraftDescription: p.DraftDescription}
}
