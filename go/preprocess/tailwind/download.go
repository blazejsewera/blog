package tailwind

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"os"
	"runtime"
)

const baseExecFilename = "bin/tailwindcss"

func download() error {
	osys, arch, err := detectOSAndArch()
	if err != nil {
		return err
	}
	upstreamFilename := upstreamExecFilename(osys, arch)
	localFilename := execFilename()

	err = files.DownloadFile(upstreamExecURL(upstreamFilename), localFilename, true)
	if err != nil {
		return err
	}

	err = checkSha256(upstreamFilename, localFilename)
	if err != nil {
		errRm := os.Remove(localFilename)
		if errRm != nil {
			return fmt.Errorf("remove local binary: %s: %w", localFilename, errRm)
		}
		return err
	}

	return nil
}

func execFilename() string {
	if runtime.GOOS == "windows" {
		return baseExecFilename + ".exe"
	} else {
		return baseExecFilename
	}
}
