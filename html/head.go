package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Head(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("head", attributes...)
}
