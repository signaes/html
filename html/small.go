package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Small(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("small", attributes...)
}
