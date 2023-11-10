package files

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func FindBySuffix(root, suffix string) (filePaths []string, err error) {
	err = filepath.WalkDir(root, func(path string, d fs.DirEntry, errW error) error {
		if errW != nil {
			return errW
		}
		if strings.HasSuffix(path, suffix) {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return filePaths, nil
}
