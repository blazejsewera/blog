package main

import (
	"fmt"
	"github.com/blazejsewera/blog/markdown"
	"github.com/blazejsewera/blog/templates"
	"github.com/blazejsewera/blog/tmpl/header"
	"html/template"
	"os"
)

func main() {
	t := newTemplate(header.WithHeader)
	template.Must(t.ParseFS(templates.TemplateFS, "index.html.tmpl"))

	fmt.Println(t.DefinedTemplates())
	//err := t.Execute(os.Stdout, header.Header{
	//	Title:    "a title",
	//	Date:     times.Now(),
	//	Author:   "me",
	//	License:  "lic",
	//	Language: "en",
	//	SiteName: "site",
	//	Abstract: "abstract",
	//	Keywords: []string{"a", "b", "c"},
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	input, err := os.Open("articles/ravioli-process.md")
	if err != nil {
		panic(err)
	}

	htmlBytes, frMetadata, isFrontmatter := markdown.Parse(input)
	fmt.Printf("%s\n", htmlBytes)
	fmt.Printf("%+#v\n", frMetadata)
	fmt.Println(frMetadata.Date.ISODate())
	fmt.Println(isFrontmatter)
}

func newTemplate(decorators ...templates.TemplateDecorator) *template.Template {
	t := template.New("index.html.tmpl")

	for _, decorator := range decorators {
		decorator(t)
	}

	return t
}
