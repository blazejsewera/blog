package postprocess

import (
	"github.com/yosssi/gohtml"
)

func MinifyHTML(htmlBytes []byte) []byte {
	return gohtml.FormatBytes(htmlBytes)
}
