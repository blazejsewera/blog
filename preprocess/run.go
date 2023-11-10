package preprocess

import (
	"github.com/blazejsewera/blog/constants"
	"github.com/blazejsewera/blog/preprocess/font"
	"github.com/blazejsewera/blog/preprocess/tailwind"
)

func Run(force constants.ForceLevel) {
	tailwind.Run(force)
	font.Download(force)
}
