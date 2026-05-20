package component

import (
	"html/template"

	"github.com/blazejsewera/blog/renderer/constants"
)

type SiteInfo struct {
	Name               string
	BlogRootURL        template.URL
	BaseRootURL        template.URL
	GithubProfileURL   template.URL
	MastodonProfileURL template.URL
	BlogSourceRootURL  string
}

var DefaultSite = SiteInfo{
	Name:               constants.DefaultSiteName,
	BlogRootURL:        constants.DefaultSiteBlogRootURL,
	BaseRootURL:        constants.DefaultSiteBaseRootURL,
	GithubProfileURL:   constants.DefaultSiteGithubProfileURL,
	MastodonProfileURL: constants.DefaultSiteMastodonProfileURL,
	BlogSourceRootURL:  constants.DefaultSiteBlogSourceRootURL,
}
