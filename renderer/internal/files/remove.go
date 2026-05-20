package files

import (
	"os"

	"github.com/blazejsewera/blog/renderer/internal/log"
)

func RemoveAll(file string) {
	log.Debug("files: removing: %s", file)
	err := os.RemoveAll(file)
	if err != nil {
		panic(err)
	}
}
