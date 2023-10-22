package main

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/page/component/article"
	"github.com/blazejsewera/blog/page/component/draft"
	"github.com/blazejsewera/blog/page/component/header"
	"github.com/blazejsewera/blog/page/post"
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
	rendered := t.Render(post.Props{
		Header: header.Props{
			Title:          metadata.Title,
			Date:           metadata.Date,
			Author:         metadata.Author,
			License:        metadata.License,
			Language:       metadata.Language,
			SiteName:       "blog.sewera.dev",
			Abstract:       metadata.Abstract,
			Keywords:       metadata.Keywords,
			ImgURL:         metadata.ImgURL,
			ImgDescription: metadata.ImgDescription,
		},
		Article: article.Props{
			Draft:      metadata.Draft,
			DraftProps: draft.Props{DraftDescription: metadata.DraftDescription},
			Abstract:   metadata.Abstract,
			RawContent: htmlBytes,
		},
	})

	fmt.Printf("targetFilename: %s\n\n", targetFilename)

	fmt.Printf("%s\n", rendered)
}
