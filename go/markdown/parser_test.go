package markdown

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/assert"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/times"
	"html/template"
	"strings"
	"testing"
)

func TestMarkdown(t *testing.T) {
	sampleMdWithFrontmatter := files.Read("parser_testdata/example.md")
	expectedHTML := files.Read("parser_testdata/example.html")
	t.Run("returns rendered HTML, metadata, target filename, and index.html-trimmed URL", func(t *testing.T) {
		parser := Parser{
			WorkingDir: "dist",
			AllArticles: []domain.ArticleMetadata{
				{
					Title:      "a title",
					SourceFile: "_site/test-article/index.md",
					Previous:   domain.PartialMetadata{Title: "previous title"},
				},
			},
		}
		input := strings.NewReader(sampleMdWithFrontmatter)
		output, metadata, targetFilename := parser.parseFile(input, "dist/test-article/index.md")

		expectedMetadata := domain.ArticleMetadata{
			Title:    "a title",
			Date:     times.Parse("2023-03-03"),
			URL:      template.URL("/test-article"),
			Previous: domain.PartialMetadata{Title: "previous title"},
		}
		expectedTargetFilename := "dist/test-article/index.html"

		assert.EqualFields(t, expectedMetadata, metadata, "Title", "URL", "Previous")
		assert.Equal(t, expectedHTML, string(output))
		assert.Equal(t, expectedTargetFilename, targetFilename)
	})

	t.Run("returns rendered HTML, domain metadata, target filename, and non-trimmed URL", func(t *testing.T) {
		parser := Parser{WorkingDir: "dist"}
		input := strings.NewReader(sampleMdWithFrontmatter)
		_, metadata, targetFilename := parser.parseFile(input, "dist/test-article/article.md")

		expectedURL := template.URL("/test-article/article.html")
		expectedTargetFilename := "dist/test-article/article.html"

		assert.Equal(t, expectedURL, metadata.URL)
		assert.Equal(t, expectedTargetFilename, targetFilename)
	})
}
