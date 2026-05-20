package tailwind

import (
	"fmt"

	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
)

func upstreamURL(filename string) string {
	return fmt.Sprintf("%s/%s/%s", constants.TailwindUpstreamURL, constants.TailwindVersion, filename)
}

func upstreamExecFilename() string {
	osys, arch := detectOSAndArch()
	baseFilename := fmt.Sprintf("%s-%s-%s", constants.TailwindBaseFilename, osys, arch)
	return files.ExecutableFilename(baseFilename)
}

func upstreamChecksumsURL() string {
	const checksumsFilename = "sha256sums.txt"
	return upstreamURL(checksumsFilename)
}
