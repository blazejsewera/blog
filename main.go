package main

import (
	"flag"
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/page"
	"github.com/blazejsewera/blog/postprocess"
	"github.com/blazejsewera/blog/preprocess"
	"github.com/blazejsewera/blog/preprocess/distdir"
	"os"
)

func main() {
	f := flag.Int("f", 0, "Force level. -f 0 is normal operation, -f 1 re-renders project, -f 2 re-downloads dependencies.")
	flag.Parse()

	force := constants.ForceLevel(*f)
	preprocess.Run(force)
	distdir.CopyIfDoesNotExist(force)

	scanner := markdown.Scanner{}
	allArticles, sourceFiles := scanner.ScanMetadata()

	postTemplate := page.Post()
	parser := &markdown.Parser{AllArticles: allArticles}
	for _, sourceFile := range sourceFiles {
		htmlBytes, metadata, targetFilename := parser.ParseFile(sourceFile)
		fmt.Printf("%s\n", metadata.URL)
		rendered, err := postTemplate.Render(page.PostPropsFrom(domain.FillDefaultIfEmpty(metadata), htmlBytes))
		if err != nil {
			panic(err)
		}
		target, err := os.Create(targetFilename)
		if err != nil {
			panic(err)
		}
		_, err = target.Write(postprocess.MinifyHTML(rendered))
		if err != nil {
			panic(err)
		}

		fmt.Printf("targetFilename: %s\n\n", targetFilename)
	}

	indexTemplate := page.Index()
	targetFilename := page.IndexMetadata.TargetFile
	rendered, err := indexTemplate.Render(page.IndexPropsFrom(page.IndexMetadata, allArticles))

	if err != nil {
		panic(err)
	}
	target, err := os.Create(targetFilename)
	if err != nil {
		panic(err)
	}
	_, err = target.Write(postprocess.MinifyHTML(rendered))
	if err != nil {
		panic(err)
	}

	fmt.Printf("targetFilename: %s\n\n", targetFilename)
}