package parse

import (
	"bufio"
	"bytes"
	"io"

	"github.com/blazejsewera/blog/renderer/internal/must"
)

func MarkdownBody(markdownReader io.Reader) []byte {
	markdownBuf := &bytes.Buffer{}
	must.Copy(markdownBuf, markdownReader)
	markdownBody := readMarkdownBody(markdownBuf.Bytes())
	return markdownBody
}

func readMarkdownBody(wholeMarkdown []byte) []byte {
	if !frontmatterExists(wholeMarkdown) {
		return wholeMarkdown
	}

	markdownReader := bufio.NewReader(bytes.NewReader(wholeMarkdown))
	skipFrontmatter(markdownReader)
	restOfMarkdown := &bytes.Buffer{}
	must.Copy(restOfMarkdown, markdownReader)
	return restOfMarkdown.Bytes()
}

func skipFrontmatter(markdownReader *bufio.Reader) {
	readLine(markdownReader)
	for range readFrontmatterLines(markdownReader) {
	}
}
