package section

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/molecule"
	"html/template"
)

func Listing(t *templates.Template) {
	templateNames := templates.Components("section/listing/listing")
	t.With(molecule.Draft).
		ParseTFS(templateNames)
}

type ListingProps struct {
	Title          string
	TabTitle       string
	Subtitle       string
	Description    string
	ImgURL         template.URL
	ImgDescription string
	Articles       []domain.ArticleMetadata
}
