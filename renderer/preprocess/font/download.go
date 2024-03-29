package font

import (
	"bufio"
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"os"
	"path"
	"sync"
)

func Download(force constants.ForceLevel) {
	fns := fontNames()
	if !allFontsExist(fns) || force >= constants.ReDownload {
		download(fns)
	}
}

func allFontsExist(fontNames []string) bool {
	for _, fontName := range fontNames {
		if !files.Exists(path.Join(constants.FontDir, fontName)) {
			return false
		}
	}
	return true
}

func download(fontNames []string) {
	log.Info("fonts: downloading")
	wg := &sync.WaitGroup{}
	for _, fontName := range fontNames {
		fontName := fontName
		wg.Add(1)
		go downloadFont(wg, fontName)
	}
	wg.Wait()
}

func fontNames() []string {
	file, err := os.Open(path.Join(constants.FontDir, constants.FontList))
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
	upstreamURL := constants.FontUpstreamURL + fontName
	targetFile := path.Join(constants.FontDir, fontName)
	err := files.DownloadFile(upstreamURL, targetFile, false)
	if err != nil {
		log.Error("font: download: %s", err)
	}
}
