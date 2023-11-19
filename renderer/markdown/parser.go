package markdown

import (
	"bytes"
	"fmt"
	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/blazejsewera/blog/renderer/domain"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"github.com/blazejsewera/blog/renderer/markdown/footnoteextension"
	"github.com/blazejsewera/blog/renderer/markdown/frontmatter"
	"github.com/yuin/goldmark"
	goldmarkhighlighting "github.com/yuin/goldmark-highlighting/v2"
	goldmarkextension "github.com/yuin/goldmark/extension"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
	"io"
	"os"
)

type Parser struct {
	AllArticles []domain.ArticleMetadata
}

func (p *Parser) ParseFile(markdownFilename string) (html []byte, metadata domain.ArticleMetadata, targetFilename string) {
	file, err := os.Open(markdownFilename)
	if err != nil {
		panic(fmt.Errorf("markdown: parse file: %w", err))
	}
	defer must.Close(file)
	return p.parseFile(file, markdownFilename)
}

func (p *Parser) parseFile(markdownReader io.Reader, markdownFilename string) (html []byte, metadata domain.ArticleMetadata, targetFilename string) {
	html, _ = parse(markdownReader)
	metadata = p.findMetadata(markdownFilename)
	return html, metadata, metadata.TargetFile
}

func (p *Parser) findMetadata(markdownSourceFile string) domain.ArticleMetadata {
	for _, article := range p.AllArticles {
		if article.EqualSource(markdownSourceFile) {
			return article
		}
	}
	panic(fmt.Errorf("find metadata: cannot find metadata for %s; consider running Scanner first", markdownSourceFile))
}

func parse(markdownReader io.Reader) (html []byte, frMetadata frontmatter.Frontmatter) {
	frMetadata, stripped := frontmatter.SplitMetadataAndMarkdown(markdownReader)

	md := goldmark.New(
		goldmark.WithRendererOptions(goldmarkhtml.WithUnsafe()),
		goldmark.WithExtensions(
			goldmarkextension.GFM,
			goldmarkextension.Typographer,
			goldmarkextension.DefinitionList,
			footnoteextension.NewFootnote(),
			goldmarkhighlighting.NewHighlighting(
				goldmarkhighlighting.WithStyle("github"),
				goldmarkhighlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
	)

	output := &bytes.Buffer{}
	err := md.Convert(stripped, output)
	if err != nil {
		panic(err)
	}

	return output.Bytes(), frMetadata
}
