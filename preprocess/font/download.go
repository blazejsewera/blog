package font

import (
	"bufio"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/must"
	"log"
	"os"
	"path"
	"sync"
)

const (
	baseUpstreamURL = "https://www.sewera.dev/swty/"
	fontDir         = "_site/font"
	fontList        = "font.list"
)

func Download(force constants.ForceLevel) {
	fns := fontNames()
	if !allFontsExist(fns) || force >= constants.ReDownload {
		download(fns)
	}
}

func allFontsExist(fontNames []string) bool {
	for _, fontName := range fontNames {
		if !files.Exists(path.Join(fontDir, fontName)) {
			return false
		}
	}
	return true
}

func download(fontNames []string) {
	wg := &sync.WaitGroup{}
	for _, fontName := range fontNames {
		fontName := fontName
		wg.Add(1)
		go downloadFont(wg, fontName)
	}
	wg.Wait()
}

func fontNames() []string {
	file, err := os.Open(path.Join(fontDir, fontList))
	if err != nil {
		panic(err)
	}
	defer must.Close(file)

	var result []string
	s := bufio.NewScanner(file)
	for s.Scan() {
		fontName := s.Text()
		result = append(result, fontName)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return result
}

func downloadFont(wg *sync.WaitGroup, fontName string) {
	defer wg.Done()
	upstreamURL := baseUpstreamURL + fontName
	targetFile := path.Join(fontDir, fontName)
	err := files.DownloadFile(upstreamURL, targetFile, false)
	if err != nil {
		log.Printf("download font: %s", err)
	}
}
