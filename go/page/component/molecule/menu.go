package molecule

import (
	"github.com/blazejsewera/blog/internal/templates"
	"html/template"
)

func Menu(t *templates.Template) {
	templateNames := templates.Components("molecule/menu")
	t.ParseTFS(templateNames)
}

type MenuProps struct {
	Title              string
	SiteName           string
	BlogSiteRootURL    template.URL
	BaseSiteRootURL    template.URL
	GithubProfileURL   template.URL
	MastodonProfileURL template.URL
}
