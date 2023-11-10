package files

import (
	"bytes"
	"fmt"
	"github.com/blazejsewera/blog/internal/must"
	"io"
	"os"
)

func Read(filename string) string {
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
