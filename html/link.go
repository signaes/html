package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Link(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("link", attributes...)
}
