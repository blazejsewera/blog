package markdown

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/assert"
	"github.com/blazejsewera/blog/internal/files"
	"html/template"
	"strings"
	"testing"
)

func TestMarkdown(t *testing.T) {
	sampleMdWithFrontmatter := files.Read("parser_testdata/example.md")
	expectedHTML := files.Read("parser_testdata/example.html")

	title := "a title"
	sourceFile := "_site/test-article/index.md"
	targetFile := "dist/test-article/index.html"
	url := template.URL("/test-article")
	previous := domain.PartialMetadata{Title: "previous title"}

	parser := &Parser{
		AllArticles: []domain.ArticleMetadata{
			{
				Title:      title,
				SourceFile: sourceFile,
				TargetFile: targetFile,
				URL:        url,
				Previous:   previous,
			},
		},
	}

	t.Run("returns rendered HTML, metadata, target filename, and index.html-trimmed URL", func(t *testing.T) {
		input := strings.NewReader(sampleMdWithFrontmatter)
		output, metadata, targetFilename := parser.parseFile(input, sourceFile)

		expectedMetadata := domain.ArticleMetadata{
			Title:    title,
			URL:      url,
			Previous: previous,
		}
		expectedTargetFilename := "dist/test-article/index.html"

		assert.EqualFields(t, expectedMetadata, metadata, "Title", "URL", "Previous")
		assert.EqualHTML(t, expectedHTML, output)
		assert.Equal(t, expectedTargetFilename, targetFilename)
	})
}
