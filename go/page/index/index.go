package index

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/header"
	"github.com/blazejsewera/blog/page/component/listing"
)

var templateNames = templates.Pages("index/index")

type Template struct {
	t *templates.Template
}

type Props struct {
	Header  header.Props
	Listing listing.Props
}

func Index() *Template {
	return &Template{templates.ParseTFS(templateNames).
		With(header.Header).
		With(listing.Listing)}
}

func (t *Template) Render(props Props) []byte {
	return t.t.Render(props)
}
