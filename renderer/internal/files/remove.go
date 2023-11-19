package files

import "os"

func RemoveAll(file string) {
	err := os.RemoveAll(file)
	if err != nil {
		panic(err)
	}
}
