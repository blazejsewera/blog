package frontmatter

import (
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/domain"
	"html/template"
	"path/filepath"
	"strings"
)

func ToArticleMetadata(f Frontmatter, markdownFilename string) domain.ArticleMetadata {
	sourceFile := markdownFilename
	url := urlFromMdFilename(markdownFilename)
	targetFile := targetFileFromMdFilename(markdownFilename)

	return domain.ArticleMetadata{
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
		Updates:          UpdatesToDomain(f.Updates),
	}
}

func UpdatesToDomain(uu []Update) []domain.Update {
	if uu == nil {
		return nil
	}
	result := make([]domain.Update, len(uu))
	for i, u := range uu {
		result[i] = u.ToDomain()
	}
	return result
}

func (u Update) ToDomain() domain.Update {
	return domain.Update{
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
