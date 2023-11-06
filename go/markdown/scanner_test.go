package markdown

import (
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/times"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestScanner(t *testing.T) {
	t.Run("scan metadata and sort it in correct order", func(t *testing.T) {
		s := Scanner{WorkingDir: "scanner_testdata"}
		expected := []domain.ArticleMetadata{
			{
				Title:    "First",
				Date:     times.Parse("2023-01-01"),
				Previous: domain.PartialMetadata{Title: "Last"},
				Next:     domain.PartialMetadata{Title: "Middle"},
			},
			{
				Title:    "Middle",
				Date:     times.Parse("2023-02-02"),
				Previous: domain.PartialMetadata{Title: "First"},
				Next:     domain.PartialMetadata{Title: "Almost last"},
			},
			{
				Title:    "Almost last",
				Date:     times.Parse("2023-03-03"),
				Previous: domain.PartialMetadata{Title: "Middle"},
				Next:     domain.PartialMetadata{Title: "Last"},
			},
			{
				Title:    "Last",
				Date:     times.Parse("2023-04-04"),
				Previous: domain.PartialMetadata{Title: "Almost last"},
				Next:     domain.PartialMetadata{Title: "First"},
			},
		}

		actual := s.ScanMetadata()

		if assert.Len(t, actual, 4) {
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
