package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Figcaption(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("figcaption", attributes...)
}
