package page

import (
	"github.com/blazejsewera/blog/domain"
)

func PropsFrom(metadata domain.ArticleMetadata, rawContent []byte) PostProps {
	return PostProps{
		Metadata:   metadata,
		RawContent: rawContent,
	}
}
