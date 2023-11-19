package tailwind

import (
	"fmt"
	"runtime"
)

func detectOSAndArch() (osys string, arch string, err error) {
	switch runtime.GOOS {
	case "linux":
		osys = "linux"
	case "darwin":
		osys = "macos"
	case "windows":
		osys = "windows"
	default:
		return "", "", fmt.Errorf("tailwind: unsupported OS: %s", runtime.GOOS)
	}

	switch runtime.GOARCH {
	case "amd64", "x86_64":
		arch = "x64"
	case "arm", "arm64":
		arch = "arm64"
	default:
		return "", "", fmt.Errorf("tailwind: unsupported architecture: %s", runtime.GOARCH)
	}

	return osys, arch, nil
}
