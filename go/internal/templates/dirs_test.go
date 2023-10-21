package templates

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComponents(t *testing.T) {
	t.Run("return the relative template location of components given template name", func(t *testing.T) {
		assert.Equal(t,
			[]string{"component/header/header.html.tmpl", "component/header/base-header.html.tmpl"},
			Components("header/header", "header/base-header"))
	})
}

func TestPages(t *testing.T) {
	t.Run("return the relative template location of pages given template name", func(t *testing.T) {
		assert.Equal(t,
			[]string{"index/index.html.tmpl", "article/article.html.tmpl"},
			Pages("index/index", "article/article"))
	})
}
