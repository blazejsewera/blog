package frontmatter

import (
	"bytes"
	"io"
)

func SplitMetadataAndMarkdown(markdownReader io.Reader) (frMetadata Frontmatter, stripped []byte) {
	b := &bytes.Buffer{}
	_, err := io.Copy(b, markdownReader)
	if err != nil {
		panic(err)
	}

	frMetadata, stripped, frMetaExists := Unmarshal(b.Bytes())
	if !frMetaExists {
		frMetadata = DefaultFrMetadata
	}
	return frMetadata, stripped
}
