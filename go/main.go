package main

import (
	"fmt"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/templates"
	"github.com/blazejsewera/blog/templates/header"
	"github.com/blazejsewera/blog/templates/index"
	"os"
)

func main() {
	_ = templates.New().
		With(index.Index).
		With(header.Header)

	input, err := os.Open("articles/ravioli-process.md")
	if err != nil {
		panic(err)
	}

	htmlBytes, frMetadata := markdown.Parse(input)
	fmt.Printf("%s\n", htmlBytes)
	fmt.Printf("%+#v\n", frMetadata)
	fmt.Printf("%s\n", frMetadata.Date)
}
