package tailwind

import (
	"github.com/blazejsewera/blog/internal/files"
	"os"
	"os/exec"
)

func Run() {
	downloadIfDoesNotExist()

	// TODO: execute tailwind
	tailwindCmd := exec.Command(execFilename(), "-h")
	tailwindCmd.Stdout = os.Stdout
	err := tailwindCmd.Run()
	if err != nil {
		panic(err)
	}
}

func downloadIfDoesNotExist() {
	if !files.Exists(execFilename()) {
		err := download()
		if err != nil {
			panic(err)
		}
	}
}
