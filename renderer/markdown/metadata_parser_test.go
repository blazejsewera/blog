package markdown

import (
	"testing"

	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/internal/assert"
	"github.com/blazejsewera/blog/renderer/internal/times"
)

func TestScanner(t *testing.T) {
	t.Run("scan metadata and sort it in correct order", func(t *testing.T) {
		s := Scanner{WorkingDir: "metadata_parser_testdata"}
		expected := []article.Metadata{
			{
				Title:    "First",
				Date:     times.MustParse("2023-01-01"),
				Previous: article.PartialMetadata{Title: "Last"},
				Next:     article.PartialMetadata{Title: "Middle"},
			},
			{
				Title:    "Middle",
				Date:     times.MustParse("2023-02-02"),
				Previous: article.PartialMetadata{Title: "First"},
				Next:     article.PartialMetadata{Title: "Almost last"},
			},
			{
				Title:    "Almost last",
				Date:     times.MustParse("2023-03-03"),
				Previous: article.PartialMetadata{Title: "Middle"},
				Next:     article.PartialMetadata{Title: "Last"},
			},
			{
				Title:    "Last",
				Date:     times.MustParse("2023-04-04"),
				Previous: article.PartialMetadata{Title: "Almost last"},
				Next:     article.PartialMetadata{Title: "First"},
			},
		}

		actual, _ := s.ScanMetadata()

		if assert.Len(t, 4, actual) {
			for i, ex := range expected {
				ac := actual[i]
				assert.Zero(t, ex.Date.Compare(ac.Date))
				assert.Equal(t, ex.Title, ac.Title)
				assert.Equal(t, ex.Previous.Title, ac.Previous.Title)
				assert.Equal(t, ex.Next.Title, ac.Next.Title)
			}
		}
	})
}
