package markdown

import (
	"bytes"
	"github.com/blazejsewera/blog/markdown/frontmatter"
	"github.com/russross/blackfriday/v2"
	"io"
)

func Parse(markdownReader io.Reader) (html []byte, frMetadata frontmatter.Frontmatter) {
	b := &bytes.Buffer{}
	_, err := io.Copy(b, markdownReader)
	if err != nil {
		panic(err)
	}

	frMetadata, stripped, frMetaExists := frontmatter.Unmarshal(b.Bytes())
	if !frMetaExists {
		frMetadata = frontmatter.DefaultFrMetadata
	}

	rd := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.Smartypants |
			blackfriday.SmartypantsFractions |
			blackfriday.SmartypantsDashes |
			blackfriday.SmartypantsLatexDashes,
	})
	md := blackfriday.Run(stripped,
		blackfriday.WithRenderer(rd),
		blackfriday.WithExtensions(blackfriday.Footnotes))

	return md, frMetadata
}
