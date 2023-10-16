package main

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/times"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/templates/header"
	"github.com/blazejsewera/blog/templates/index"
	"os"
)

func main() {
	t := index.Index()

	rendered := t.Render(index.Props{
		Header: header.Props{
			Title:    "Hello",
			Date:     times.Now(),
			Author:   "Me",
			License:  "License",
			Language: "en-US",
			SiteName: "blog.sewera.dev",
			Abstract: "This is an abstract",
			Keywords: []string{
				"hello",
				"world",
			},
			ImgURL:         "./some/url",
			ImgDescription: "Some picture",
		},
		Posts: []string{
			"first",
			"second",
		},
	})

	fmt.Printf("%s\n", rendered)

	input, err := os.Open("articles/ravioli-process.md")
	if err != nil {
		panic(err)
	}

	htmlBytes, frMetadata := markdown.Parse(input)
	fmt.Printf("%s\n", htmlBytes)
	fmt.Printf("%+#v\n", frMetadata)
	fmt.Printf("%s\n", frMetadata.Date)
}
