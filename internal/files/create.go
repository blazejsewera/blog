package files

import (
	"fmt"
	"io/fs"
	"os"
)

func CreateDirIfDoesNotExist(name string) error {
	if !Exists(name) {
		err := os.Mkdir(name, 0755)
		if err != nil {
			return fmt.Errorf("dir: create: %w", err)
		}
	}
	return nil
}

func CreateFileWr(name string, executable bool) (*os.File, error) {
	var mode fs.FileMode
	if executable {
		mode = 0755
	} else {
		mode = 0644
	}

	file, err := os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
	if err != nil {
		return nil, fmt.Errorf("file: create for write: %w", err)
	}
	return file, nil
}
