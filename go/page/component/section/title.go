package section

import (
	"github.com/blazejsewera/blog/page/component/molecule"
	"html/template"
)

type TitleProps struct {
	Title    string
	Subtitle string
	ImgURL   template.URL
}

func (p TitleProps) Menu() molecule.MenuProps {
	return molecule.MenuPropsWithTitle(p.Title)
}
