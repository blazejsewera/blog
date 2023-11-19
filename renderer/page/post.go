package page

import (
	"github.com/blazejsewera/blog/renderer/domain"
	"github.com/blazejsewera/blog/renderer/internal/templates"
	"github.com/blazejsewera/blog/renderer/page/component/meta"
	"github.com/blazejsewera/blog/renderer/page/component/section"
)

type PostTemplate struct {
	t *templates.Template
}

type PostProps struct {
	Metadata   domain.ArticleMetadata
	RawContent []byte
}

func (p PostProps) Header() meta.HeaderProps {
	return meta.HeaderPropsFromDomain(p.Metadata)
}

func (p PostProps) Site() domain.Site {
	return domain.DefaultSite
}

func (p PostProps) Title() section.TitleProps {
	return section.TitlePropsFromDomain(p.Metadata)
}

func (p PostProps) Article() section.ArticleProps {
	return section.ArticlePropsFromDomainAndRaw(p.Metadata, p.RawContent)
}

func (p PostProps) Footer() section.FooterProps {
	return section.FooterPropsFromDomain(p.Metadata)
}

func Post() *PostTemplate {
	return &PostTemplate{templates.ParseAll()}
}

func (t *PostTemplate) Render(props PostProps) ([]byte, error) {
	return t.t.Render("Post", props)
}
