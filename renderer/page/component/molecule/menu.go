package molecule

import (
	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/page/component"
)

type MenuProps struct {
	Title string
}

func (p MenuProps) Site() component.SiteInfo {
	return component.DefaultSite
}

func MenuPropsOf(title string) MenuProps {
	return MenuProps{title}
}

func MenuPropsFromDomain(m article.Metadata) MenuProps {
	return MenuProps{m.Title}
}
