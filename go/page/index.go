package page

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/meta"
	"github.com/blazejsewera/blog/page/component/section"
)

type IndexTemplate struct {
	t *templates.Template
}

type IndexProps struct {
	Header  meta.HeaderProps
	Listing section.ArticleProps
}

func Index() *IndexTemplate {
	templateNames := templates.Pages("index")
	return &IndexTemplate{templates.ParseTFS(templateNames).
		With(meta.Header).
		With(section.Listing)}
}

func (t *IndexTemplate) Render(props IndexProps) []byte {
	return t.t.Render(props)
}
