package frontmatter

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/blazejsewera/blog/internal/times"
	"gopkg.in/yaml.v3"
	"io"
)

type Frontmatter struct {
	Layout           string     `yaml:"layout"`
	Title            string     `yaml:"title"`
	Subtitle         string     `yaml:"subtitle"`
	Date             times.Time `yaml:"date"`
	Author           string     `yaml:"author"`
	License          string     `yaml:"license"`
	Language         string     `yaml:"language"`
	Draft            bool       `yaml:"draft"`
	DraftDescription string     `yaml:"draftDescription"`
	ImgURL           string     `yaml:"imgUrl"`
	ImgDescription   string     `yaml:"imgDescription"`
	Abstract         string     `yaml:"abstract"`
	Keywords         []string   `yaml:"keywords"`
}

var DefaultFrMetadata = Frontmatter{
	Layout:   "Article",
	Date:     times.Now(),
	Author:   "Blazej Sewera",
	License:  "CC-BY",
	Language: "en-US",
}

func Unmarshal(markdown []byte) (frMetadata Frontmatter, stripped []byte, frMetaExists bool) {
	br := bufio.NewReader(bytes.NewReader(markdown))
	if !detect(br) {
		return Frontmatter{}, markdown, false
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
				frMetadata = Frontmatter{}
				err := yaml.Unmarshal(buf.Bytes(), &frMetadata)
				if err != nil {
					panic(err)
				}

				rest := &bytes.Buffer{}
				_, err = io.Copy(rest, br)
				if err != nil {
					panic(err)
				}
				return frMetadata, rest.Bytes(), true
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

	return Frontmatter{}, markdown, false
}

func detect(br *bufio.Reader) bool {
	delimiterLength := 4
	head, err := br.Peek(delimiterLength)
	if err != nil {
		panic(err)
	}

	return string(head) == "---\n"
}
