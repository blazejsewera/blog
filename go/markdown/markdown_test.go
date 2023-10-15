package markdown_test

import (
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/markdown/frontmatter"
	"github.com/blazejsewera/blog/times"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

// language=Markdown
var sampleMd = `
## H2 heading

Some text
`

// language=YAML
var sampleMdWithFrontmatter = `---
layout: a layout
title: a title
date: 2023-03-03
---
` + sampleMd

// language=HTML
const expectedHTML = `<h2>H2 heading</h2>

<p>Some text</p>
`

var expectedFrMetadata = frontmatter.Frontmatter{
	Layout: "a layout",
	Title:  "a title",
	Date:   times.Parse("2023-03-03"),
}

func TestMarkdown(t *testing.T) {
	t.Run("returns rendered HTML and parsed metadata from Markdown with Frontmatter", func(t *testing.T) {
		input := strings.NewReader(sampleMdWithFrontmatter)
		output, actualFrMetadata := markdown.Parse(input)

		assert.Equal(t, expectedHTML, string(output))
		assert.Equal(t, expectedFrMetadata, actualFrMetadata)
	})

	t.Run("returns rendered HTML and default metadata from Markdown without Frontmatter", func(t *testing.T) {
		input := strings.NewReader(sampleMd)
		output, actualFrMetadata := markdown.Parse(input)

		assert.Equal(t, expectedHTML, string(output))
		assert.Equal(t, frontmatter.DefaultFrMetadata, actualFrMetadata)
	})
}
