package page_test

import (
	"bytes"
	"github.com/blazejsewera/blog/internal/maps"
	"github.com/blazejsewera/blog/internal/templates"
	"github.com/blazejsewera/blog/internal/workingdir"
	"github.com/blazejsewera/blog/page"
	"github.com/blazejsewera/blog/page/component/section"
	"testing"
)

func TestPost(t *testing.T) {
	workingdir.SetToProjectRoot()

	t.Run("renders properly", func(t *testing.T) {
		data := maps.Union(headerData, articleData)

		props := page.PostProps{
			Header: headerProps,
			Article: section.ArticleProps{
				Draft:            true,
				DraftDescription: data["draftDescription"],
				Abstract:         data["abstract"],
				RawContent:       rawContent,
			},
		}

		rendered := page.Post().Render(props)

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