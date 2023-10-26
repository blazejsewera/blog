package molecule

import (
	"github.com/blazejsewera/blog/domain"
)

type MenuProps struct {
	Title string
}

func (p MenuProps) Site() domain.Site {
	return domain.DefaultSite
}

func MenuPropsOf(title string) MenuProps {
	return MenuProps{title}
}

func MenuPropsFromDomain(m domain.ArticleMetadata) MenuProps {
	return MenuProps{m.Title}
}
