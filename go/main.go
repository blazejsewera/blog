package main

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/page/post"
	"github.com/blazejsewera/blog/preprocess/tailwind"
	"os"
)

func main() {
	css := tailwind.Run()

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
	rendered := t.Render(post.PropsFrom(htmlBytes, css, metadata))
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