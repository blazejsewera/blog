package font

import (
	"path"
	"strings"
	"sync"

	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
)

func Download(force constants.ForceLevel) {
	fns := fontNames()
	if !allFontsExist(fns) || force >= constants.ReDownload {
		download(fns)
		log.Info("fonts: done")
	}
}

func fontNames() []string {
	fontNamesFile := files.Read(constants.FontListFile)
	var result []string
	strings.Split(fontNamesFile, "\n")
	for _, line := range strings.Split(fontNamesFile, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
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
	log.Debug("fonts: downloading: %v", fontNames)
	wg := &sync.WaitGroup{}
	for _, fontName := range fontNames {
		wg.Go(func() { downloadFont(fontName) })
	}
	wg.Wait()
}

func downloadFont(fontName string) {
	upstreamURL := constants.FontUpstreamURL + fontName
	targetFile := path.Join(constants.FontDir, fontName)
	err := files.DownloadFile(upstreamURL, targetFile, false)
	if err != nil {
		log.Error("font: download: %s", err)
		return
	}
	log.Debug("font: downloaded: source: %s; target: %s", upstreamURL, targetFile)
}
