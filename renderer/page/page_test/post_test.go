package page_test

import (
	"bytes"
	"github.com/blazejsewera/blog/renderer/internal/maps"
	"github.com/blazejsewera/blog/renderer/internal/templates"
	"github.com/blazejsewera/blog/renderer/internal/workingdir"
	"github.com/blazejsewera/blog/renderer/page"
	"testing"
)

func TestPost(t *testing.T) {
	workingdir.SetToProjectRoot()

	t.Run("renders properly", func(t *testing.T) {
		data := maps.Union(headerData, articleData)

		props := page.PostProps{
			Metadata:   articleMetadata,
			RawContent: rawContent,
		}

		rendered, err := page.Post().Render(props)

		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if bytes.Contains(rendered, templates.ErrorBytes) {
			t.Errorf(templates.ErrorMessage)
		}

		for _, v := range data {
			value := []byte(v)
			if !bytes.Contains(rendered, value) {
				t.Errorf("rendered does not contain '%s'", value)
			}
		}
	})
}
