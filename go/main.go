package main

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/times"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/postprocess"
	"github.com/blazejsewera/blog/templates/header"
	"github.com/blazejsewera/blog/templates/index"
	"github.com/blazejsewera/blog/templates/listing"
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
		Listing: listing.Props{
			Articles: []listing.ArticleInfo{
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

	err := files.CreateDirIfDoesNotExist("dist")
	if err != nil {
		panic(err)
	}
	file, err := os.Create("dist/index.html")
	if err != nil {
		panic(err)
	}
	_, err = file.Write(rendered)
	if err != nil {
		panic(err)
	}

	err = files.CopyDir("dist", "_site")
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
}
