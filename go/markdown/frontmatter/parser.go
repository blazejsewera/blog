package frontmatter

import (
	"bufio"
	"bytes"
	"errors"
	"io"
)

type Frontmatter map[string]string

func Unmarshal(markdown []byte) (frMetadata Frontmatter, stripped []byte, isFrontmatter bool) {
	br := bufio.NewReader(bytes.NewReader(markdown))
	if !detect(br) {
		return nil, markdown, false
	}

	inFrontmatter := false
	buf := &bytes.Buffer{}

	for {
		s, isPrefix, err := br.ReadLine()
		if err != nil {
			panic(err)
		} else if isPrefix {
			panic("read line: line was too long")
		}

		if string(s) == "---" {
			if inFrontmatter {
				// TODO: parse Frontmatter
				rest := &bytes.Buffer{}
				_, err := io.Copy(rest, br)
				if err != nil {
					panic(err)
				}
				return Frontmatter{"buf": buf.String()}, rest.Bytes(), true
			} else {
				inFrontmatter = true
			}
		} else {
			_, err1 := buf.Write(s)
			_, err2 := buf.WriteString("\n")
			if err := errors.Join(err1, err2); err != nil {
				panic(err)
			}
		}
	}

	return nil, markdown, false
}

func detect(br *bufio.Reader) bool {
	delimiterLength := 4
	head, err := br.Peek(delimiterLength)
	if err != nil {
		panic(err)
	}

	return string(head) == "---\n"
}
