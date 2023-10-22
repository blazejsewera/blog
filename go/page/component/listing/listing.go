package listing

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/draft"
	"html/template"
)

var templateNames = templates.Components("listing/listing")

func Listing(t *templates.Template) {
	t.With(draft.Draft).
		ParseTFS(templateNames)
}

type Props struct {
	Title          string
	TabTitle       string
	Subtitle       string
	Description    string
	ImgURL         template.URL
	ImgDescription string
	Articles       []domain.ArticleMetadata
}
