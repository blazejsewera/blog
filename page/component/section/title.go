package section

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/page/component/molecule"
	"html/template"
)

type TitleProps struct {
	Title    string
	Subtitle string
	ImgURL   template.URL
	Metadata domain.ArticleMetadata
}

func (p TitleProps) Menu() molecule.MenuProps {
	return molecule.MenuPropsFromDomain(p.Metadata)
}

func TitlePropsFromDomain(metadata domain.ArticleMetadata) TitleProps {
	return TitleProps{
		Title:    metadata.Title,
		Subtitle: metadata.Subtitle,
		ImgURL:   metadata.ImgURL,
		Metadata: metadata,
	}
}
