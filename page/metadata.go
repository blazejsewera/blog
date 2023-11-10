package page

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/page/component/section"
)

func PostPropsFrom(metadata domain.ArticleMetadata, rawContent []byte) PostProps {
	return PostProps{
		Metadata:   metadata,
		RawContent: rawContent,
	}
}

func IndexPropsFrom(metadata domain.ArticleMetadata, allArticles []domain.ArticleMetadata) IndexProps {
	return IndexProps{
		Metadata: metadata,
		Listing:  section.ListingPropsFromAllArticles(allArticles),
	}
}
