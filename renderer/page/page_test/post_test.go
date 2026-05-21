package page_test

import (
	"testing"

	"github.com/blazejsewera/blog/renderer/internal/maps"
	"github.com/blazejsewera/blog/renderer/internal/workingdir"
	"github.com/blazejsewera/blog/renderer/page"
)

func TestPost(t *testing.T) {
	workingdir.SetToProjectRoot()

	t.Run("renders header and article without errors", func(t *testing.T) {
		expectedData := maps.Union(headerData, articleData)

		props := page.PostProps{
			Metadata:   articleMetadata,
			RawContent: rawContent,
		}

		renderedHtml, err := page.Post().Render(props)

		assertRenderPageContains(t, renderedPage{renderedHtml, err}, expectedData)
	})
}
