package files

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/blazejsewera/blog/renderer/internal/log"
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
	log.Debug("files: found files: %v; root: %s; suffix: %s", filePaths, root, suffix)
	return filePaths, nil
}
