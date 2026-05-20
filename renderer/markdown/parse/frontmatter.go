package parse

import (
	"bufio"
	"bytes"
	"io"
	"iter"
	"strings"

	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"github.com/blazejsewera/blog/renderer/internal/times"
	"gopkg.in/yaml.v3"
)

func Frontmatter(markdownReader io.Reader, markdownFilename string) article.Metadata {
	b := &bytes.Buffer{}
	must.Copy(b, markdownReader)
	frontmatter := readFrontmatter(b.Bytes())
	return mapToArticleMetadata(frontmatter, markdownFilename)
}

const frontmatterYamlDelimiter = "---"

type rawFrontmatter struct {
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
	Updates          []update   `yaml:"updates"`
}

type update struct {
	Date    times.Time `yaml:"date"`
	DiffURL string     `yaml:"diffUrl"`
}

var defaultFrontmatter = rawFrontmatter{
	Date:     times.Now(),
	Author:   "Blazej Sewera",
	License:  "CC-BY",
	Language: "en-US",
}

func readFrontmatter(wholeMarkdown []byte) (frontmatter rawFrontmatter) {
	if !frontmatterExists(wholeMarkdown) {
		return defaultFrontmatter
	}
	markdownReader := bufio.NewReader(bytes.NewReader(wholeMarkdown))

	frontmatterBuf := readFrontmatterBuf(markdownReader)
	frontmatter = rawFrontmatter{}
	err := yaml.Unmarshal(frontmatterBuf.Bytes(), &frontmatter)
	if err != nil {
		panic(err)
	}
	return frontmatter
}

func frontmatterExists(wholeMarkdown []byte) bool {
	reader := bufio.NewReader(bytes.NewReader(wholeMarkdown))
	head, _ := readLine(reader)
	return isFrontmatterYamlDelimiter(head)
}

func readFrontmatterBuf(markdownReader *bufio.Reader) *bytes.Buffer {
	firstLine, _ := readLine(markdownReader)
	if !isFrontmatterYamlDelimiter(firstLine) {
		panic("markdown: frontmatter was not detected correctly; check if you run frontmatterExists before parsing")
	}

	buf := &bytes.Buffer{}
	for line := range readFrontmatterLines(markdownReader) {
		appendLine(buf, line)
	}
	return buf
}

func appendLine(buf *bytes.Buffer, line string) {
	_, err := buf.WriteString(line)
	if err != nil {
		panic(err)
	}
}

func readFrontmatterLines(r *bufio.Reader) iter.Seq[string] {
	return func(yield func(s string) bool) {
		for {
			line, eof := readLine(r)
			if eof {
				return
			}
			if isFrontmatterYamlDelimiter(line) {
				return
			}
			if !yield(line) {
				return
			}
		}
	}
}

func readLine(r *bufio.Reader) (string, bool) {
	line, err := r.ReadString('\n')
	if err == io.EOF {
		return line, true
	}
	if err != nil {
		panic(err)
	}
	return line, false
}

func isFrontmatterYamlDelimiter(s string) bool {
	return strings.TrimSpace(s) == frontmatterYamlDelimiter
}
