package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Track(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("track", attributes...)
}
