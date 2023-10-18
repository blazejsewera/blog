package files

import "os"

func CreateDirIfDoesNotExist(name string) error {
	if !Exists(name) {
		err := os.Mkdir(name, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
