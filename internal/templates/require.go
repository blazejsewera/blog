package templates

import "fmt"

func require(s string) (string, error) {
	if s == "" {
		return "", fmt.Errorf("require: value was empty")
	}
	return s, nil
}
