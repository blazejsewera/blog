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
	log.Info("tailwind: downloading: upstreamFilename=%s", upstreamFilename)
	localFilename := execFilename()

	err := files.DownloadFile(upstreamURL(upstreamFilename), localFilename, true)
	if err != nil {
		panic(fmt.Errorf("tailwind: download: %w", err))
	}

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
