package molecule

import "html/template"

type MenuProps struct {
	Title              string
	SiteName           string
	BlogSiteRootURL    template.URL
	BaseSiteRootURL    template.URL
	GithubProfileURL   template.URL
	MastodonProfileURL template.URL
}
