package render

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/page"
)

type IndexRenderer struct {
	indexTemplate *page.IndexTemplate
}

func NewIndexRenderer(indexTemplate *page.IndexTemplate) *IndexRenderer {
	return &IndexRenderer{indexTemplate}
}

func (r *IndexRenderer) Render(allArticles []domain.ArticleMetadata) error {
	metadata := page.IndexMetadata
	targetFile := metadata.TargetFile
	sourceFile := metadata.SourceFile

	rendered, err := r.indexTemplate.Render(page.IndexPropsFrom(metadata, allArticles))
	if err != nil {
		return renderErr(sourceFile, err)
	}
	return save(targetFile, rendered, sourceFile)
}
