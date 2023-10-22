package preprocess

import (
	"github.com/blazejsewera/blog/preprocess/font"
	"github.com/blazejsewera/blog/preprocess/tailwind"
)

func Run() {
	tailwind.Run(false)
	font.Download(false)
}
