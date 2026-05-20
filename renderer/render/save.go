package render

import (
	"fmt"
	"os"

	"github.com/blazejsewera/blog/renderer/internal/log"
	"github.com/blazejsewera/blog/renderer/internal/must"
)

func save(targetFile string, rendered []byte, sourceFile string) error {
	target, err := os.Create(targetFile)
	if err != nil {
		return renderErr(sourceFile, err)
	}
	defer must.Close(target)
	_, err = target.Write(rendered)
	if err != nil {
		return renderErr(sourceFile, err)
	}
	log.Info("render: saved: source: %s; target: %s", sourceFile, targetFile)
	return nil
}

func renderErr(sourceFile string, err error) error {
	return fmt.Errorf("render: %s: %w", sourceFile, err)
}
