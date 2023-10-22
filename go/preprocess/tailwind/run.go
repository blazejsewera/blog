package tailwind

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/must"
	"io"
	"os/exec"
)

const configFile = "preprocess/tailwind/tailwind.config.js"
const tailwindStyleFile = "_site/style/tailwind.css"

func Run(force bool) {
	if files.Exists(tailwindStyleFile) && !force {
		return
	}
	downloadIfDoesNotExist()

	cssBuf := &bytes.Buffer{}
	err := runTailwind(cssBuf)
	if err != nil {
		panic(err)
	}

	err = writeCSSFile(cssBuf)
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

func runTailwind(cssBuf *bytes.Buffer) error {
	tailwindCmd := exec.Command(execFilename(), "--config", configFile)
	tailwindCmd.Stdout = cssBuf
	err := tailwindCmd.Run()
	if err != nil {
		return fmt.Errorf("run tailwind: %w; maybe you have a wrong binary version for your OS/arch", err)
	}
	return nil
}

func writeCSSFile(cssBuf *bytes.Buffer) error {
	file, err := files.CreateFileWr(tailwindStyleFile, false)
	if err != nil {
		return err
	}
	defer must.Close(file)

	_, err = io.Copy(file, cssBuf)
	return err
}
