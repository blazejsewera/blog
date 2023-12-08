package render

import (
	"fmt"
	"github.com/blazejsewera/blog/renderer/internal/log"
	"os"
)

func save(targetFile string, rendered []byte, sourceFile string) error {
	target, err := os.Create(targetFile)
	if err != nil {
		return renderErr(sourceFile, err)
	}
	_, err = target.Write(rendered)
	if err != nil {
		return renderErr(sourceFile, err)
	}
	log.Info("rendered: source: %s; target: %s", sourceFile, targetFile)
	return nil
}

func renderErr(sourceFile string, err error) error {
	return fmt.Errorf("render: %s: %w", sourceFile, err)
}
