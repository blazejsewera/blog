package files

import (
	"errors"
	"os"
)

func Exists(filename string) bool {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		return false
	} else {
		return true
	}
}
