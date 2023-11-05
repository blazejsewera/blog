package markdown

import (
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/must"
	"github.com/blazejsewera/blog/markdown/frontmatter"
	"os"
	"slices"
)

type Scanner struct {
	WorkingDir string
}

func (s *Scanner) ScanMetadata() []domain.ArticleMetadata {
	filePaths, err := files.FindBySuffix(s.WorkingDir, constants.MdExt)
	if err != nil {
		panic(err)
	}
	var result []domain.ArticleMetadata
	for _, markdownFilename := range filePaths {
		result = append(result, scanFile(s.WorkingDir, markdownFilename))
	}

	slices.SortFunc(result, func(a, b domain.ArticleMetadata) int {
		return a.Date.Compare(b.Date)
	})
	return result
}

func scanFile(workingDir, markdownFilename string) domain.ArticleMetadata {
	file, err := os.Open(markdownFilename)
	if err != nil {
		panic(fmt.Errorf("markdown: parse file: %w", err))
	}
	defer must.Close(file)

	frMetadata, _ := frontmatter.SplitMetadataAndMarkdown(file)
	return frMetadata.ToArticleMetadata(workingDir, markdownFilename)
}
