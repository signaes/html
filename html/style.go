package html

import (
	"fmt"

	"github.com/signaes/elements/elements"
)

func Style(attributes ...fmt.Stringer) func(children ...fmt.Stringer) elements.Element {
	return elements.New("style", attributes...)
}
