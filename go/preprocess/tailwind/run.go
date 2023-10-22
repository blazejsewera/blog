package tailwind

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"os/exec"
)

const configFile = "preprocess/tailwind/tailwind.config.json"

func Run() (css []byte) {
	downloadIfDoesNotExist()

	cssBuf := &bytes.Buffer{}
	// TODO: execute tailwind
	tailwindCmd := exec.Command(execFilename(), "--config", configFile)
	tailwindCmd.Stdout = cssBuf
	err := tailwindCmd.Run()
	if err != nil {
		panic(fmt.Errorf("run tailwind: %w; maybe you have a wrong binary version for your OS/arch", err))
	}

	return cssBuf.Bytes()
}

func downloadIfDoesNotExist() {
	if !files.Exists(execFilename()) {
		err := download()
		if err != nil {
			panic(err)
		}
	}
}
