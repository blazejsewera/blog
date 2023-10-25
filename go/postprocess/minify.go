package postprocess

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"log"
)

const (
	mimeHTML = "text/html"
	mimeCSS  = "text/css"
)

var minifier = minify.New()

func init() {
	minifier.AddFunc(mimeHTML, html.Minify)
	minifier.AddFunc(mimeCSS, css.Minify)
}

func MinifyHTML(htmlBytes []byte) []byte {
	b, err := minifier.Bytes(mimeHTML, htmlBytes)
	if err != nil {
		log.Print("warn: could not minify HTML")
	}
	return b
}
