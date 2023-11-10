package distdir

import (
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/internal/files"
	"log"
)

func CopyIfDoesNotExist(force constants.ForceLevel) {
	if force >= constants.ReRender {
		log.Print("info: dist: remove")
		files.RemoveAll(constants.DistDir)
	}
	if !files.Exists(constants.DistDir) {
		log.Print("info: dist: copying")
		err := files.CopyDir(constants.DistDir, constants.SiteDir)
		if err != nil {
			panic(fmt.Errorf("dist: copy: %w", err))
		}
	}
}
