package distdir

import (
	"fmt"
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/internal/files"
)

func CopyIfDoesNotExist(force constants.ForceLevel) {
	if force >= constants.ReRender {
		files.RemoveAll(constants.DistDir)
	}
	if !files.Exists(constants.DistDir) {
		err := files.CopyDir(constants.DistDir, constants.SiteDir)
		if err != nil {
			panic(fmt.Errorf("copy dist: %w", err))
		}
	}
}
