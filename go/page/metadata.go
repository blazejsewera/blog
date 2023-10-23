package page

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/page/component/meta"
	"github.com/blazejsewera/blog/page/component/section"
)

func PropsFrom(htmlBytes []byte, metadata domain.ArticleMetadata) PostProps {
	return PostProps{
		Header: meta.HeaderProps{
			Title:          metadata.Title,
			Date:           metadata.Date,
			Author:         metadata.Author,
			License:        metadata.License,
			Language:       metadata.Language,
			SiteName:       "blog.sewera.dev",
			Abstract:       metadata.Abstract,
			Keywords:       metadata.Keywords,
			ImgURL:         metadata.ImgURL,
			ImgDescription: metadata.ImgDescription,
		},
		Article: section.ArticleProps{
			Draft:            metadata.Draft,
			DraftDescription: metadata.DraftDescription,
			Abstract:         metadata.Abstract,
			RawContent:       htmlBytes,
		},
	}
}
