package tailwind

import "fmt"

const (
	tailwindVersion = "v3.3.3"
	baseURL         = "https://github.com/tailwindlabs/tailwindcss/releases/download"
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
	return fmt.Sprintf("%s/%s/%s", baseURL, tailwindVersion, filename)
}
