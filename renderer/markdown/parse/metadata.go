package parse

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/constants"
)

func mapToArticleMetadata(f rawFrontmatter, markdownFilename string) article.Metadata {
	sourceFile := markdownFilename
	url := urlFromMdFilename(markdownFilename)
	targetFile := targetFileFromMdFilename(markdownFilename)

	return article.Metadata{
		Title:            f.Title,
		Subtitle:         f.Subtitle,
		Date:             f.Date,
		Draft:            f.Draft,
		DraftDescription: f.DraftDescription,
		Author:           f.Author,
		Abstract:         f.Abstract,
		Keywords:         f.Keywords,
		Language:         f.Language,
		License:          f.License,
		ImgURL:           template.URL(f.ImgURL),
		ImgDescription:   f.ImgDescription,
		URL:              template.URL(url),
		SourceFile:       sourceFile,
		TargetFile:       targetFile,
		Updates:          toArticleUpdates(f.Updates),
	}
}

func toArticleUpdates(uu []update) []article.Update {
	if uu == nil {
		return nil
	}
	result := make([]article.Update, len(uu))
	for i, u := range uu {
		result[i] = u.toArticleUpdate()
	}
	return result
}

func (u update) toArticleUpdate() article.Update {
	return article.Update{
		Date:    u.Date,
		DiffURL: u.DiffURL,
	}
}

func urlFromMdFilename(markdownFilename string) string {
	htmlFilename := bareTarget(markdownFilename)
	posix := filepath.ToSlash(htmlFilename)
	return strings.TrimSuffix(posix, "/index.html")
}

func targetFileFromMdFilename(markdownFilename string) string {
	return constants.DistDir + bareTarget(markdownFilename)
}

func bareTarget(markdownFilename string) string {
	trimmedOfSite := strings.TrimPrefix(markdownFilename, constants.SiteDir)
	trimmedOfExt := strings.TrimSuffix(trimmedOfSite, constants.MdExt)
	return trimmedOfExt + constants.HtmlExt
}
