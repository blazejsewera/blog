package main

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/markdown"
)

func main() {
	err := files.CopyDir("dist", "_site")
	if err != nil {
		panic(err)
	}

	filePaths, err := files.FindBySuffix("dist", ".md")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", filePaths)

	parser := &markdown.Parser{WorkingDir: "dist"}
	htmlBytes, metadata, targetFilename := parser.ParseFile(filePaths[0])
	fmt.Printf("%s\n", htmlBytes)
	fmt.Printf("%+#v\n", metadata)
	fmt.Printf("%s\n", targetFilename)
}
