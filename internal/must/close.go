package must

import (
	"fmt"
	"io"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		panic(fmt.Errorf("must close: %w", err))
	}
}
