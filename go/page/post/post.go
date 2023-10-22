package post

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/article"
	"github.com/blazejsewera/blog/page/component/header"
)

var templateNames = templates.Pages("post/post")

type Template struct {
	t *templates.Template
}

type Props struct {
	Header  header.Props
	Article article.Props
}

func Post() *Template {
	return &Template{templates.ParseTFS(templateNames).
		With(header.Header).
		With(article.Article)}
}

func (t *Template) Render(props Props) []byte {
	return t.t.Render(props)
}
