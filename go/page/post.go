package page

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/meta"
	"github.com/blazejsewera/blog/page/component/section"
)

type PostTemplate struct {
	t *templates.Template
}

type PostProps struct {
	Header  meta.HeaderProps
	Article section.ArticleProps
}

func Post() *PostTemplate {
	return &PostTemplate{templates.ParseAll()}
}

func (t *PostTemplate) Render(props PostProps) []byte {
	return t.t.Render(props)
}
