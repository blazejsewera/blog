package tailwind

import (
	"fmt"
	"github.com/blazejsewera/blog/renderer/constants"
)

func upstreamExecFilename(osys string, arch string) string {
	const baseFilename = "tailwindcss"
	if osys == "windows" {
		return fmt.Sprintf("%s-%s-%s.exe", baseFilename, osys, arch)
	}
	return fmt.Sprintf("%s-%s-%s", baseFilename, osys, arch)
}

func upstreamExecURL(upstreamFilename string) string {
	return upstreamURL(upstreamFilename)
}

func upstreamChecksumsURL() string {
	const checksumsFilename = "sha256sums.txt"
	return upstreamURL(checksumsFilename)
}

func upstreamURL(filename string) string {
	return fmt.Sprintf("%s/%s/%s", constants.TailwindUpstreamURL, constants.TailwindVersion, filename)
}
