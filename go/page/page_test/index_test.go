package page_test

import (
	"bytes"
	"github.com/blazejsewera/blog/internal/maps"
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/internal/workingdir"
	"github.com/blazejsewera/blog/page"
	"github.com/blazejsewera/blog/page/component/section"
	"html/template"
	"testing"
)

func TestIndex(t *testing.T) {
	workingdir.SetToProjectRoot()

	t.Run("renders properly", func(t *testing.T) {
		data := maps.Union(headerData, listingData)
		props := page.IndexProps{
			Header: headerProps,
			Listing: section.ListingProps{
				Title:          data["title"],
				TabTitle:       data["tabTitle"],
				Subtitle:       data["subtitle"],
				Description:    data["description"],
				ImgURL:         template.URL(data["imgURL"]),
				ImgDescription: data["imgDescription"],
				Articles:       articles,
			},
		}

		rendered, err := page.Index().Render(props)

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
