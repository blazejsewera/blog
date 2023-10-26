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
	Listing section.ListingProps
}

//goland:noinspection GoUnusedExportedFunction
func Index() *IndexTemplate {
	return &IndexTemplate{templates.ParseAll()}
}

func (t *IndexTemplate) Render(props IndexProps) ([]byte, error) {
	return t.t.Render("Index", props)
}
