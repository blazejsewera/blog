package page

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/meta"
	"github.com/blazejsewera/blog/page/component/section"
)

type IndexTemplate struct {
	t *templates.Template
}

type IndexProps struct {
	Metadata domain.ArticleMetadata
	Listing  section.ListingProps
}

func (p IndexProps) Header() meta.HeaderProps {
	return meta.HeaderPropsFromDomain(p.Metadata)
}

//goland:noinspection GoUnusedExportedFunction
func Index() *IndexTemplate {
	return &IndexTemplate{templates.ParseAll()}
}

func (t *IndexTemplate) Render(props IndexProps) ([]byte, error) {
	return t.t.Render("Index", props)
}
