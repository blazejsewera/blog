package page_test

import (
	"testing"

	"github.com/blazejsewera/blog/renderer/internal/maps"
	"github.com/blazejsewera/blog/renderer/internal/workingdir"
	"github.com/blazejsewera/blog/renderer/page"
	"github.com/blazejsewera/blog/renderer/page/component/section"
)

func TestIndex(t *testing.T) {
	workingdir.SetToProjectRoot()

	t.Run("renders header and listing without errors", func(t *testing.T) {
		expectedData := maps.Union(headerData, listingData)
		props := page.IndexProps{
			Metadata: articleMetadata,
			Listing:  section.ListingPropsFromAllArticles(articles),
		}

		renderedHtml, err := page.Index().Render(props)
		assertRenderPageContains(t, renderedPage{renderedHtml, err}, expectedData)
	})
}
