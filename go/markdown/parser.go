package markdown

import (
	"bytes"
	"github.com/blazejsewera/blog/markdown/frontmatter"
	"github.com/russross/blackfriday/v2"
	"io"
)

func Parse(htmlWriter io.Writer, markdownReader io.Reader) (frMetadata frontmatter.Frontmatter, isFrontmatter bool) {
	b := &bytes.Buffer{}
	_, err := io.Copy(b, markdownReader)
	if err != nil {
		panic(err)
	}

	frMetadata, stripped, isFrontmatter := frontmatter.Unmarshal(b.Bytes())

	rd := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.Smartypants |
			blackfriday.SmartypantsFractions |
			blackfriday.SmartypantsDashes |
			blackfriday.SmartypantsLatexDashes,
	})
	md := blackfriday.Run(stripped,
		blackfriday.WithRenderer(rd),
		blackfriday.WithExtensions(blackfriday.Footnotes))

	_, err = htmlWriter.Write(md)
	if err != nil {
		panic(err)
	}

	return frMetadata, isFrontmatter
}
