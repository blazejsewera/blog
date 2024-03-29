package domain

import "html/template"

type Site struct {
	Name               string
	BlogRootURL        template.URL
	BaseRootURL        template.URL
	GithubProfileURL   template.URL
	MastodonProfileURL template.URL
	BlogSourceRootURL  string
}

const (
	DefaultSiteName               = "blog.sewera.dev"
	DefaultSiteBlogRootURL        = "/"
	DefaultSiteBaseRootURL        = "https://www.sewera.dev"
	DefaultSiteGithubProfileURL   = "https://github.com/sewera"
	DefaultSiteMastodonProfileURL = "https://hachyderm.io/@sewera"
	DefaultSiteBlogSourceRootURL  = "https://github.com/blazejsewera/blog/tree/master/"
)

var DefaultSite = Site{
	Name:               DefaultSiteName,
	BlogRootURL:        DefaultSiteBlogRootURL,
	BaseRootURL:        DefaultSiteBaseRootURL,
	GithubProfileURL:   DefaultSiteGithubProfileURL,
	MastodonProfileURL: DefaultSiteMastodonProfileURL,
	BlogSourceRootURL:  DefaultSiteBlogSourceRootURL,
}
