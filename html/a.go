package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func A(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("a", attributes...)
}
