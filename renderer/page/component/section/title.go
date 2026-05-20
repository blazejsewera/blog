package section

import (
	"html/template"

	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/page/component/molecule"
)

type TitleProps struct {
	Title    string
	Subtitle string
	ImgURL   template.URL
	Metadata article.Metadata
}

func (p TitleProps) Menu() molecule.MenuProps {
	return molecule.MenuPropsFromDomain(p.Metadata)
}

func TitlePropsFromDomain(metadata article.Metadata) TitleProps {
	return TitleProps{
		Title:    metadata.Title,
		Subtitle: metadata.Subtitle,
		ImgURL:   metadata.ImgURL,
		Metadata: metadata,
	}
}
