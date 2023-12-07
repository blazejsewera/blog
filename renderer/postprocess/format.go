package postprocess

import (
	"github.com/yosssi/gohtml"
)

func FormatHTML(htmlBytes []byte) []byte {
	return gohtml.FormatBytes(htmlBytes)
}
