package tailwind

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/must"
	"io"
	"os"
	"strings"
)

func checkSha256(upstreamFilename string, localFilename string) error {
	expectedChecksum, err := downloadAndExtractChecksum(upstreamFilename)
	if err != nil {
		return err
	}

	f, err := os.Open(localFilename)
	if err != nil {
		return err
	}
	defer must.Close(f)

	h := sha256.New()
	if _, err = io.Copy(h, f); err != nil {
		return err
	}

	actualChecksum := fmt.Sprintf("%x", h.Sum(nil))
	if actualChecksum != expectedChecksum {
		return fmt.Errorf(
			"checksums: not equal: expected='%s'; actual='%s'; filename='%s'",
			expectedChecksum,
			actualChecksum,
			localFilename,
		)
	}
	return nil
}

func downloadAndExtractChecksum(upstreamFilename string) (sha256checksum string, err error) {
	err = files.DownloadFile(upstreamChecksumsURL(), constants.TailwindChecksum, false)
	if err != nil {
		return "", err
	}

	checksums, err := os.Open(constants.TailwindChecksum)
	if err != nil {
		return "", err
	}

	sha256checksum, err = findChecksum(checksums, upstreamFilename)
	if err != nil {
		return "", err
	}

	err = os.Remove(constants.TailwindChecksum)
	if err != nil {
		return "", fmt.Errorf("remove checksums file: %s: %w", constants.TailwindChecksum, err)
	}
	return sha256checksum, nil
}

func findChecksum(checksums io.Reader, upstreamFilename string) (checksum string, err error) {
	s := bufio.NewScanner(checksums)
	for s.Scan() {
		line := s.Text()
		f := strings.Fields(line)
		if len(f) != 2 {
			return "", fmt.Errorf("checksums: wrong file format; expected line 'downloadAndExtractChecksum filename'")
		}
		checksum = f[0]
		filename := f[1]
		if strings.Contains(filename, upstreamFilename) {
			return checksum, nil
		}
	}

	if err = s.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("checksums: not found for: %s", upstreamFilename)
}
