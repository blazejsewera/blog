package main

import (
	"flag"
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/page"
	"github.com/blazejsewera/blog/postprocess"
	"github.com/blazejsewera/blog/preprocess"
	"github.com/blazejsewera/blog/preprocess/distdir"
	"os"
)

func main() {
	f := flag.Int("f", 0, "Force level. -f0 is normal operation, -f1 re-renders project, -f2 re-downloads dependencies.")
	flag.Parse()

	force := constants.ForceLevel(*f)
	preprocess.Run(force)
	distdir.CopyIfDoesNotExist(force)

	filePaths, err := files.FindBySuffix(constants.DistDir, constants.MdExt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n\n", filePaths)

	parser := &markdown.Parser{WorkingDir: constants.DistDir}
	htmlBytes, metadata, targetFilename := parser.ParseFile(filePaths[0])
	t := page.Post()
	rendered := t.Render(page.PropsFrom(htmlBytes, metadata))
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
