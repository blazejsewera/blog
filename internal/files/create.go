package files

import (
	"io/fs"
	"os"
)

func CreateDirIfDoesNotExist(name string) error {
	if !Exists(name) {
		err := os.Mkdir(name, 0755)
		if err != nil {
			return err
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

	return os.OpenFile(name, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, mode)
}
