package section

import (
	"github.com/blazejsewera/blog/domain"
	"html/template"
)

type ListingProps struct {
	Title          string
	TabTitle       string
	Subtitle       string
	Description    string
	ImgURL         template.URL
	ImgDescription string
	Articles       []domain.ArticleMetadata
}
