package section

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/page/component/molecule"
)

type TitleProps struct {
	Metadata domain.ArticleMetadata
}

func (p TitleProps) Menu() molecule.MenuProps {
	return molecule.MenuPropsFromDomain(p.Metadata)
}
