package markdown

import (
	"bytes"
	"fmt"
	"io"
	"os"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"github.com/blazejsewera/blog/renderer/markdown/footnoteextension"
	"github.com/blazejsewera/blog/renderer/markdown/parse"
	"github.com/yuin/goldmark"
	goldmarkhighlighting "github.com/yuin/goldmark-highlighting/v2"
	goldmarkextension "github.com/yuin/goldmark/extension"
	goldmarkparser "github.com/yuin/goldmark/parser"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

type Renderer struct {
	AllArticles []article.Metadata
}

func (p *Renderer) RenderFile(markdownFilename string) (html []byte, metadata article.Metadata, targetFilename string) {
	file, err := os.Open(markdownFilename)
	if err != nil {
		panic(fmt.Errorf("markdown: parse: %w", err))
	}
	defer must.Close(file)
	log.Debug("markdown: parsing: %s", markdownFilename)
	return p.renderFileAndParseMetadata(file, markdownFilename)
}

func (p *Renderer) renderFileAndParseMetadata(markdownReader io.Reader, markdownFilename string) (html []byte, metadata article.Metadata, targetFilename string) {
	html = renderHtmlFromMarkdown(markdownReader)
	metadata = p.findMetadata(markdownFilename)
	log.Debug("markdown: parsed: %s", markdownFilename)
	return html, metadata, metadata.TargetFile
}

func (p *Renderer) findMetadata(markdownSourceFile string) article.Metadata {
	for _, anArticle := range p.AllArticles {
		if anArticle.EqualSource(markdownSourceFile) {
			return anArticle
		}
	}
	panic(fmt.Errorf("find metadata: cannot find metadata for %s; consider running Scanner first", markdownSourceFile))
}

func renderHtmlFromMarkdown(markdownReader io.Reader) (html []byte) {
	markdownBody := parse.MarkdownBody(markdownReader)

	md := goldmark.New(
		goldmark.WithParserOptions(
			goldmarkparser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(goldmarkhtml.WithUnsafe()),
		goldmark.WithExtensions(
			goldmarkextension.GFM,
			goldmarkextension.Typographer,
			goldmarkextension.DefinitionList,
			footnoteextension.NewFootnote(),
			goldmarkhighlighting.NewHighlighting(
				goldmarkhighlighting.WithStyle("xcode"),
				goldmarkhighlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
	)

	output := &bytes.Buffer{}
	err := md.Convert(markdownBody, output)
	if err != nil {
		panic(err)
	}

	return output.Bytes()
}
