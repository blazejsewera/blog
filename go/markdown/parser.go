package markdown

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/markdown/footnoteextension"
	"io"
	"os"
	"path/filepath"
	"strings"

	chromahtml "github.com/alecthomas/chroma/v2/formatters/html"
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/must"
	"github.com/blazejsewera/blog/markdown/frontmatter"
	"github.com/yuin/goldmark"
	goldmarkhighlighting "github.com/yuin/goldmark-highlighting/v2"
	goldmarkextension "github.com/yuin/goldmark/extension"
	goldmarkhtml "github.com/yuin/goldmark/renderer/html"
)

type Parser struct {
	WorkingDir string
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
	html, frMetadata := parse(markdownReader)
	metadata = frMetadata.ToArticleMetadata(
		sourceFileFromMdFilename(markdownFilename),
		p.urlFromMdFilename(markdownFilename))
	return html, metadata, fileFromMdFilename(markdownFilename)
}

func (p *Parser) urlFromMdFilename(markdownFilename string) string {
	trimmed := strings.TrimPrefix(markdownFilename, p.WorkingDir)
	htmlFilename := fileFromMdFilename(trimmed)
	posix := filepath.ToSlash(htmlFilename)
	return strings.TrimSuffix(posix, "/index.html")
}

func sourceFileFromMdFilename(markdownFilename string) string {
	trimmed := strings.TrimPrefix(markdownFilename, constants.DistDir)
	return constants.SiteDir + trimmed
}

func fileFromMdFilename(markdownFilename string) string {
	trimmed := strings.TrimSuffix(markdownFilename, constants.MdExt)
	return trimmed + constants.HtmlExt
}

func parse(markdownReader io.Reader) (html []byte, frMetadata frontmatter.Frontmatter) {
	b := &bytes.Buffer{}
	_, err := io.Copy(b, markdownReader)
	if err != nil {
		panic(err)
	}

	frMetadata, stripped, frMetaExists := frontmatter.Unmarshal(b.Bytes())
	if !frMetaExists {
		frMetadata = frontmatter.DefaultFrMetadata
	}

	md := goldmark.New(
		goldmark.WithRendererOptions(goldmarkhtml.WithUnsafe()),
		goldmark.WithExtensions(
			goldmarkextension.GFM,
			goldmarkextension.Typographer,
			goldmarkextension.DefinitionList,
			footnoteextension.NewFootnote(),
			goldmarkhighlighting.NewHighlighting(
				goldmarkhighlighting.WithStyle("catppuccin-mocha"),
				goldmarkhighlighting.WithFormatOptions(
					chromahtml.WithLineNumbers(true),
				),
			),
		),
	)

	output := &bytes.Buffer{}
	err = md.Convert(stripped, output)
	if err != nil {
		panic(err)
	}

	return output.Bytes(), frMetadata
}
