package tailwind

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/log"
	"github.com/blazejsewera/blog/internal/must"
	"io"
	"os/exec"
)

const configFile = "preprocess/tailwind/tailwind.config.js"
const tailwindStyleFile = "_site/style/tailwind.css"

func Run(force constants.ForceLevel) {
	if files.Exists(tailwindStyleFile) && force == constants.NoForce {
		return
	}
	if !files.Exists(execFilename()) || force >= constants.ReDownload {
		download()
	}

	log.Info("tailwind: running")
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

func runTailwind(cssBuf *bytes.Buffer) error {
	tailwindCmd := exec.Command(execFilename(), "--config", configFile, "--minify")
	tailwindCmd.Stdout = cssBuf
	err := tailwindCmd.Run()
	if err != nil {
		return fmt.Errorf("tailwind: run: %w; maybe you have a wrong binary version for your OS/arch", err)
	}
	return nil
}

func writeCSSFile(cssBuf *bytes.Buffer) error {
	file, err := files.CreateFileWr(tailwindStyleFile, false)
	if err != nil {
		return fmt.Errorf("tailwind: write css: %w", err)
	}
	defer must.Close(file)

	_, err = io.Copy(file, cssBuf)
	if err != nil {
		return fmt.Errorf("tailwind: write css: %w", err)
	}
	return nil
}
