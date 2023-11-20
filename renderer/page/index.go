package page

import (
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/domain"
	"github.com/blazejsewera/blog/renderer/internal/templates"
	"github.com/blazejsewera/blog/renderer/internal/times"
	"github.com/blazejsewera/blog/renderer/page/component/meta"
	"github.com/blazejsewera/blog/renderer/page/component/section"
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

func (p IndexProps) Site() domain.Site {
	return domain.DefaultSite
}

func (p IndexProps) Title() section.TitleProps {
	return section.TitlePropsFromDomain(p.Metadata)
}

func (p IndexProps) Footer() section.FooterProps {
	return section.FooterPropsFromDomain(p.Metadata)
}

func Index() *IndexTemplate {
	return &IndexTemplate{templates.ParseAll()}
}

func (t *IndexTemplate) Render(props IndexProps) ([]byte, error) {
	return t.t.Render("Index", props)
}

var IndexMetadata = domain.FillDefaultIfEmpty(domain.ArticleMetadata{
	Title:      "Software development blog",
	Subtitle:   "By Blazej Sewera",
	Date:       times.Now(),
	URL:        "/",
	SourceFile: "renderer/page/index.html.tmpl",
	TargetFile: constants.DistDir + "/index.html",
	Keywords: []string{
		"software engineering",
		"software development",
		"programming",
		"technical blog",
		"Blazej Sewera",
	},
	ImgURL:         "/image/nordic-sky.jpg",
	ImgDescription: "Nordic sky",
})
