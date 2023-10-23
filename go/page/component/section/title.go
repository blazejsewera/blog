package section

import (
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/page/component/molecule"
	"html/template"
)

func Title(t *templates.Template) {
	templateNames := templates.Components("section/title")
	t.With(molecule.Draft).ParseTFS(templateNames)
}

type TitleProps struct {
	Title    string
	Subtitle string
	ImgURL   template.URL
}
