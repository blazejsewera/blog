package frontmatter

import (
	"bytes"
	"io"
)

func ParseFrontmatter(markdownReader io.Reader) Frontmatter {
	frontmatter, _ := splitMetadataAndMarkdown(markdownReader)
	return frontmatter
}

func MarkdownBody(markdownReader io.Reader) []byte {
	_, markdownBody := splitMetadataAndMarkdown(markdownReader)
	return markdownBody
}

func splitMetadataAndMarkdown(markdownReader io.Reader) (frontmatter Frontmatter, markdownBody []byte) {
	b := &bytes.Buffer{}
	_, err := io.Copy(b, markdownReader)
	if err != nil {
		panic(err)
	}

	frontmatter, markdownBody = unmarshal(b.Bytes())
	return frontmatter, markdownBody
}
