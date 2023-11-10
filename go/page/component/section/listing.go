package section

import (
	"github.com/blazejsewera/blog/domain"
	"slices"
)

type ListingProps struct {
	Articles []domain.ArticleMetadata
}

func ListingPropsFromAllArticles(allArticles []domain.ArticleMetadata) ListingProps {
	articles := slices.Clone(allArticles)
	slices.Reverse(articles)

	return ListingProps{articles}
}
