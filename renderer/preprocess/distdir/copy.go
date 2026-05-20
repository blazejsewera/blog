package distdir

import (
	"fmt"

	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/internal/files"
	"github.com/blazejsewera/blog/renderer/internal/log"
)

func CopyIfDoesNotExist(force constants.ForceLevel) {
	if force >= constants.RemoveAndReRender {
		log.Debug("dist: removing")
		files.RemoveAll(constants.DistDir)
		log.Info("dist: removed")
	}
	if force >= constants.ReRender || !files.Exists(constants.DistDir) {
		log.Debug("dist: copying")
		err := files.CopyDir(constants.DistDir, constants.SiteDir)
		if err != nil {
			panic(fmt.Errorf("dist: copy: %w", err))
		}
		log.Info("dist: copied: source: %s; target: %s", constants.SiteDir, constants.DistDir)
	}
}
