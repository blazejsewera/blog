package files

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"slices"

	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/internal/must"
)

var skipFiles = []string{".git", ".gitignore", ".gitkeep"}

func CopyDir(dst, src string) error {
	log.Debug("files: copying directory: source: %s; target: %s", src, dst)
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
			return fmt.Errorf("files: copy directory: %s", err)
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
				return fmt.Errorf("files: copy directory: %w", err)
			}
		default:
			if err = Copy(destPath, sourcePath); err != nil {
				return fmt.Errorf("files: copy directory: %w", err)
			}
		}
	}
	return nil
}

func Copy(dst, src string) error {
	log.Debug("files: copying file: source: %s; target: %s", src, dst)
	out, err := CreateFileWr(dst, false)
	if err != nil {
		return err
	}
	defer must.Close(out)

	in, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("files: copy file: open: %w", err)
	}
	defer must.Close(in)

	_, err = io.Copy(out, in)
	if err != nil {
		return fmt.Errorf("files: copy file: %w", err)
	}
	return nil
}

func CopySymlink(dst, src string) error {
	log.Debug("files: copying symlink: source: %s; target: %s", src, dst)
	link, err := os.Readlink(src)
	if err != nil {
		return fmt.Errorf("files: copy symlink: %w", err)
	}
	return os.Symlink(link, dst)
}
