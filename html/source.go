package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Source(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("source", attributes...)
}
