package markdown

import (
	"fmt"
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/domain"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"github.com/blazejsewera/blog/renderer/markdown/frontmatter"
	"os"
	"slices"
)

type Scanner struct {
	// WorkingDir is the base directory from which the Scanner starts looking for Markdown files.
	// It is _site by default.
	WorkingDir string
}

func (s *Scanner) ScanMetadata() (allArticles []domain.ArticleMetadata, sourceFiles []string) {
	filePaths, err := files.FindBySuffix(s.workingDirectory(), constants.MdExt)
	if err != nil {
		panic(err)
	}
	var articles []domain.ArticleMetadata
	for _, markdownFilename := range filePaths {
		articles = append(articles, scanFile(markdownFilename))
	}

	slices.SortFunc(articles, func(a, b domain.ArticleMetadata) int {
		return a.Date.Compare(b.Date)
	})

	allArticles = linkArticlesCyclic(articles)
	sourceFiles = sources(allArticles)
	return allArticles, sourceFiles
}

func (s *Scanner) workingDirectory() string {
	if s.WorkingDir == "" {
		return "_site"
	} else {
		return s.WorkingDir
	}
}

func scanFile(markdownFilename string) domain.ArticleMetadata {
	file, err := os.Open(markdownFilename)
	if err != nil {
		panic(fmt.Errorf("markdown: parse file: %w", err))
	}
	defer must.Close(file)

	frMetadata, _ := frontmatter.SplitMetadataAndMarkdown(file)
	return frontmatter.ToArticleMetadata(frMetadata, markdownFilename)
}

func linkArticlesCyclic(articles []domain.ArticleMetadata) []domain.ArticleMetadata {
	limit := len(articles)

	previous := func(i int) domain.ArticleMetadata {
		if i == 0 {
			return articles[limit-1]
		}
		return articles[i-1]
	}
	next := func(i int) domain.ArticleMetadata {
		if i == limit-1 {
			return articles[0]
		}
		return articles[i+1]
	}
	for i := 0; i < limit; i++ {
		articles[i].Previous = domain.PartialFromArticleMetadata(previous(i))
		articles[i].Next = domain.PartialFromArticleMetadata(next(i))
	}
	return articles
}

func sources(articles []domain.ArticleMetadata) []string {
	var result []string
	for _, article := range articles {
		result = append(result, article.SourceFile)
	}
	return result
}
