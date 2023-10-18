package tailwind

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/must"
	"io"
	"io/fs"
	"net/http"
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

	err = downloadFile(upstreamExecURL(upstreamFilename), localFilename, true)
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

func downloadFile(url string, targetFile string, executable bool) error {
	var mode fs.FileMode
	if executable {
		mode = 0755
	} else {
		mode = 0644
	}

	file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	defer must.Close(file)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer must.Close(res.Body)

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return err
	}
	return nil
}
