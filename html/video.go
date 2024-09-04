package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Video(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("video", attributes...)
}
