package main

import (
	"fmt"
	"github.com/blazejsewera/blog/domain"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/times"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/page/component/header"
	"github.com/blazejsewera/blog/page/component/listing"
	"github.com/blazejsewera/blog/page/index"
	"github.com/blazejsewera/blog/postprocess"
	"os"
)

func main() {
	t := index.Index()

	_ = t.Render(index.Props{
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
		Listing: listing.Props{
			Articles: []domain.ArticleMetadata{
				{
					Title:          "Hello",
					Date:           times.Now(),
					URL:            "/hello",
					ImgURL:         "/hello.jpg",
					ImgDescription: "Hello img",
				},
				{
					Title:          "World",
					Date:           times.Now(),
					URL:            "/world",
					ImgURL:         "/world.jpg",
					ImgDescription: "World img",
				},
			},
		},
	})

	err := files.CopyDir("dist", "_site")
	if err != nil {
		panic(err)
	}

	postprocess.Run()

	input, err := os.Open("_site/ravioli-process/index.md")
	if err != nil {
		panic(err)
	}

	htmlBytes, frMetadata := markdown.Parse(input)
	fmt.Printf("%s\n", htmlBytes)
	fmt.Printf("%+#v\n", frMetadata)
	fmt.Printf("%s\n", frMetadata.Date)

	filePaths, err := files.FindBySuffix("dist", ".md")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", filePaths)
}
