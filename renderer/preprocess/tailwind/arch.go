package tailwind

import (
	"fmt"
	"runtime"

	"github.com/blazejsewera/blog/renderer/constants"
)

const (
	tailwindOSLinux   = "linux"
	tailwindOSWindows = "windows"
	tailwindOSMacOS   = "macos"

	tailwindArchX64   = "x64"
	tailwindArchArm64 = "arm64"
)

func detectOSAndArch() (osys string, arch string) {
	switch runtime.GOOS {
	case constants.OSLinux:
		osys = tailwindOSLinux
	case constants.OSMacOS:
		osys = tailwindOSMacOS
	case constants.OSWindows:
		osys = tailwindOSWindows
	default:
		panic(fmt.Errorf("tailwind: unsupported OS: %s", runtime.GOOS))
	}

	switch runtime.GOARCH {
	case constants.ArchAmd64, constants.ArchIntel64:
		arch = tailwindArchX64
	case constants.ArchArm64, constants.ArchArm:
		arch = tailwindArchArm64
	default:
		panic(fmt.Errorf("tailwind: unsupported architecture: %s", runtime.GOARCH))
	}

	return osys, arch
}
