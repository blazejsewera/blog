package tailwind

import (
	"fmt"
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
	"os"
	"runtime"
)

func download() {
	osys, arch, err := detectOSAndArch()
	if err != nil {
		panic(fmt.Errorf("tailwind: download: %w", err))
	}
	log.Info("tailwind: downloading: sys=%s arch=%s", osys, arch)
	upstreamFilename := upstreamExecFilename(osys, arch)
	localFilename := execFilename()

	err = files.DownloadFile(upstreamExecURL(upstreamFilename), localFilename, true)
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

//goland:noinspection GoBoolExpressions
func execFilename() string {
	if runtime.GOOS == "windows" {
		return constants.TailwindBinary + ".exe"
	} else {
		return constants.TailwindBinary
	}
}
