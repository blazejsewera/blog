package frontmatter

import (
	"github.com/blazejsewera/blog/constants"
	"path/filepath"
	"strings"
)

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
