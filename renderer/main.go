package main

import (
	"flag"
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/markdown"
	"github.com/blazejsewera/blog/renderer/page"
	"github.com/blazejsewera/blog/renderer/preprocess"
	"github.com/blazejsewera/blog/renderer/preprocess/distdir"
	"github.com/blazejsewera/blog/renderer/render"
)

func main() {
	force, verbosity := parseCmdArgs()
	log.SetVerbosity(verbosity)

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

	log.Info("success")
}

func parseCmdArgs() (constants.ForceLevel, constants.VerbosityLevel) {
	const forceLevelDescription = `Force level.
	-f=0 is normal operation,
	-f=1 re-runs tailwind and re-renders project,
	-f=2 removes dist and does the above,
	-f=3 re-downloads dependencies and does the above.`
	f := flag.Int("f", 0, forceLevelDescription)

	const verbosityLevelDescription = `Verbosity level.
	-v=0 silent,
	-v=1 errors only,
	-v=2 warnings and above,
	-v=3 info and above,
	-v=4 debug and above`
	v := flag.Int("v", 0, verbosityLevelDescription)

	flag.Parse()
	force := constants.ForceLevel(*f)
	verbosity := constants.VerbosityLevel(*v)
	return force, verbosity
}
