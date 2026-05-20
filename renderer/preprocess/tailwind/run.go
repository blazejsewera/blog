package tailwind

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/internal/must"
)

func Run(force constants.ForceLevel) {
	if files.Exists(constants.TailwindStyleFile) && force == constants.NoForce {
		return
	}
	if !files.Exists(execFilename()) || force >= constants.ReDownload {
		download()
	}

	cssBuf := &bytes.Buffer{}
	err := runTailwind(cssBuf)
	if err != nil {
		panic(err)
	}

	err = writeCSSFile(cssBuf)
	if err != nil {
		panic(err)
	}
	log.Info("tailwind: done")
}

func runTailwind(cssBuf *bytes.Buffer) error {
	log.Debug("tailwind: running: version: %s; config: %s", constants.TailwindVersion, constants.TailwindConfigFile)
	tailwindCmd := exec.Command(execFilename(), "--config", constants.TailwindConfigFile, "--minify")
	tailwindCmd.Stdout = cssBuf
	errBuf := &bytes.Buffer{}
	tailwindCmd.Stderr = errBuf

	err := tailwindCmd.Run()
	if err != nil {
		return fmt.Errorf("tailwind: run: %w; maybe you have a wrong binary version for your OS/arch;\nstderr:\n%s", err, errBuf.String())
	}
	return nil
}

func writeCSSFile(cssBuf *bytes.Buffer) error {
	file, err := files.CreateFileWr(constants.TailwindStyleFile, false)
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
