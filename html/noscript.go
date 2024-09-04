package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Noscript(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("noscript", attributes...)
}
