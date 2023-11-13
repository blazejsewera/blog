package distdir

import (
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/internal/files"
	"github.com/blazejsewera/blog/internal/log"
)

func CopyIfDoesNotExist(force constants.ForceLevel) {
	if force >= constants.RemoveAndReRender {
		log.Info("dist: remove")
		files.RemoveAll(constants.DistDir)
	} else if force >= constants.ReRender || !files.Exists(constants.DistDir) {
		log.Info("dist: copying")
		err := files.CopyDir(constants.DistDir, constants.SiteDir)
		if err != nil {
			panic(fmt.Errorf("dist: copy: %w", err))
		}
	}
}
