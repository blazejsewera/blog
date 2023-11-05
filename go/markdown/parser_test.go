package markdown

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/times"
	"github.com/blazejsewera/blog/internal/workingdir"
	"github.com/stretchr/testify/assert"
	"html/template"
	"strings"
	"testing"
)

func TestMarkdown(t *testing.T) {
	workingdir.SetToProjectRoot()

	sampleMdWithFrontmatter := `---
title: a title
date: 2023-03-03
---

## H2 heading

Some text.

> Some quote[^1]

[^1]: Some footnote
`

	expectedHTML := `<h2>H2 heading</h2>
<p>Some text.</p>
<blockquote>
<p>Some quote<sup id="fnref:1"><a href="#fn:1" class="footnote-link" role="doc-noteref">1</a></sup></p>
</blockquote>
<div class="footnotes" role="doc-endnotes">
<h2 id="footnote-label">Footnotes</h2>
<ol>
<li id="fn:1">
<p>Some footnote&#160;<a href="#fnref:1" class="footnote-backlink" role="doc-backlink">[show in text]</a></p>
</li>
</ol>
</div>
`

	t.Run("returns rendered HTML, metadata, target filename, and index.html-trimmed URL", func(t *testing.T) {
		parser := Parser{WorkingDir: "dist"}
		input := strings.NewReader(sampleMdWithFrontmatter)
		output, metadata, targetFilename := parser.parseFile(input, "dist/test-article/index.md")

		expectedMetadata := domain.ArticleMetadata{
			Title: "a title",
			Date:  times.Parse("2023-03-03"),
			URL:   template.URL("/test-article"),
		}
		expectedTargetFilename := "dist/test-article/index.html"

		assert.Equal(t, expectedHTML, string(output))
		assert.Equal(t, expectedMetadata, metadata)
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
