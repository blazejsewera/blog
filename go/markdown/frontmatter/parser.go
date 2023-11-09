package frontmatter

import (
	"bufio"
	"bytes"
	"errors"
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/times"
	"gopkg.in/yaml.v3"
	"html/template"
	"io"
)

type Frontmatter struct {
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
	Updates          []Update   `yaml:"updates"`
}

type Update struct {
	Date    times.Time `yaml:"date"`
	DiffURL string     `yaml:"diffUrl"`
}

var DefaultFrMetadata = Frontmatter{
	Date:     times.Now(),
	Author:   "Blazej Sewera",
	License:  "CC-BY",
	Language: "en-US",
}

func ToArticleMetadata(f Frontmatter, markdownFilename string) domain.ArticleMetadata {
	sourceFile := markdownFilename
	url := urlFromMdFilename(markdownFilename)
	targetFile := targetFileFromMdFilename(markdownFilename)

	return domain.ArticleMetadata{
		Title:            f.Title,
		Subtitle:         f.Subtitle,
		Date:             f.Date,
		Draft:            f.Draft,
		DraftDescription: f.DraftDescription,
		Author:           f.Author,
		Abstract:         f.Abstract,
		Keywords:         f.Keywords,
		Language:         f.Language,
		License:          f.License,
		ImgURL:           template.URL(f.ImgURL),
		ImgDescription:   f.ImgDescription,
		URL:              template.URL(url),
		SourceFile:       sourceFile,
		TargetFile:       targetFile,
		Updates:          UpdatesToDomain(f.Updates),
	}
}

func UpdatesToDomain(uu []Update) []domain.Update {
	if uu == nil {
		return nil
	}
	result := make([]domain.Update, len(uu))
	for i, u := range uu {
		result[i] = u.ToDomain()
	}
	return result
}

func (u Update) ToDomain() domain.Update {
	return domain.Update{
		Date:    u.Date,
		DiffURL: u.DiffURL,
	}
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
