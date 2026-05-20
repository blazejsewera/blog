package render

import (
	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/markdown"
	"github.com/blazejsewera/blog/renderer/page"
)

type PostRenderer struct {
	renderer     *markdown.Renderer
	postTemplate *page.PostTemplate
}

func NewPostRenderer(parser *markdown.Renderer, postTemplate *page.PostTemplate) *PostRenderer {
	return &PostRenderer{parser, postTemplate}
}

func (r *PostRenderer) Render(sourceFile string) error {
	log.Debug("render: rendering post")
	htmlBytes, metadata, targetFile := r.renderer.RenderFile(sourceFile)
	rendered, err := r.postTemplate.Render(page.PostPropsFrom(article.FillDefaultIfEmpty(metadata), htmlBytes))
	if err != nil {
		return renderErr(sourceFile, err)
	}
	return save(targetFile, rendered, sourceFile)
}
