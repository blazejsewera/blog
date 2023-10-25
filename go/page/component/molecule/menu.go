package molecule

import (
	"github.com/blazejsewera/blog/constants"
	"html/template"
)

type MenuProps struct {
	Title              string
	SiteName           string
	BlogSiteRootURL    template.URL
	BaseSiteRootURL    template.URL
	GithubProfileURL   template.URL
	MastodonProfileURL template.URL
}

func MenuPropsWithTitle(title string) MenuProps {
	return MenuProps{
		Title:              title,
		SiteName:           constants.SiteName,
		BlogSiteRootURL:    constants.BlogSiteRootURL,
		BaseSiteRootURL:    constants.BaseSiteRootURL,
		GithubProfileURL:   constants.GithubProfileURL,
		MastodonProfileURL: constants.MastodonProfileURL,
	}
}
