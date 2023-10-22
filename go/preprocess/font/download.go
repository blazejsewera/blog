package font

import (
	"bufio"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/must"
	"os"
)

func DownloadIfAbsent() {
	if files.Exists("_site/font/SuisseIntl-Book-WebS.woff2") {
		return
	}
	download()
}

func download() {
	file, err := os.Open("_site/font/font.list")
	if err != nil {
		panic(err)
	}
	defer must.Close(file)

	var fontNames []string = []string{}

	s := bufio.NewScanner(file)
	for s.Scan() {
		fontName := s.Text()
		fontNames = append(fontNames, fontName)
	}

	if err := s.Err(); err != nil {
		panic(err)
	}
}
