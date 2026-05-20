package page

import (
	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/page/component/section"
)

func PostPropsFrom(metadata article.Metadata, rawContent []byte) PostProps {
	return PostProps{
		Metadata:   metadata,
		RawContent: rawContent,
	}
}

func IndexPropsFrom(metadata article.Metadata, allArticles []article.Metadata) IndexProps {
	return IndexProps{
		Metadata: metadata,
		Listing:  section.ListingPropsFromAllArticles(allArticles),
	}
}
