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
	t.Run("returns rendered HTML, metadata, target filename, and index.html-trimmed URL", func(t *testing.T) {
		title := "a title"
		sourceFile := "_site/test-article/index.md"
		sourceInDist := "dist/test-article/index.md"
		targetFile := "dist/test-article/index.html"
		url := template.URL("/test-article")
		previous := domain.PartialMetadata{Title: "previous title"}

		parser := Parser{
			WorkingDir: "dist",
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
		input := strings.NewReader(sampleMdWithFrontmatter)
		output, metadata, targetFilename := parser.parseFile(input, sourceInDist)

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
