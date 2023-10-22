package main

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/page/post"
	"os"
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
	fmt.Printf("%v\n\n", filePaths)

	parser := &markdown.Parser{WorkingDir: "dist"}
	htmlBytes, metadata, targetFilename := parser.ParseFile(filePaths[0])
	t := post.Post()
	rendered := t.Render(post.PropsFrom(htmlBytes, metadata))
	target, err := os.Create(targetFilename)
	if err != nil {
		panic(err)
	}
	_, err = target.Write(rendered)
	if err != nil {
		panic(err)
	}

	fmt.Printf("targetFilename: %s\n\n", targetFilename)
}
