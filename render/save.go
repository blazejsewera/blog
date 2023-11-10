package render

import (
	"fmt"
	"github.com/blazejsewera/blog/postprocess"
	"os"
)

func save(targetFile string, rendered []byte, sourceFile string) error {
	target, err := os.Create(targetFile)
	if err != nil {
		return renderErr(sourceFile, err)
	}
	_, err = target.Write(postprocess.MinifyHTML(rendered))
	if err != nil {
		return renderErr(sourceFile, err)
	}
	return nil
}

func renderErr(sourceFile string, err error) error {
	return fmt.Errorf("render: %s: %w", sourceFile, err)
}
