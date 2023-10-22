package post

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/page/component/article"
	"github.com/blazejsewera/blog/page/component/header"
)

func PropsFrom(htmlBytes []byte, css []byte, metadata domain.ArticleMetadata) Props {
	return Props{
		Header: header.Props{
			RawCSS:         css,
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
		Article: article.Props{
			Draft:            metadata.Draft,
			DraftDescription: metadata.DraftDescription,
			Abstract:         metadata.Abstract,
			RawContent:       htmlBytes,
		},
	}
}
