package distdir

import (
	"fmt"
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
)

func CopyIfDoesNotExist(force constants.ForceLevel) {
	if force >= constants.RemoveAndReRender {
		log.Info("dist: remove")
		files.RemoveAll(constants.DistDir)
	}
	if force >= constants.ReRender || !files.Exists(constants.DistDir) {
		log.Info("dist: copying")
		err := files.CopyDir(constants.DistDir, constants.SiteDir)
		if err != nil {
			panic(fmt.Errorf("dist: copy: %w", err))
		}
	}
}
