package workingdir

import (
	"fmt"
	"os"
	"path"
)

const maxIters = 6

func SetToProjectRoot() {
	if isProjectRoot(getCurrentWD()) {
		return
	}

	for i := 0; i < maxIters; i++ {
		parent := path.Dir(getCurrentWD())
		err := set(parent)
		if err != nil {
			panic(fmt.Errorf("set working directory: %w", err))
		}
		if isProjectRoot(parent) {
			return
		}
	}

	panic(fmt.Errorf("set working directory: failed to find project root after %d iterations", maxIters))
}

func getCurrentWD() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return path.Clean(wd)
}

func isProjectRoot(dir string) bool {
	if !isDir(dir) {
		return false
	}

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(fmt.Errorf("is project root: read dir: %s: %w", dir, err))
	}
	for _, file := range files {
		if file.Name() == "LICENSE" {
			return true
		}
	}
	return false
}

func isDir(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		panic(fmt.Errorf("is dir: %w", err))
	}
	return fileInfo.Mode()&os.ModeType == os.ModeDir
}

func set(newWD string) error {
	return os.Chdir(newWD)
}
