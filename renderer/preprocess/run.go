package preprocess

import (
	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/preprocess/font"
	"github.com/blazejsewera/blog/renderer/preprocess/tailwind"
)

func Run(force constants.ForceLevel) {
	tailwind.Run(force)
	font.Download(force)
}
