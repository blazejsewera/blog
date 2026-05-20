package markdown

import (
	"fmt"
	"os"
	"slices"

	"github.com/blazejsewera/blog/renderer/article"
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"github.com/blazejsewera/blog/renderer/markdown/parse"
)

type Scanner struct {
	// WorkingDir is the base directory from which the Scanner starts looking for Markdown files.
	// It is constants.SiteDir by default.
	WorkingDir string
}

func (s *Scanner) ScanAllArticlesMetadata() (allArticles []article.Metadata, sourceFiles []string) {
	filePaths, err := files.FindBySuffix(s.workingDirectory(), constants.MdExt)
	if err != nil {
		panic(err)
	}
	var articles []article.Metadata
	for _, markdownFilename := range filePaths {
		articles = append(articles, scanFile(markdownFilename))
	}

	slices.SortFunc(articles, func(a, b article.Metadata) int {
		return a.Date.Compare(b.Date)
	})

	allArticles = linkArticlesCyclic(articles)
	sourceFiles = sources(allArticles)
	return allArticles, sourceFiles
}

func (s *Scanner) workingDirectory() string {
	if s.WorkingDir == "" {
		return constants.SiteDir
	}
	return s.WorkingDir
}

func scanFile(markdownFilename string) article.Metadata {
	file, err := os.Open(markdownFilename)
	if err != nil {
		panic(fmt.Errorf("markdown: parse file: %w", err))
	}
	defer must.Close(file)

	return parse.Frontmatter(file, markdownFilename)
}

func linkArticlesCyclic(articles []article.Metadata) []article.Metadata {
	limit := len(articles)

	previous := func(i int) article.Metadata {
		if i == 0 {
			return articles[limit-1]
		}
		return articles[i-1]
	}
	next := func(i int) article.Metadata {
		if i == limit-1 {
			return articles[0]
		}
		return articles[i+1]
	}
	for i := 0; i < limit; i++ {
		articles[i].Previous = article.PartialFromMetadata(previous(i))
		articles[i].Next = article.PartialFromMetadata(next(i))
	}
	return articles
}

func sources(articles []article.Metadata) []string {
	var result []string
	for _, anArticle := range articles {
		result = append(result, anArticle.SourceFile)
	}
	return result
}
