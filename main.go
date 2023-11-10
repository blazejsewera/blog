package main

import (
	"flag"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/page"
	"github.com/blazejsewera/blog/preprocess"
	"github.com/blazejsewera/blog/preprocess/distdir"
	"github.com/blazejsewera/blog/render"
)

func main() {
	force := parseCmdArgs()

	preprocess.Run(force)
	distdir.CopyIfDoesNotExist(force)

	scanner := markdown.Scanner{}
	allArticles, sourceFiles := scanner.ScanMetadata()

	parser := &markdown.Parser{AllArticles: allArticles}
	postTemplate := page.Post()
	postRenderer := render.NewPostRenderer(parser, postTemplate)
	for _, sourceFile := range sourceFiles {
		err := postRenderer.Render(sourceFile)
		if err != nil {
			panic(err)
		}
	}

	indexTemplate := page.Index()
	indexRenderer := render.NewIndexRenderer(indexTemplate)
	err := indexRenderer.Render(allArticles)
	if err != nil {
		panic(err)
	}
}

func parseCmdArgs() constants.ForceLevel {
	f := flag.Int("f", 0, "Force level. -f 0 is normal operation, -f 1 re-renders project, -f 2 re-downloads dependencies.")
	flag.Parse()

	force := constants.ForceLevel(*f)
	return force
}
