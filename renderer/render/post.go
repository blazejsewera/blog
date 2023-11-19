package render

import (
	"github.com/blazejsewera/blog/renderer/domain"
	"github.com/blazejsewera/blog/renderer/markdown"
	"github.com/blazejsewera/blog/renderer/page"
)

type PostRenderer struct {
	parser       *markdown.Parser
	postTemplate *page.PostTemplate
}

func NewPostRenderer(parser *markdown.Parser, postTemplate *page.PostTemplate) *PostRenderer {
	return &PostRenderer{parser, postTemplate}
}

func (r *PostRenderer) Render(sourceFile string) error {
	htmlBytes, metadata, targetFile := r.parser.ParseFile(sourceFile)
	rendered, err := r.postTemplate.Render(page.PostPropsFrom(domain.FillDefaultIfEmpty(metadata), htmlBytes))
	if err != nil {
		return renderErr(sourceFile, err)
	}
	return save(targetFile, rendered, sourceFile)
}
