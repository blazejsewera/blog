package frontmatter

import (
	"github.com/blazejsewera/blog/constants"
	"path/filepath"
	"strings"
)

func urlFromMdFilename(workingDir, markdownFilename string) string {
	trimmed := strings.TrimPrefix(markdownFilename, workingDir)
	htmlFilename := targetFileFromMdFilename(trimmed)
	posix := filepath.ToSlash(htmlFilename)
	return strings.TrimSuffix(posix, "/index.html")
}

func sourceFileFromMdFilename(markdownFilename string) string {
	trimmed := strings.TrimPrefix(markdownFilename, constants.DistDir)
	return constants.SiteDir + trimmed
}

func targetFileFromMdFilename(markdownFilename string) string {
	trimmed := strings.TrimSuffix(markdownFilename, constants.MdExt)
	return trimmed + constants.HtmlExt
}
