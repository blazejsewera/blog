package domain

import "html/template"

type Site struct {
	Name               string
	BlogRootURL        template.URL
	BaseRootURL        template.URL
	GithubProfileURL   template.URL
	MastodonProfileURL template.URL
}

const (
	DefaultSiteName               = "blog.sewera.dev"
	DefaultSiteBlogRootURL        = "https://blog.sewera.dev"
	DefaultSiteBaseRootURL        = "https://www.sewera.dev"
	DefaultSiteGithubProfileURL   = "https://github.com/sewera"
	DefaultSiteMastodonProfileURL = "https://hachyderm.io/@sewera"
)

var DefaultSite = Site{
	Name:               DefaultSiteName,
	BlogRootURL:        DefaultSiteBlogRootURL,
	BaseRootURL:        DefaultSiteBaseRootURL,
	GithubProfileURL:   DefaultSiteGithubProfileURL,
	MastodonProfileURL: DefaultSiteMastodonProfileURL,
}
