package files

import (
	"errors"
	"github.com/blazejsewera/blog/internal/must"
	"io"
	"os"
	"path/filepath"
	"slices"
)

var skipFiles = []string{".git", ".gitignore", ".gitkeep"}

func CopyDir(dst, src string) error {
	entries, err1 := os.ReadDir(src)
	err2 := CreateDirIfDoesNotExist(dst)
	if err := errors.Join(err1, err2); err != nil {
		return err
	}

	for _, entry := range entries {
		if slices.Contains(skipFiles, entry.Name()) {
			continue
		}

		sourcePath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dst, entry.Name())

		fileInfo, err := os.Stat(sourcePath)
		if err != nil {
			return err
		}

		switch fileInfo.Mode() & os.ModeType {
		case os.ModeDir:
			if err = CreateDirIfDoesNotExist(destPath); err != nil {
				return err
			}
			if err = CopyDir(destPath, sourcePath); err != nil {
				return err
			}
		case os.ModeSymlink:
			if err = CopySymlink(destPath, sourcePath); err != nil {
				return err
			}
		default:
			if err = Copy(destPath, sourcePath); err != nil {
				return err
			}
		}
	}
	return nil
}

func Copy(dst, src string) error {
	out, err := CreateFileWr(dst, false)
	if err != nil {
		return err
	}
	defer must.Close(out)

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer must.Close(in)

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return nil
}

func CopySymlink(dst, src string) error {
	link, err := os.Readlink(src)
	if err != nil {
		return err
	}
	return os.Symlink(link, dst)
}
