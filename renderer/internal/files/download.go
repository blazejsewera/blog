package files

import (
	"fmt"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"io"
	"net/http"
)

func DownloadFile(url string, targetFile string, executable bool) error {
	file, err := CreateFileWr(targetFile, executable)
	if err != nil {
		return err
	}
	defer must.Close(file)

	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("file: download: %w", err)
	}
	defer must.Close(res.Body)

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("file: download: copy buffer: %w", err)
	}
	return nil
}
