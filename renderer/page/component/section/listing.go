package section

import (
	"slices"

	"github.com/blazejsewera/blog/renderer/article"
)

type ListingProps struct {
	Articles []article.Metadata
}

func ListingPropsFromAllArticles(allArticles []article.Metadata) ListingProps {
	articles := slices.Clone(allArticles)
	slices.Reverse(articles)

	return ListingProps{articles}
}
