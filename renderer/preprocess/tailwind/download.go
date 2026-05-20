package tailwind

import (
	"fmt"
	"os"

	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
)

func download() {
	upstreamFilename := upstreamExecFilename()
	upstreamUrl := upstreamURL(upstreamFilename)
	log.Debug("tailwind: downloading: source: %s", upstreamUrl)

	localFilename := execFilename()
	err := files.DownloadFile(upstreamUrl, localFilename, true)
	if err != nil {
		panic(fmt.Errorf("tailwind: download: %w", err))
	}
	log.Debug("tailwind: downloaded: source: %s; target: %s", upstreamUrl, localFilename)

	err = checkSha256(upstreamFilename, localFilename)
	if err != nil {
		errRm := os.Remove(localFilename)
		if errRm != nil {
			panic(fmt.Errorf("tailwind: remove local binary: %s: %w", localFilename, errRm))
		}
		panic(fmt.Errorf("tailwind: download: %w", err))
	}

}

func execFilename() string {
	return files.ExecutableFilename(constants.TailwindBinary)
}
