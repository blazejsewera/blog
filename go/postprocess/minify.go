package postprocess

import (
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"log"
)

const (
	mimeHTML = "text/html"
	mimeCSS  = "text/css"
	mimeJS   = "application/javascript"
)

var minifier = minify.New()

func init() {
	minifier.AddFunc(mimeHTML, html.Minify)
	minifier.AddFunc(mimeCSS, css.Minify)
	minifier.AddFunc(mimeJS, js.Minify)
}

func MinifyHTML(htmlBytes []byte) []byte {
	b, err := minifier.Bytes(mimeHTML, htmlBytes)
	if err != nil {
		log.Print("warn: could not minify HTML")
	}
	return b
}
