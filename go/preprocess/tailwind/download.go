package tailwind

import (
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"log"
	"os"
	"runtime"
)

const baseExecFilename = "bin/tailwindcss"

func download() {
	log.Printf("info: downloading tailwind")
	osys, arch, err := detectOSAndArch()
	if err != nil {
		panic(fmt.Errorf("download tailwind: %w", err))
	}
	upstreamFilename := upstreamExecFilename(osys, arch)
	localFilename := execFilename()

	err = files.DownloadFile(upstreamExecURL(upstreamFilename), localFilename, true)
	if err != nil {
		panic(fmt.Errorf("download tailwind: %w", err))
	}

	err = checkSha256(upstreamFilename, localFilename)
	if err != nil {
		errRm := os.Remove(localFilename)
		if errRm != nil {
			panic(fmt.Errorf("remove local binary: %s: %w", localFilename, errRm))
		}
		panic(fmt.Errorf("download tailwind: %w", err))
	}

}

//goland:noinspection GoBoolExpressions
func execFilename() string {
	if runtime.GOOS == "windows" {
		return baseExecFilename + ".exe"
	} else {
		return baseExecFilename
	}
}
