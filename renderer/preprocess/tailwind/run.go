package tailwind

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
)

func Run(force constants.ForceLevel, verbosity constants.VerbosityLevel) {
	if files.Exists(constants.TailwindStyleFile) && force == constants.NoForce {
		return
	}
	if !files.Exists(execFilename()) || force >= constants.ReDownload {
		download()
	}

	err := runTailwind(verbosity)
	if err != nil {
		panic(err)
	}
	log.Info("tailwind: done")
}

func runTailwind(verbosity constants.VerbosityLevel) error {
	tailwindOutputOptimization := "--minify"
	if verbosity >= constants.Debug {
		tailwindOutputOptimization = ""
	}
	log.Debug("tailwind: running: version: %s; config: %s; optimization: %s",
		constants.TailwindVersion, constants.TailwindConfigFile, tailwindOutputOptimization)

	tailwindCmd := exec.Command(execFilename(),
		"--input", constants.TailwindConfigFile,
		"--output", constants.TailwindStyleFile,
		tailwindOutputOptimization)
	outBuf := &bytes.Buffer{}
	tailwindCmd.Stdout = outBuf
	errBuf := &bytes.Buffer{}
	tailwindCmd.Stderr = errBuf

	err := tailwindCmd.Run()
	if err != nil {
		return fmt.Errorf("tailwind: run: %w; maybe you have a wrong binary version for your OS/arch;\nstderr:\n%s", err, errBuf.String())
	}

	log.Debug("tailwind: stdout:\n%s", outBuf.String())
	log.Debug("tailwind: stderr:\n%s", errBuf.String())
	return nil
}
