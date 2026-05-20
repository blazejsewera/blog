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
	"github.com/blazejsewera/blog/renderer/markdown/frontmatter"
	"github.com/yuin/goldmark"
	goldmarkhighlighting "github.com/yuin/goldmark-highlighting/v2"
	goldmarkextension "github.com/yuin/goldmark/extension"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

type Parser struct {
	AllArticles []article.Metadata
}

func (p *Parser) ParseFile(markdownFilename string) (html []byte, metadata article.Metadata, targetFilename string) {
	file, err := os.Open(markdownFilename)
	if err != nil {
		panic(fmt.Errorf("markdown: parse: %w", err))
	}
	defer must.Close(file)
	log.Debug("markdown: parsing: %s", markdownFilename)
	return p.parseFile(file, markdownFilename)
}

func (p *Parser) parseFile(markdownReader io.Reader, markdownFilename string) (html []byte, metadata article.Metadata, targetFilename string) {
	html = parse(markdownReader)
	metadata = p.findMetadata(markdownFilename)
	log.Debug("markdown: parsed: %s", markdownFilename)
	return html, metadata, metadata.TargetFile
}

func (p *Parser) findMetadata(markdownSourceFile string) article.Metadata {
	for _, anArticle := range p.AllArticles {
		if anArticle.EqualSource(markdownSourceFile) {
			return anArticle
		}
	}
	panic(fmt.Errorf("find metadata: cannot find metadata for %s; consider running Scanner first", markdownSourceFile))
}

func parse(markdownReader io.Reader) (html []byte) {
	markdownBody := frontmatter.MarkdownBody(markdownReader)

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
	err := md.Convert(markdownBody, output)
	if err != nil {
		panic(err)
	}

	return output.Bytes()
}
