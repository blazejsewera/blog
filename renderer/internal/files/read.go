package files

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/internal/must"
)

func Read(filename string) string {
	log.Debug("files: reading file: %s", filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Errorf("read file: %w", err))
	}
	defer must.Close(file)

	b := &bytes.Buffer{}
	_, err = io.Copy(b, file)
	if err != nil {
		panic(fmt.Errorf("read file: %w", err))
	}

	return b.String()
}
