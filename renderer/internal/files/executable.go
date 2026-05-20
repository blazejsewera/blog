package files

import (
	"runtime"

	"github.com/blazejsewera/blog/renderer/constants"
)

func ExecutableFilename(baseFilename string) string {
	if runtime.GOOS == constants.OSWindows {
		return baseFilename + ".exe"
	}
	return baseFilename
}
