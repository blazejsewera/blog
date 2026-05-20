package preprocess

import (
	"sync"

	"github.com/blazejsewera/blog/renderer/constants"
	"github.com/blazejsewera/blog/renderer/preprocess/font"
	"github.com/blazejsewera/blog/renderer/preprocess/tailwind"
)

func Run(force constants.ForceLevel) {
	wg := sync.WaitGroup{}
	wg.Go(func() {
		tailwind.Run(force)
	})
	wg.Go(func() {
		font.Download(force)
	})
	wg.Wait()
}
